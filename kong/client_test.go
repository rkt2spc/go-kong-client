package kong

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*******************************************************************************
*                               [TEST] NewClient                               *
*******************************************************************************/

// Test scenario schema
type NewClientScenario struct {
	Description    string
	Config         *Config
	ExpectedClient Client
	ExpectedError  error
}

// Test scenarios
var NewClientScenarios = []NewClientScenario{
	NewClientMissingConfigScenario(),
	NewClientSuccessScenario(),
}

// Run test for each scenario
func TestNewClient(t *testing.T) {
	t.Parallel()
	for _, sc := range NewClientScenarios {
		scenario := sc
		t.Run(scenario.Description, func(t *testing.T) {
			client, err := NewClient(scenario.Config)
			assert.Equal(t, scenario.ExpectedClient, client)
			assert.Equal(t, scenario.ExpectedError, err)
		})
	}
}

////////////////////////// TEST SCENARIOS DEFINITIONS //////////////////////////

func NewClientMissingConfigScenario() NewClientScenario {
	return NewClientScenario{
		Description:    "NewClient with missing config",
		Config:         nil,
		ExpectedClient: nil,
		ExpectedError:  ErrMissingConfig,
	}
}

func NewClientSuccessScenario() NewClientScenario {
	config := &Config{AdminAPIUrl: "http://dummy-kong-admin"}

	return NewClientScenario{
		Description:    "NewClient success",
		Config:         config,
		ExpectedClient: &clientImpl{config: config, httpClient: http.DefaultClient},
		ExpectedError:  nil,
	}
}

/*******************************************************************************
*                        [TEST] NewClientWithHTTPClient                        *
*******************************************************************************/

// Test scenario schema
type NewClientWithHTTPClientScenario struct {
	Description    string
	Config         *Config
	HTTPClient     *http.Client
	ExpectedClient Client
	ExpectedError  error
}

// Test scenarios
var NewClientWithHTTPClientScenarios = []NewClientWithHTTPClientScenario{
	NewClientWithHTTPClientMissingConfigScenario(),
	NewClientWithHTTPClientMissingHTTPClientScenario(),
	NewClientWithHTTPClientSuccessScenario(),
}

// Run test for each scenario
func TestNewClientWithHTTPClient(t *testing.T) {
	t.Parallel()
	for _, sc := range NewClientWithHTTPClientScenarios {
		scenario := sc
		t.Run(scenario.Description, func(t *testing.T) {
			client, err := NewClientWithHTTPClient(scenario.Config, scenario.HTTPClient)
			assert.Equal(t, scenario.ExpectedClient, client)
			assert.Equal(t, scenario.ExpectedError, err)
		})
	}
}

////////////////////////// TEST SCENARIOS DEFINITIONS //////////////////////////

func NewClientWithHTTPClientMissingConfigScenario() NewClientWithHTTPClientScenario {
	return NewClientWithHTTPClientScenario{
		Description:    "NewClientWithHTTPClient with missing config",
		Config:         nil,
		HTTPClient:     nil,
		ExpectedClient: nil,
		ExpectedError:  ErrMissingConfig,
	}
}

func NewClientWithHTTPClientMissingHTTPClientScenario() NewClientWithHTTPClientScenario {
	return NewClientWithHTTPClientScenario{
		Description:    "NewClientWithHTTPClient with missing http-client",
		Config:         &Config{AdminAPIUrl: "http://dummy-kong-admin"},
		HTTPClient:     nil,
		ExpectedClient: nil,
		ExpectedError:  ErrMissingHTTPClient,
	}
}

func NewClientWithHTTPClientSuccessScenario() NewClientWithHTTPClientScenario {
	config := &Config{AdminAPIUrl: "http://dummy-kong-admin"}
	httpClient := http.DefaultClient

	return NewClientWithHTTPClientScenario{
		Description:    "NewClient success",
		Config:         config,
		HTTPClient:     httpClient,
		ExpectedClient: &clientImpl{config: config, httpClient: httpClient},
		ExpectedError:  nil,
	}
}
