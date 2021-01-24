# Node Explorer

Atlas provides Tendermint node crawling functionality that allows for users to
explore the topology of various Tendermint-based networks and understand what
kinds of nodes exist in a network through node metadata (version, location, etc...).

Atlas periodically crawls Tendermint-based networks and persists nodes to a
database. It does so by being provided an initial list of seed nodes with an
optional network (chain-id) identifier. These seed nodes are then used to further
explore the topology of the various networks these nodes run on.

Internally, Atlas maintains a pool of nodes for which to crawl. Random nodes are
picked off of and removed from this pool when crawling. Then, any undiscovered
peers are added to the pool to be subsequently crawled as well. If the pool is
ever exhausted, the pool is reseeded and the crawling beings again after some
time duration (see below). This process runs in its own separate goroutine.

In order not to keep around nodes that are no longer reachable or are part of
their respective network around, Atlas also runs a recheck process, also in a
separate goroutine, where it fetches all stale nodes and rechecks them for their
availability and potentially any new information.

The following configuration parameters, which may be provided as environment
variables or in a config file, are used to tune the crawling functionality:

- `ipstack API key`: The API key for the [ipstack](https://ipstack.com/) service.
  This service provides free IP to geolocation APIs which are used to determine
  geographical information about crawled nodes.
- `crawl interval`: The time duration between successive crawling attempts. A new
  crawl is only triggered after the internal node pool is exhausted and the crawl
  interval ticker is triggered. Note, depending how the node pool is depleted and
  if and how many new peers are discovered, subsequent node crawling attempts may
  not always reflect this duration accurately.
- `recheck interval`: The time duration between successive stale node rechecks
  sweeps. During every trigger of this interval, Atlas will check for all stale
  nodes and recheck if they are still reachable and update any relevant information
  about each node.
- `reseed size`: The max capacity of the list of nodes for which Atlas will attempt
  to reseed the internal node pool between successive crawl attempts.
- `seeds`: The initial list of comma-delimited seed nodes for Atlas to crawl.
  This list initially populates the internal node pool. A seed node takes the
  form of `[host]:[port];[network]`, where `;[network]` is optional
  (e.g. `http://1.255.51.125:26657;cosmoshub-3`). It's ideal to provide a large
  enough list of healthy and reachable nodes in order for Atlas to successfully
  explore the various networks the seed nodes represent.
