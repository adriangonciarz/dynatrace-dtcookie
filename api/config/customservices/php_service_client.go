package customservices

import (
	api "github.com/dtcookie/dynatrace/api/config"
)

type PHPServiceClient struct {
	client *ServiceClient
}

func NewPHPService(baseURL string, token string) *PHPServiceClient {
	return &PHPServiceClient{client: NewService(baseURL, token)}
}

func (cs *PHPServiceClient) Create(customService *CustomService, technology Technology) (*api.EntityShortRepresentation, error) {
	return cs.client.Create(customService, Technologies.PHP)
}

func (cs *PHPServiceClient) Update(customService *CustomService, technology Technology) error {
	return cs.client.Update(customService, Technologies.PHP)
}

func (cs *PHPServiceClient) Delete(id string, technology Technology) error {
	return cs.client.Delete(id, Technologies.PHP)
}

func (cs *PHPServiceClient) Get(id string, includeProcessGroupReferences bool) (*CustomService, error) {
	var err error
	if customService, err := cs.client.Get(id, Technologies.PHP, includeProcessGroupReferences); err == nil {
		return customService, nil
	}
	return nil, err
}

func (cs *PHPServiceClient) List() (*api.StubList, error) {
	return cs.client.List(Technologies.PHP)
}

func (cs *PHPServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id, false)
}

func (cs *PHPServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if stubList, err := cs.List(); err == nil {
		for _, stub := range stubList.Values {
			ids = append(ids, stub.ID)
		}
	}
	return ids, err
}
