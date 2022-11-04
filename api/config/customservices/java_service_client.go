package customservices

import (
	api "github.com/dtcookie/dynatrace/api/config"
)

type JavaServiceClient struct {
	client *ServiceClient
}

func NewJavaService(baseURL string, token string) *JavaServiceClient {
	return &JavaServiceClient{client: NewService(baseURL, token)}
}

func (cs *JavaServiceClient) Create(customService *CustomService, technology Technology) (*api.EntityShortRepresentation, error) {
	return cs.client.Create(customService, Technologies.Java)
}

func (cs *JavaServiceClient) Update(customService *CustomService, technology Technology) error {
	return cs.client.Update(customService, Technologies.Java)
}

func (cs *JavaServiceClient) Delete(id string, technology Technology) error {
	return cs.client.Delete(id, Technologies.Java)
}

func (cs *JavaServiceClient) Get(id string, includeProcessGroupReferences bool) (*CustomService, error) {
	var err error
	if customService, err := cs.client.Get(id, Technologies.Java, includeProcessGroupReferences); err == nil {
		customService.Technology = Technologies.Java
		return customService, nil
	}
	return nil, err
}

func (cs *JavaServiceClient) List() (*api.StubList, error) {
	return cs.client.List(Technologies.Java)
}

func (cs *JavaServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id, false)
}

func (cs *JavaServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if stubList, err := cs.List(); err == nil {
		for _, stub := range stubList.Values {
			ids = append(ids, stub.ID)
		}
	}
	return ids, err
}
