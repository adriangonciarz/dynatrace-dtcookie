package customservices

import (
	api "github.com/dtcookie/dynatrace/api/config"
)

type DotNetServiceClient struct {
	client *ServiceClient
}

func NewDotNetService(baseURL string, token string) *DotNetServiceClient {
	return &DotNetServiceClient{client: NewService(baseURL, token)}
}

func (cs *DotNetServiceClient) Create(customService *CustomService, technology Technology) (*api.EntityShortRepresentation, error) {
	return cs.client.Create(customService, Technologies.DotNet)
}

func (cs *DotNetServiceClient) Update(customService *CustomService, technology Technology) error {
	return cs.client.Update(customService, Technologies.DotNet)
}

func (cs *DotNetServiceClient) Delete(id string, technology Technology) error {
	return cs.client.Delete(id, Technologies.DotNet)
}

func (cs *DotNetServiceClient) Get(id string, includeProcessGroupReferences bool) (*CustomService, error) {
	var err error
	if customService, err := cs.client.Get(id, Technologies.DotNet, includeProcessGroupReferences); err == nil {
		customService.Technology = Technologies.DotNet
		return customService, nil
	}
	return nil, err
}

func (cs *DotNetServiceClient) List() (*api.StubList, error) {
	return cs.client.List(Technologies.DotNet)
}

func (cs *DotNetServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id, false)
}

func (cs *DotNetServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if stubList, err := cs.List(); err == nil {
		for _, stub := range stubList.Values {
			ids = append(ids, stub.ID)
		}
	}
	return ids, err
}
