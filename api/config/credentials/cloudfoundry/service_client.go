package cloudfoundry

import (
	"encoding/json"
	"errors"
	"fmt"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

// ServiceClient TODO: documentation
type ServiceClient struct {
	client *rest.Client
}

// NewService creates a new Service Client
// baseURL should look like this: "https://siz65484.live.dynatrace.com/api/config/v1"
// token is an API Token
func NewService(baseURL string, token string) *ServiceClient {
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ServiceClient{client: client}
}

// Create TODO: documentation
func (cs *ServiceClient) Create(config *CloudFoundryCredentials) (*api.EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.POST("/cloudFoundry/credentials", config, 201); err != nil {
		return nil, err
	}
	var stub api.EntityShortRepresentation
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(id string, config *CloudFoundryCredentials) error {
	if _, err := cs.client.PUT(fmt.Sprintf("/cloudFoundry/credentials/%s", id), config, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the configuration to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/cloudFoundry/credentials/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*CloudFoundryCredentials, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the configuration to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/cloudFoundry/credentials/%s", id), 200); err != nil {
		return nil, err
	}
	var autoTag CloudFoundryCredentials
	if err = json.Unmarshal(bytes, &autoTag); err != nil {
		return nil, err
	}
	return &autoTag, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) ListAll() (*api.StubList, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/cloudFoundry/credentials", 200); err != nil {
		return nil, err
	}
	var stubList api.StubList
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}
