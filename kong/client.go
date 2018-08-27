package kong

import (
	"context"
	"net/http"

	"github.com/rocketspacer/go-kong-client/kong/dto"
)

// Client represent a Kong Admin API client
type Client interface {
	/*******************************************************************************
	* Node API                                                                     *
	* https://docs.konghq.com/0.14.x/admin-api/#information-routes                 *
	*******************************************************************************/

	// https://docs.konghq.com/0.14.x/admin-api/#retrieve-node-information
	GetNodeInfo(ctx context.Context) (*dto.NodeInfo, error)

	// https://docs.konghq.com/0.14.x/admin-api/#retrieve-node-status
	GetNodeStatus(ctx context.Context) (*dto.NodeStatus, error)

	/*******************************************************************************
	* Service API                                                                  *
	* https://docs.konghq.com/0.14.x/admin-api/#service-object                     *
	*******************************************************************************/

	// https://docs.konghq.com/0.14.x/admin-api/#add-service
	AddService(ctx context.Context, svc *dto.Service) (*dto.Service, error)

	/*******************************************************************************
	* Route API                                                                    *
	*******************************************************************************/
}

type clientImpl struct {
	config     *Config
	httpClient *http.Client
}

// NewClient create a new concrete implementation of the Client interface
func NewClient(config *Config) (Client, error) {
	return NewClientWithHTTPClient(config, http.DefaultClient)
}

// NewClientWithHTTPClient create a new concrete implementation of the Client interface with a custom http.Client
func NewClientWithHTTPClient(config *Config, httpClient *http.Client) (Client, error) {
	if config == nil {
		return nil, ErrMissingConfig
	}

	if httpClient == nil {
		return nil, ErrMissingHTTPClient
	}

	return &clientImpl{
		config:     config,
		httpClient: httpClient,
	}, nil
}
