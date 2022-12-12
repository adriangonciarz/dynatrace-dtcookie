package networkzones

import (
	"encoding/json"
	"fmt"

	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
	"github.com/google/uuid"
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
func (cs *ServiceClient) Create(nz *NetworkZone) (*NetworkZone, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.PUT(fmt.Sprintf("/networkZones/%s", uuid.New().String()), nz, 201); err != nil {
		return nil, err
	}
	var networkZone NetworkZone
	if err = json.Unmarshal(bytes, &networkZone); err != nil {
		return nil, err
	}
	return &networkZone, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(id string, nz *NetworkZone) error {
	if _, err := cs.client.PUT(fmt.Sprintf("/networkZones/%s", id), nz, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if _, err := cs.client.DELETE(fmt.Sprintf("/networkZones/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*NetworkZone, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/networkZones/%s", id), 200); err != nil {
		return nil, err
	}
	var nz NetworkZone
	if err = json.Unmarshal(bytes, &nz); err != nil {
		return nil, err
	}
	return &nz, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) List() (*NetworkZones, error) {
	var err error
	var bytes []byte
	if bytes, err = cs.client.GET("/networkZones", 200); err != nil {
		return nil, err
	}
	var networkZones *NetworkZones
	if err = json.Unmarshal(bytes, &networkZones); err != nil {
		return nil, err
	}
	return networkZones, nil
}

func (cs *ServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

func (cs *ServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if networkZones, err := cs.List(); err == nil {
		for _, nz := range networkZones.Zones {
			ids = append(ids, *nz.ID)
		}
	}
	return ids, err
}
