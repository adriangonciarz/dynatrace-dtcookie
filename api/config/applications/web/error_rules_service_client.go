package web

import (
	api "github.com/dtcookie/dynatrace/api/config"
)

// ServiceClient TODO: documentation
type ErrorRulesServiceClient struct {
	client *ServiceClient
}

// NewService creates a new Service Client
// baseURL should look like this: "https://siz65484.live.dynatrace.com/api/config/v1"
// token is an API Token
func NewErrorRulesService(baseURL string, token string) *ErrorRulesServiceClient {
	return &ErrorRulesServiceClient{client: NewService(baseURL, token)}
}

func (cs *ErrorRulesServiceClient) Get(id string) (*ApplicationErrorRules, error) {
	return cs.client.GetErrorRules(id)
}

func (cs *ErrorRulesServiceClient) Store(config *ApplicationErrorRules) error {
	return cs.client.StoreErrorRules(config)
}

func (cs *ErrorRulesServiceClient) List() (*api.StubList, error) {
	return cs.client.List()
}

func (cs *ErrorRulesServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

func (cs *ErrorRulesServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if stubList, err := cs.List(); err == nil {
		for _, stub := range stubList.Values {
			ids = append(ids, stub.ID)
		}
	}
	return ids, err
}
