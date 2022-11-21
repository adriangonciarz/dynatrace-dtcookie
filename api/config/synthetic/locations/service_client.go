package locations

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
	"github.com/dtcookie/opt"
)

// ServiceClient TODO: documentation
type ServiceClient struct {
	client *rest.Client
}

// NewService creates a new Service Client
// baseURL should look like this: "https://siz65484.live.dynatrace.com/api/v1"
// token is an API Token
func NewService(baseURL string, token string) *ServiceClient {
	baseURL = strings.ReplaceAll(baseURL, "/api/config/v1", "/api/v1")
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ServiceClient{client: client}
}

// Create TODO: documentation
func (cs *ServiceClient) Create(config *PrivateSyntheticLocation) (string, error) {
	var err error
	var bytes []byte

	if len(opt.String(config.ID)) > 0 {
		return "", errors.New("you must not provide an ID within the configuration payload upon creation")
	}

	if bytes, err = cs.client.POST("/synthetic/locations", config, 200); err != nil {
		return "", err
	}

	var stub struct {
		ID string `json:"entityId"`
	}

	if err = json.Unmarshal(bytes, &stub); err != nil {
		return "", err
	}
	return stub.ID, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(config *PrivateSyntheticLocation) error {
	if len(opt.String(config.ID)) == 0 {
		return errors.New("the configuration doesn't contain an ID")
	}
	if _, err := cs.client.PUT(fmt.Sprintf("/synthetic/locations/%s", opt.String(config.ID)), config, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the configuration to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("synthetic/locations/%s", id), 204); err != nil {
		return err
	}
	return nil
}

func (cs *ServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*PrivateSyntheticLocation, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the configuration to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/synthetic/locations/%s", id), 200); err != nil {
		return nil, err
	}
	var location PrivateSyntheticLocation
	if err = json.Unmarshal(bytes, &location); err != nil {
		return nil, err
	}
	return &location, nil
}

func (cs *ServiceClient) ListInterface() (interface{}, error) {
	credentialList, err := cs.ListAll()
	if err != nil {
		return nil, err
	}
	return credentialList, nil
}

func (cs *ServiceClient) LIST() ([]string, error) {
	list, err := cs.ListAll()
	if err != nil {
		return nil, err
	}
	res := []string{}
	for _, location := range list.Locations {
		res = append(res, *location.ID)
	}
	return res, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) ListAll() (*SyntheticLocations, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/synthetic/locations?type=PRIVATE", 200); err != nil {
		return nil, err
	}
	var stubList SyntheticLocations
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}
