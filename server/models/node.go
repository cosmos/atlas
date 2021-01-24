package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/cosmos/atlas/server/httputil"
)

// LocationJSON defines the JSON-encodeable type for a Location.
type LocationJSON struct {
	GormModelJSON

	Country   string `json:"country"`
	Region    string `json:"region"`
	City      string `json:"city"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Location defines the geographical location of a crawled Tendermint node.
type Location struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Country   string
	Region    string
	City      string
	Latitude  string
	Longitude string
}

// NodeJSON defines the JSON-encodeable type for a Node.
type NodeJSON struct {
	GormModelJSON

	Location LocationJSON `json:"location"`
	Address  string       `json:"address"`
	RPCPort  string       `json:"rpc_port"`
	P2PPort  string       `json:"p2p_port"`
	Moniker  string       `json:"moniker"`
	NodeID   string       `json:"node_id"`
	Network  string       `json:"network"`
	Version  string       `json:"version"`
	TxIndex  string       `json:"tx_index"`
}

// Node defines a crawled Tendermint node.
type Node struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	LocationID uint
	Location   Location
	Address    string
	RPCPort    string `gorm:"column:rpc_port"`
	P2PPort    string `gorm:"column:p2p_port"`
	Moniker    string
	NodeID     string
	Network    string
	Version    string
	TxIndex    string
}

// MarshalJSON implements custom JSON marshaling for the Location model.
func (l Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.NewLocationJSON())
}

func (l Location) NewLocationJSON() LocationJSON {
	return LocationJSON{
		GormModelJSON: GormModelJSON{
			ID:        l.ID,
			CreatedAt: l.CreatedAt,
			UpdatedAt: l.UpdatedAt,
		},
		Country:   l.Country,
		Region:    l.Region,
		City:      l.City,
		Latitude:  l.Latitude,
		Longitude: l.Longitude,
	}
}

// MarshalJSON implements custom JSON marshaling for the Node model.
func (n Node) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.NewNodeJSON())
}

func (n Node) NewNodeJSON() NodeJSON {
	return NodeJSON{
		GormModelJSON: GormModelJSON{
			ID:        n.ID,
			CreatedAt: n.CreatedAt,
			UpdatedAt: n.UpdatedAt,
		},
		Location: n.Location.NewLocationJSON(),
		Address:  n.Address,
		RPCPort:  n.RPCPort,
		P2PPort:  n.P2PPort,
		Moniker:  n.Moniker,
		NodeID:   n.NodeID,
		Network:  n.Network,
		Version:  n.Version,
		TxIndex:  n.TxIndex,
	}
}

// Upsert creates or updates a Location record. If no record exists, a new one
// will be created. Otherwise, the existing record is updated. An error is returned
// upon failure. The updated or created record is returned upon success.
func (l Location) Upsert(db *gorm.DB) (Location, error) {
	if l.Longitude == "" || l.Latitude == "" {
		return Location{}, errors.New("longitude and latitude are required")
	}

	var record Location

	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("latitude = ? AND longitude = ?", l.Latitude, l.Longitude).First(&record).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := tx.Create(&l).Error; err != nil {
					return fmt.Errorf("failed to create location: %w", err)
				}

				// commit the tx
				return nil
			} else {
				return fmt.Errorf("failed to query for location: %w", err)
			}
		}

		record.Country = l.Country
		record.Region = l.Region
		record.City = l.City
		record.Latitude = l.Latitude
		record.Longitude = l.Longitude
		if err := tx.Save(&record).Error; err != nil {
			return fmt.Errorf("failed to update location: %w", err)
		}

		// commit the tx
		return nil
	})
	if err != nil {
		return Location{}, err
	}

	return QueryLocation(db, map[string]interface{}{"latitude": l.Latitude, "longitude": l.Longitude})
}

// QueryLocation performs a query for a Location record. The resulting record,
// if it exists, is returned. If the query fails or the record does not exist,
// an error is returned.
func QueryLocation(db *gorm.DB, query map[string]interface{}) (Location, error) {
	var record Location

	if err := db.Where(query).First(&record).Error; err != nil {
		return Location{}, fmt.Errorf("failed to query location: %w", err)
	}

	return record, nil
}

// Upsert creates or updates a Node record. If no record exists, a new one will
// be created. Otherwise, the existing record is updated. An error is returned
// upon failure. The updated or created record is returned upon success.
func (n Node) Upsert(db *gorm.DB) (Node, error) {
	var record Node

	err := db.Transaction(func(tx *gorm.DB) error {
		loc, err := n.Location.Upsert(tx)
		if err != nil {
			return err
		}

		n.Location = loc

		err = tx.Where("address = ? AND network = ?", n.Address, n.Network).First(&record).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := tx.Create(&n).Error; err != nil {
					return fmt.Errorf("failed to create node: %w", err)
				}

				// commit the tx
				return nil
			} else {
				return fmt.Errorf("failed to query for node: %w", err)
			}
		}

		if err := tx.Model(&record).Association("Location").Replace(&n.Location); err != nil {
			return fmt.Errorf("failed to update node location: %w", err)
		}

		record.Address = n.Address
		record.RPCPort = n.RPCPort
		record.P2PPort = n.P2PPort
		record.Moniker = n.Moniker
		record.NodeID = n.NodeID
		record.Network = n.Network
		record.Version = n.Version
		record.TxIndex = n.TxIndex
		if err := tx.Save(&record).Error; err != nil {
			return fmt.Errorf("failed to update node: %w", err)
		}

		// commit the tx
		return nil
	})
	if err != nil {
		return Node{}, err
	}

	return QueryNode(db, map[string]interface{}{"address": n.Address, "network": n.Network})
}

// Delete attempts to delete a Node record by its address. An error is returned
// upon query or delete failure. An error is not returned if the record does not
// exist.
func (n Node) Delete(db *gorm.DB) error {
	if err := db.Where("address = ?", n.Address).Delete(&n).Error; err != nil {
		return fmt.Errorf("failed to delete node: %w", err)
	}

	return nil
}

// GetAllNodes returns a slice of Node records paginated by an offset, order
// and limit. An error is returned upon database query failure.
func GetAllNodes(db *gorm.DB, pq httputil.PaginationQuery) ([]Node, Paginator, error) {
	var (
		nodes []Node
		total int64
	)

	tx := db.Preload(clause.Associations)

	if err := tx.Scopes(paginateScope(pq, &nodes)).Error; err != nil {
		return nil, Paginator{}, fmt.Errorf("failed to query for nodes: %w", err)
	}

	if err := db.Model(&Node{}).Count(&total).Error; err != nil {
		return nil, Paginator{}, fmt.Errorf("failed to query for node count: %w", err)
	}

	return nodes, buildPaginator(pq, total), nil
}

// SearchNodes performs a paginated query for a set of Node records by moniker,
// network, version or location. If any empty query is provided, we return a
// paginated list of all Node records. Otherwise, if no matching Node records
// exist, an empty slice is returned.
func SearchNodes(db *gorm.DB, query string, pq httputil.PaginationQuery) ([]Node, Paginator, error) {
	if query == "" {
		return GetAllNodes(db, pq)
	}

	type queryRow struct {
		NodeID uint
	}

	rows, err := db.Raw(`SELECT DISTINCT
  ON (node_id) results.node_id AS node_id
FROM
  (
    SELECT
      n.id AS node_id,
      n.moniker,
      n.network,
      n.version,
      l.country,
      l.region,
      l.city
    FROM
      nodes n
      LEFT JOIN
        locations l
        ON (n.location_id = l.id)
    WHERE
      to_tsvector('english', COALESCE(n.moniker, '') || ' ' || COALESCE(n.network, '') || ' ' || COALESCE(n.version, '') || ' ' || COALESCE(l.country, '') || ' ' || COALESCE(l.region, '') || ' ' || COALESCE(l.city, '')) @@ websearch_to_tsquery('english', ?)
  )
  AS results;
`, query).Rows()
	if err != nil {
		return nil, Paginator{}, fmt.Errorf("failed to search for nodes: %w", err)
	}

	defer rows.Close()

	nodeIDs := []uint{}
	for rows.Next() {
		var qr queryRow
		if err := db.ScanRows(rows, &qr); err != nil {
			return nil, Paginator{}, fmt.Errorf("failed to search for nodes: %w", err)
		}

		nodeIDs = append(nodeIDs, qr.NodeID)
	}

	if len(nodeIDs) == 0 {
		return []Node{}, Paginator{}, nil
	}

	var nodes []Node

	if err := db.Preload(clause.Associations).
		Offset(int(offsetFromPage(pq))).
		Limit(int(pq.Limit)).
		Order(buildOrderBy(pq)).
		Find(&nodes, nodeIDs).Error; err != nil {
		return nil, Paginator{}, fmt.Errorf("failed to search for nodes: %w", err)
	}

	return nodes, buildPaginator(pq, int64(len(nodeIDs))), nil
}

// QueryNode performs a query for a Node record. The resulting record, if it
// exists, is returned. If the query fails or the record does not exist, an
// error is returned.
func QueryNode(db *gorm.DB, query map[string]interface{}) (Node, error) {
	var record Node

	if err := db.Preload(clause.Associations).Where(query).First(&record).Error; err != nil {
		return Node{}, fmt.Errorf("failed to query node: %w", err)
	}

	return record, nil
}

// GetStaleNodes returns all nodes that are stale. A node is considered stale if
// the updated_at timestamp is less than the provided time. An error is returned
// upon database failure.
func GetStaleNodes(db *gorm.DB, t time.Time) ([]Node, error) {
	var nodes []Node

	if err := db.Preload(clause.Associations).Where("updated_at < ?", t).Find(&nodes).Error; err != nil {
		return nil, err
	}

	return nodes, nil
}
