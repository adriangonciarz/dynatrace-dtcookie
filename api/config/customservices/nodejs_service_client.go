package customservices

import (
	api "github.com/dtcookie/dynatrace/api/config"
)

type NodeJSServiceClient struct {
	client *ServiceClient
}

func NewNodeJSService(baseURL string, token string) *NodeJSServiceClient {
	return &NodeJSServiceClient{client: NewService(baseURL, token)}
}

func (cs *NodeJSServiceClient) Create(customService *CustomService, technology Technology) (*api.EntityShortRepresentation, error) {
	return cs.client.Create(customService, Technologies.NodeJS)
}

func (cs *NodeJSServiceClient) Update(customService *CustomService, technology Technology) error {
	return cs.client.Update(customService, Technologies.NodeJS)
}

func (cs *NodeJSServiceClient) Delete(id string, technology Technology) error {
	return cs.client.Delete(id, Technologies.NodeJS)
}

func (cs *NodeJSServiceClient) Get(id string, includeProcessGroupReferences bool) (*CustomService, error) {
	var err error
	if customService, err := cs.client.Get(id, Technologies.NodeJS, includeProcessGroupReferences); err == nil {
		customService.Technology = Technologies.NodeJS
		return customService, nil
	}
	return nil, err
}

func (cs *NodeJSServiceClient) List() (*api.StubList, error) {
	return cs.client.List(Technologies.NodeJS)
}

func (cs *NodeJSServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id, false)
}

func (cs *NodeJSServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if stubList, err := cs.List(); err == nil {
		for _, stub := range stubList.Values {
			ids = append(ids, stub.ID)
		}
	}
	return ids, err
}
