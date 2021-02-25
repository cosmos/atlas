package crawl

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/harwoeck/ipstack"
	lru "github.com/hashicorp/golang-lru"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/cosmos/atlas/config"
	"github.com/cosmos/atlas/server/models"
)

const (
	defaultP2PPort    = "26656"
	locationCacheSize = 1000
	ipClientHTTPS     = false
	ipClientTimeoutS  = 5
)

// Crawler implements the Tendermint p2p network crawler.
type Crawler struct {
	logger   zerolog.Logger
	pool     *NodePool
	ipClient *ipstack.Client
	locCache *lru.ARCCache
	seeds    []string
	doneCh   chan struct{}

	mtx sync.Mutex
	db  *gorm.DB
	// tmpPeers is buffer that is used to hold peers that will later be added to
	// the pool after a crawl is complete.
	tmpPeers []Peer

	crawlInterval   time.Duration
	recheckInterval time.Duration
}

func NewCrawler(logger zerolog.Logger, cfg config.Config, db *gorm.DB) (*Crawler, error) {
	locCache, err := lru.NewARC(locationCacheSize)
	if err != nil {
		return nil, err
	}

	logger = logger.With().Str("module", "node_crawler").Logger()

	if cfg.Duration(config.NodeCrawlInterval).Seconds() == 0 {
		logger.Info().Msg("node crawling disabled")
		return nil, nil
	}

	return &Crawler{
		logger:          logger,
		db:              db,
		seeds:           strings.Split(cfg.String(config.NodeSeeds), ","),
		crawlInterval:   cfg.Duration(config.NodeCrawlInterval),
		recheckInterval: cfg.Duration(config.NodeRecheckInterval),
		ipClient:        ipstack.NewClient(cfg.String(config.IPStackKey), ipClientHTTPS, ipClientTimeoutS),
		locCache:        locCache,
		pool:            NewNodePool(uint(cfg.Int(config.NodeReseedSize))),
		doneCh:          make(chan struct{}),
	}, nil
}

// Stop signals to the crawler that it should halt and exit all spawned goroutines.
func (c *Crawler) Stop() {
	close(c.doneCh)
}

// Start starts a blocking process in which a random node is selected from the
// node pool and crawled. For each successful crawl, it'll be persisted or updated
// and its peers will be added to the node pool if they do not already exist.
// This process continues indefinitely until all nodes are exhausted from the pool.
// When the pool is empty and after crawlInterval seconds since the last complete
// crawl, a random set of nodes from the DB are added to reseed the pool.
func (c *Crawler) Start() {
	// seed the pool with the initial set of seeds before crawling
	c.pool.Seed(c.seeds)

	go c.RecheckNodes()

	ticker := time.NewTicker(c.crawlInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.logger.Info().Msg("starting to crawl nodes")

			// reset the peer buffer
			c.mtx.Lock()
			c.tmpPeers = make([]Peer, 0)
			c.mtx.Unlock()

			var wg sync.WaitGroup
			nc := 0
			start := time.Now()

			// Keep picking a pseudo-random node from the pool to crawl until the pool
			// is exhausted.
			peer, ok := c.pool.RandomNode()
			for ok {
				wg.Add(1)
				c.pool.DeleteNode(peer)

				go func(p Peer) {
					defer wg.Done()
					c.CrawlNode(p)
				}(peer)

				// pick the next pseudo-random node
				peer, ok = c.pool.RandomNode()

				if nc%50 == 0 {
					c.logger.Info().Int("size", c.pool.Size()).Msg("node pool size")
				}

				nc++
			}

			// wait for all crawlers to complete
			wg.Wait()

			// add all peers from the temp buffer to the node pool
			c.mtx.Lock()
			for _, p := range c.tmpPeers {
				c.logger.Debug().Str("rpc_address", p.RPCAddr).Msg("adding peer to node pool")
				c.pool.AddNode(p)
			}
			c.mtx.Unlock()

			elapsed := time.Since(start).Seconds()

			c.logger.Info().Int("num_crawled", nc).
				Float64("elapsed", elapsed).
				Msg("node crawl complete; reseeding node pool")
			c.pool.Reseed()

		case <-c.doneCh:
			return
		}
	}
}

// RecheckNodes starts a blocking process where every recheckInterval duration
// the crawler checks for all stale nodes that need to be rechecked. For each
// stale node, the node is added back into the node pool to be re-crawled and
// updated (or removed).
func (c *Crawler) RecheckNodes() {
	ticker := time.NewTicker(c.recheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.logger.Info().Msg("rechecking nodes...")

			nodes, err := models.GetStaleNodes(c.db, time.Now())
			if err != nil {
				c.logger.Info().Err(err).Msg("failed to get all stale nodes")
				continue
			}

			for _, node := range nodes {
				nodeP2PAddr := fmt.Sprintf("%s:%s", node.Address, node.P2PPort)
				nodeRPCAddr := fmt.Sprintf("http://%s:%s", node.Address, node.RPCPort)

				p := Peer{RPCAddr: nodeRPCAddr, Network: node.Network}
				if !c.pool.HasNode(p) {
					c.logger.Debug().
						Str("p2p_address", nodeP2PAddr).
						Str("rpc_address", nodeRPCAddr).
						Time("last_sync", node.UpdatedAt).
						Msg("adding stale node to node pool")
					c.pool.AddNode(p)
				}
			}

		case <-c.doneCh:
			return
		}
	}
}

