package client

import (
	"encoding/json"

	"github.com/noot/atomic-swap/common/rpcclient"
	"github.com/noot/atomic-swap/common/rpctypes"
)

// Query calls net_query.
func (c *Client) Query(maddr string) (*rpctypes.QueryPeerResponse, error) {
	const (
		method = "net_queryPeer"
	)

	req := &rpctypes.QueryPeerRequest{
		Multiaddr: maddr,
	}

	params, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := rpcclient.PostRPC(c.endpoint, method, string(params))
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, resp.Error
	}

	var res *rpctypes.QueryPeerResponse
	if err = json.Unmarshal(resp.Result, &res); err != nil {
		return nil, err
	}

	return res, nil
}
