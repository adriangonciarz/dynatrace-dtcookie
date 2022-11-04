package service

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

// ServiceClient holds REST API connection details
type ServiceClient struct {
	client *rest.Client
}

// NewService creates a new Service Client
// baseURL should look like this: "https://siz65484.live.dynatrace.com/api/v1"
// token is an API Token
func NewService(baseURL string, token string) *ServiceClient {
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ServiceClient{client: client}
}

// List function retrieves the list of all available services in the environment
func (cs *ServiceClient) List() (Services, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/entity/services", 200); err != nil {
		return nil, err
	}
	var serviceList Services
	if err = json.Unmarshal(bytes, &serviceList); err != nil {
		return nil, err
	}

	return serviceList, nil
}

func (cs *ServiceClient) ListInterface() (interface{}, error) {
	return cs.List()
}