package customservices

import (
	api "github.com/dtcookie/dynatrace/api/config"
)

type GoServiceClient struct {
	client *ServiceClient
}

func NewGoService(baseURL string, token string) *GoServiceClient {
	return &GoServiceClient{client: NewService(baseURL, token)}
}

func (cs *GoServiceClient) Create(customService *CustomService, technology Technology) (*api.EntityShortRepresentation, error) {
	return cs.client.Create(customService, Technologies.Go)
}

func (cs *GoServiceClient) Update(customService *CustomService, technology Technology) error {
	return cs.client.Update(customService, Technologies.Go)
}

func (cs *GoServiceClient) Delete(id string, technology Technology) error {
	return cs.client.Delete(id, Technologies.Go)
}

func (cs *GoServiceClient) Get(id string, includeProcessGroupReferences bool) (*CustomService, error) {
	var err error
	if customService, err := cs.client.Get(id, Technologies.Go, includeProcessGroupReferences); err == nil {
		return customService, nil
	}
	return nil, err
}

func (cs *GoServiceClient) List() (*api.StubList, error) {
	return cs.client.List(Technologies.Go)
}

func (cs *GoServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id, false)
}

func (cs *GoServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if stubList, err := cs.List(); err == nil {
		for _, stub := range stubList.Values {
			ids = append(ids, stub.ID)
		}
	}
	return ids, err
}
