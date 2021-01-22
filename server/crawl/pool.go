package crawl

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Peer defines a node structure that exists in the NodePool. Every Peer should
// have an RPC address defined, but a network is not strictly required.
type Peer struct {
	RPCAddr string
	Network string
}

func (p Peer) String() string {
	if p.RPCAddr != "" && p.Network != "" {
		return p.RPCAddr + ";" + p.Network
	}

	return p.RPCAddr
}

// NodePool implements an abstraction over a pool of nodes for which to crawl.
// It also contains a collection of nodes for which to reseed the pool when it's
// empty. Once the reseed list has reached capacity, a random node is removed
// when another is added.
type NodePool struct {
	rw sync.RWMutex

	nodes       map[Peer]struct{}
	reseedNodes []Peer
	rng         *rand.Rand
}

func NewNodePool(reseedCap uint) *NodePool {
	return &NodePool{
		nodes:       make(map[Peer]struct{}),
		reseedNodes: make([]Peer, 0, reseedCap),
		rng:         rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// Size returns the size of the pool.
func (np *NodePool) Size() int {
	np.rw.RLock()
	defer np.rw.RUnlock()
	return len(np.nodes)
}

// Seed seeds the node pool with a given set of nodes. For every seed, we split
// it on a ';' delimiter to get the RPC address and the network (if provided).
func (np *NodePool) Seed(seeds []string) {
	for _, s := range seeds {
		tokens := strings.Split(s, ";")
		switch len(tokens) {
		case 1:
			np.AddNode(Peer{RPCAddr: tokens[0]})

		case 2:
			np.AddNode(Peer{RPCAddr: tokens[0], Network: tokens[1]})
		}
	}
}

// RandomNode returns a random node, based on Golang's map semantics, from the
// pool.
func (np *NodePool) RandomNode() (Peer, bool) {
	np.rw.RLock()
	defer np.rw.RUnlock()

	for nodeRPCAddr := range np.nodes {
		return nodeRPCAddr, true
	}

	return Peer{}, false
}

// AddNode adds a node to the node pool by adding it to the internal node list.
// In addition, we attempt to add it to the internal reseed node list. If the
// reseed list is full, it replaces a random node in the reseed list, otherwise
// it is directly added to it.
func (np *NodePool) AddNode(p Peer) {
	np.rw.Lock()
	defer np.rw.Unlock()

	np.nodes[p] = struct{}{}

	if len(np.reseedNodes) < cap(np.reseedNodes) {
		np.reseedNodes = append(np.reseedNodes, p)
	} else {
		// replace random node with the new node
		i := np.rng.Intn(len(np.reseedNodes))
		np.reseedNodes[i] = p
	}
}

// HasNode returns true if a node exists in the node pool and false otherwise.
func (np *NodePool) HasNode(p Peer) bool {
	np.rw.RLock()
	defer np.rw.RUnlock()

	_, ok := np.nodes[p]
	return ok
}

// DeleteNode removes a node from the node pool if it exists.
func (np *NodePool) DeleteNode(p Peer) {
	np.rw.Lock()
	defer np.rw.Unlock()
	delete(np.nodes, p)
}

// Reseed seeds the node pool with all the nodes found in the internal reseed
// list.
func (np *NodePool) Reseed() {
	np.rw.Lock()
	defer np.rw.Unlock()

	for _, p := range np.reseedNodes {
		np.nodes[p] = struct{}{}
	}
}
