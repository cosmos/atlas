BEGIN;
-- 
-- Create locations table
-- 
CREATE TABLE IF NOT EXISTS locations (
  id SERIAL PRIMARY KEY,
  country VARCHAR NOT NULL,
  region VARCHAR NOT NULL,
  city VARCHAR NOT NULL,
  latitude VARCHAR NOT NULL,
  longitude VARCHAR NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_locations_country ON locations(country);
CREATE INDEX IF NOT EXISTS idx_locations_city ON locations(city);
CREATE UNIQUE INDEX IF NOT EXISTS idx_locations_lat_long ON locations(latitude, longitude);
-- 
-- Create nodes table
-- 
CREATE TABLE IF NOT EXISTS nodes (
  id SERIAL PRIMARY KEY,
  location_id INT NOT NULL,
  address VARCHAR NOT NULL,
  rpc_port VARCHAR NOT NULL,
  p2p_port VARCHAR NOT NULL,
  moniker VARCHAR NOT NULL,
  node_id VARCHAR NOT NULL,
  network VARCHAR NOT NULL,
  version VARCHAR NOT NULL,
  tx_index VARCHAR NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  FOREIGN KEY (location_id) REFERENCES locations(id)
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_nodes_addr_network ON nodes(address, network);
CREATE INDEX IF NOT EXISTS idx_nodes_moniker ON nodes(moniker);
COMMIT;
