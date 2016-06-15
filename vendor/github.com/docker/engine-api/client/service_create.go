package client

import (
	"encoding/json"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"golang.org/x/net/context"
)

// ServiceCreate creates a new Service.
func (cli *Client) ServiceCreate(ctx context.Context, service swarm.ServiceSpec) (types.ServiceCreateResponse, error) {
	var response types.ServiceCreateResponse
	resp, err := cli.post(ctx, "/services/create", nil, service, nil)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(resp.body).Decode(&response)
	ensureReaderClosed(resp)
	return response, err
}
