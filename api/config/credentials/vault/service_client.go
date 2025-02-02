package vault

import (
	"encoding/json"
	"errors"
	"fmt"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
	"github.com/dtcookie/opt"
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
func (cs *ServiceClient) Create(config *Credentials) (*api.EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if len(opt.String(config.ID)) > 0 {
		return nil, errors.New("you must not provide an ID within the configuration payload upon creation")
	}

	if bytes, err = cs.client.POST("/credentials", config, 201); err != nil {
		return nil, err
	}
	var stub api.EntityShortRepresentation
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(config *Credentials) error {
	if len(opt.String(config.ID)) == 0 {
		return errors.New("the configuration doesn't contain an ID")
	}
	if _, err := cs.client.PUT(fmt.Sprintf("/credentials/%s", opt.String(config.ID)), config, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the configuration to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/credentials/%s", id), 204); err != nil {
		return err
	}
	return nil
}

func (cs *ServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*Credentials, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the configuration to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/credentials/%s", id), 200); err != nil {
		return nil, err
	}
	var credentials Credentials
	if err = json.Unmarshal(bytes, &credentials); err != nil {
		return nil, err
	}
	return &credentials, nil
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
	for _, credential := range list.Credentials {
		res = append(res, *credential.ID)
	}
	return res, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) ListAll() (*CredentialsList, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/credentials", 200); err != nil {
		return nil, err
	}
	var stubList CredentialsList
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}
