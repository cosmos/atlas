package crawl

import (
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/harwoeck/ipstack"
	tmrpchttp "github.com/tendermint/tendermint/rpc/client/http"
	jsonrpcclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"

	"github.com/cosmos/atlas/server/models"
)

var clientTimeout = 5 * time.Second

func newRPCClient(remote string, timeout time.Duration) (*tmrpchttp.HTTP, error) {
	httpClient, err := jsonrpcclient.DefaultHTTPClient(remote)
	if err != nil {
		return nil, err
	}

	httpClient.Timeout = timeout
	return tmrpchttp.NewWithClient(remote, "/websocket", httpClient)
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
