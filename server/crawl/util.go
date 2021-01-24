package crawl

import (
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/harwoeck/ipstack"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	libclient "github.com/tendermint/tendermint/rpc/lib/client"

	"github.com/cosmos/atlas/server/models"
)

var clientTimeout = 15 * time.Second

func newRPCClient(remote string, timeout time.Duration) *rpcclient.HTTP {
	httpClient := libclient.DefaultHTTPClient(remote)
	httpClient.Timeout = timeout
	return rpcclient.NewHTTPWithClient(remote, "/websocket", httpClient)
}

func parsePort(nodeAddr string) string {
	u, err := url.Parse(nodeAddr)
	if err != nil {
		return ""
	}

	return u.Port()
}

func parseHostname(nodeAddr string) string {
	u, err := url.Parse(nodeAddr)
	if err != nil {
		return ""
	}

	return u.Hostname()
}

func locationFromIPResp(r *ipstack.Response) models.Location {
	return models.Location{
		Country:   r.CountryName,
		Region:    r.RegionName,
		City:      r.City,
		Latitude:  fmt.Sprintf("%f", r.Latitude),
		Longitude: fmt.Sprintf("%f", r.Longitude),
	}
}

func pingAddress(address string, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}

	defer conn.Close()
	return true
}
