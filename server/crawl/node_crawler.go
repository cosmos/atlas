package crawl

import (
	"errors"
	"fmt"
	"strings"
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
)

// Crawler implements the Tendermint p2p network crawler.
type Crawler struct {
	logger   zerolog.Logger
	db       *gorm.DB
	pool     *NodePool
	ipClient *ipstack.Client
	locCache *lru.ARCCache
	seeds    []string
	doneCh   chan struct{}

	crawlInterval   time.Duration
	recheckInterval time.Duration
}

func NewCrawler(logger zerolog.Logger, cfg config.Config, db *gorm.DB) (*Crawler, error) {
	locCache, err := lru.NewARC(locationCacheSize)
	if err != nil {
		return nil, err
	}

	return &Crawler{
		logger:          logger.With().Str("module", "node_crawler").Logger(),
		db:              db,
		seeds:           strings.Split(cfg.String(config.NodeSeeds), ","),
		crawlInterval:   cfg.Duration(config.NodeCrawlInterval),
		recheckInterval: cfg.Duration(config.NodeRecheckInterval),
		ipClient:        ipstack.NewClient(cfg.String(config.IPStackKey), false, 5),
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
	c.logger.Info().Msg("starting node crawler...")

	// seed the pool with the initial set of seeds before crawling
	c.pool.Seed(c.seeds)

	go c.RecheckNodes()

	ticker := time.NewTicker(c.crawlInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Keep picking a pseudo-random node from the pool to crawl until the pool
			// is exhausted.
			nodeRPCAddr, ok := c.pool.RandomNode()
			for ok {
				c.CrawlNode(nodeRPCAddr)
				c.pool.DeleteNode(nodeRPCAddr)

				// pick the next pseudo-random node
				nodeRPCAddr, ok = c.pool.RandomNode()
			}

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

				c.logger.Debug().
					Str("p2p_address", nodeP2PAddr).
					Str("rpc_address", nodeRPCAddr).
					Msg("adding node to node pool")
				c.pool.AddNode(nodeRPCAddr)
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
func (c *Crawler) CrawlNode(nodeRPCAddr string) {
	host := parseHostname(nodeRPCAddr)
	nodeP2PAddr := fmt.Sprintf("%s:%s", host, defaultP2PPort)

	node := models.Node{
		Address: host,
		RPCPort: parsePort(nodeRPCAddr),
		P2PPort: defaultP2PPort,
	}

	c.logger.Debug().
		Str("p2p_address", nodeP2PAddr).
		Str("rpc_address", nodeRPCAddr).
		Msg("pinging node...")

	// Attempt to ping the node where upon failure, we remove the node from the
	// database.
	if ok := pingAddress(nodeP2PAddr, 5*time.Second); !ok {
		c.logger.Info().
			Str("p2p_address", nodeP2PAddr).
			Str("rpc_address", nodeRPCAddr).
			Msg("failed to ping node; deleting...")

		if err := node.Delete(c.db); err != nil {
			c.logger.Info().
				Err(err).
				Str("p2p_address", nodeP2PAddr).
				Str("rpc_address", nodeRPCAddr).
				Msg("failed to delete node")
		}

		return
	}

	// Grab the node's geolocation information. Failure indicates we
	// continue to crawl the node.
	loc, err := c.GetGeolocation(node.Address)
	if err != nil {
		c.logger.Info().
			Err(err).
			Str("p2p_address", nodeP2PAddr).
			Str("rpc_address", nodeRPCAddr).
			Msg("failed to get node geolocation")
		return
	}

	node.Location = loc
	client := newRPCClient(nodeRPCAddr)

	// Attempt to get the node's status which provides us with rich information
	// about the node. Upon failure, we still crawl and persist the node but we
	// lack most useful information about the node.
	status, err := client.Status()
	if err != nil {
		c.logger.Info().
			Err(err).Str("p2p_address", nodeP2PAddr).
			Str("rpc_address", nodeRPCAddr).
			Msg("failed to get node status")
	} else {
		node.Moniker = status.NodeInfo.Moniker
		node.NodeID = string(status.NodeInfo.ID())
		node.Network = status.NodeInfo.Network
		node.Version = status.NodeInfo.Version
		node.TxIndex = status.NodeInfo.Other.TxIndex

		netInfo, err := client.NetInfo()
		if err != nil {
			c.logger.Info().
				Err(err).
				Str("p2p_address", nodeP2PAddr).
				Str("rpc_address", nodeRPCAddr).
				Msg("failed to get node net info")
			return
		}

		// add the relevant peers to pool
		for _, p := range netInfo.Peers {
			peerRPCPort := parsePort(p.NodeInfo.Other.RPCAddress)
			peerRPCAddress := fmt.Sprintf("http://%s:%s", p.RemoteIP, peerRPCPort)
			peer := models.Node{
				Address: p.RemoteIP,
				Network: node.Network,
			}

			// only add the peer to the pool if we haven't (re)discovered it
			_, err := models.QueryNode(
				c.db,
				map[string]interface{}{"address": peer.Address, "network": peer.Network},
			)
			if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
				c.logger.Debug().
					Str("peer_rpc_address", peerRPCAddress).
					Msg("adding peer to node pool")
				c.pool.AddNode(peerRPCAddress)
			}
		}
	}

	if _, err := node.Upsert(c.db); err != nil {
		c.logger.Info().
			Err(err).
			Str("p2p_address", nodeP2PAddr).
			Str("rpc_address", nodeRPCAddr).
			Msg("failed to save node")
	} else {
		c.logger.Info().
			Str("p2p_address", nodeP2PAddr).
			Str("rpc_address", nodeRPCAddr).
			Msg("successfully crawled and saved node")
	}
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
