package requestattributes

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
func (cs *ServiceClient) Create(item *RequestAttribute) (*api.EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if len(opt.String(item.ID)) > 0 {
		return nil, errors.New("you MUST NOT provide an ID within the payload upon creation")
	}

	if bytes, err = cs.client.POST("/service/requestAttributes", item, 201); err != nil {
		return nil, err
	}
	var stub api.EntityShortRepresentation
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(item *RequestAttribute) error {
	if len(opt.String(item.ID)) == 0 {
		return errors.New("the item doesn't contain an ID")
	}
	if _, err := cs.client.PUT(fmt.Sprintf("/service/requestAttributes/%s", opt.String(item.ID)), item, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the item to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/service/requestAttributes/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*RequestAttribute, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the MaintenanceWindow to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/service/requestAttributes/%s", id), 200); err != nil {
		return nil, err
	}
	var item RequestAttribute
	if err = json.Unmarshal(bytes, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) ListAll() (*api.StubList, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/service/requestAttributes", 200); err != nil {
		return nil, err
	}
	var stubList api.StubList
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}

func (cs *ServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

func (cs *ServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if stubList, err := cs.ListAll(); err == nil {
		for _, stub := range stubList.Values {
			ids = append(ids, stub.ID)
		}
	}
	return ids, err
}

func (cs *ServiceClient) ListInterface() (interface{}, error) {
	return cs.ListAll()
}