// CrawlNode performs the main crawling functionality for a Tendermint node. It
// accepts a node RPC address and attempts to ping that node's P2P address by
// using the RPC address and the default P2P port of 26656. If the P2P address
// cannot be reached, the node is deleted if it exists in the database. Otherwise,
// we attempt to get additional metadata aboout the node via it's RPC address
// and its set of peers. For every peer that doesn't exist in the node pool, it
// is added.
func (c *Crawler) CrawlNode(p Peer) {
	host := parseHostname(p.RPCAddr)
	nodeP2PAddr := fmt.Sprintf("%s:%s", host, defaultP2PPort)

	node := models.Node{
		Address: host,
		RPCPort: parsePort(p.RPCAddr),
		P2PPort: defaultP2PPort,
		Network: p.Network,
	}

	var deleteNode bool
	defer func() {
		if deleteNode {
			c.deleteNode(node)
		}
	}()

	c.logger.Debug().Str("p2p_address", nodeP2PAddr).Str("rpc_address", p.RPCAddr).Msg("pinging node...")

	// Attempt to ping the node where upon failure, we remove the node from the
	// database.
	if ok := pingAddress(nodeP2PAddr, 5*time.Second); !ok {
		c.logger.Info().
			Str("p2p_address", nodeP2PAddr).
			Str("rpc_address", p.RPCAddr).
			Msg("failed to ping node; deleting...")

		deleteNode = true
		return
	}

	// Grab the node's geolocation information where upon failure, we remove the
	// node from the database.
	loc, err := c.GetGeolocation(node.Address)
	if err != nil {
		c.logger.Error().
			Err(err).
			Str("p2p_address", nodeP2PAddr).
			Str("rpc_address", p.RPCAddr).
			Msg("failed to get node geolocation; deleting...")

		deleteNode = true
		return
	}

	node.Location = loc

	client, err := newRPCClient(p.RPCAddr, clientTimeout)
	if err != nil {
		c.logger.Error().
			Err(err).
			Str("p2p_address", nodeP2PAddr).
			Str("rpc_address", p.RPCAddr).
			Msg("failed to create RPC client")

		return
	}

	// Attempt to get the node's status which provides us with node metadata.
	// Upon failure, we return and prevent further crawling if the network is
	// unknown due to the lack of any useful information about the node.
	status, err := client.Status(context.Background())
	if err != nil {
		c.logger.Error().
			Err(err).
			Str("p2p_address", nodeP2PAddr).
			Str("rpc_address", p.RPCAddr).
			Msg("failed to get node status")

		if node.Network == "" {
			deleteNode = true
			return
		}
	} else {
		node.Moniker = status.NodeInfo.Moniker
		node.NodeID = string(status.NodeInfo.ID())
		node.Version = status.NodeInfo.Version
		node.TxIndex = status.NodeInfo.Other.TxIndex

		if node.Network == "" {
			node.Network = status.NodeInfo.Network
		}

		netInfo, err := client.NetInfo(context.Background())
		if err != nil {
			c.logger.Error().
				Err(err).
				Str("p2p_address", nodeP2PAddr).
				Str("rpc_address", p.RPCAddr).
				Msg("failed to get node net info")
		} else {
			// Add the relevant peers to the temp buffer which will later be added to
			// the node pool.
			for _, p := range netInfo.Peers {
				peerRPCPort := parsePort(p.NodeInfo.Other.RPCAddress)
				peerRPCAddress := fmt.Sprintf("http://%s:%s", p.RemoteIP, peerRPCPort)

				// only add the peer to the pool if we haven't (re)discovered it
				_, err := models.QueryNode(
					c.db,
					map[string]interface{}{"address": p.RemoteIP, "network": node.Network},
				)
				if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
					c.mtx.Lock()
					c.tmpPeers = append(c.tmpPeers, Peer{RPCAddr: peerRPCAddress, Network: node.Network})
					c.mtx.Unlock()
				}
			}
		}
	}

	c.upsertNode(node)
}

// GetGeolocation returns a Location record containing geolocation information
// for a given node. It will first check to see if the location already exists
// in cache. If the record does not exist in the cache, a Node record is queried
// by the provided address. If that record does not exist, we perform a query
// against the ipstack API and write to the cache. An error is returned if the
// database query fails.
func (c *Crawler) GetGeolocation(addr string) (models.Location, error) {
	// return the location from cache if it exists
	if loc, ok := c.locCache.Get(addr); ok {
		return loc.(models.Location), nil
	}

	var loc models.Location

	// Query for the Node record and if the record exists, use that Location.
	// Otherwise, perform a query against ipstack using the provided address.
	node, err := models.QueryNode(c.db, map[string]interface{}{"address": addr})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ipResp, err := c.ipClient.Check(addr)
			if err != nil {
				return models.Location{}, err
			}

			loc = locationFromIPResp(ipResp)
		} else {
			return models.Location{}, err
		}
	} else {
		loc = node.Location
	}

	// write to cache
	c.locCache.Add(addr, loc)

	return loc, nil
}

// deleteNode provides a thread-safe way of deleting the given node from the
// database. Concurrent goroutines are spawned for each node to crawl, so we
// use the crawler's mutex to prevent any issues with concurrent database
// operations.
func (c *Crawler) deleteNode(n models.Node) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if err := n.Delete(c.db); err != nil {
		c.logger.Error().Err(err).Str("rpc_address", n.Address).Msg("failed to delete node")
	}
}

// upsertNode provides a thread-safe way of updating the given node from the
// database. Concurrent goroutines are spawned for each node to crawl, so we
// use the crawler's mutex to prevent any issues with concurrent database
// operations.
func (c *Crawler) upsertNode(n models.Node) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if _, err := n.Upsert(c.db); err != nil {
		c.logger.Error().Err(err).Str("rpc_address", n.Address).Msg("failed to save node")
	} else {
		c.logger.Info().Str("rpc_address", n.Address).Msg("successfully crawled and saved node")
	}
}
