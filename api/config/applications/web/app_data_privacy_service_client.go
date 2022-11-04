package web

import (
	api "github.com/dtcookie/dynatrace/api/config"
)

// ServiceClient TODO: documentation
type AppDataPrivacyServiceClient struct {
	client *ServiceClient
}

// NewService creates a new Service Client
// baseURL should look like this: "https://siz65484.live.dynatrace.com/api/config/v1"
// token is an API Token
func NewAppDataPrivacyService(baseURL string, token string) *AppDataPrivacyServiceClient {
	return &AppDataPrivacyServiceClient{client: NewService(baseURL, token)}
}

func (cs *AppDataPrivacyServiceClient) Get(id string) (*ApplicationDataPrivacy, error) {
	return cs.client.GetAppDataPrivacy(id)
}

func (cs *AppDataPrivacyServiceClient) Store(config *ApplicationDataPrivacy) error {
	return cs.client.StoreAppDataPrivacy(config)
}

func (cs *AppDataPrivacyServiceClient) List() (*api.StubList, error) {
	return cs.client.List()
}

func (cs *AppDataPrivacyServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

func (cs *AppDataPrivacyServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if stubList, err := cs.List(); err == nil {
		for _, stub := range stubList.Values {
			ids = append(ids, stub.ID)
		}
	}
	return ids, err
}
