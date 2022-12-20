package apitokens

import (
	"encoding/json"
	"errors"
	"fmt"

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
func (cs *ServiceClient) Create(item *ApiToken) (*ApiToken, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.POST("/apiTokens", item, 201); err != nil {
		return nil, err
	}
	var resultToken ApiToken
	if err = json.Unmarshal(bytes, &resultToken); err != nil {
		return nil, err
	}
	if item.Enabled == nil || *item.Enabled == false {
		item.Enabled = opt.NewBool(false)
		if _, err = cs.client.PUT(fmt.Sprintf("/apiTokens/%s", *resultToken.ID), item, 204); err != nil {
			return nil, err
		}
		resultToken.Enabled = item.Enabled
	}
	resultToken.Name = item.Name
	resultToken.Scopes = item.Scopes
	return &resultToken, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(id string, item *ApiToken) error {
	if _, err := cs.client.PUT(fmt.Sprintf("/apiTokens/%s", id), item, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the MaintenanceWindow to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/apiTokens/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*ApiToken, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the MaintenanceWindow to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/apiTokens/%s", id), 200); err != nil {
		return nil, err
	}
	var item ApiToken
	if err = json.Unmarshal(bytes, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) List() (*TokenList, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/apiTokens?pageSize=10000&sort=-creationDate", 200); err != nil {
		return nil, err
	}
	var tokenList TokenList
	if err = json.Unmarshal(bytes, &tokenList); err != nil {
		return nil, err
	}
	return &tokenList, nil
}

func (cs *ServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

func (cs *ServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if tokenList, err := cs.List(); err == nil {
		for _, token := range tokenList.ApiTokens {
			ids = append(ids, *token.ID)
		}
	}
	return ids, err
}
