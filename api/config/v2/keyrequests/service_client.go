package keyrequests

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"

	v2common "github.com/dtcookie/dynatrace/api/config/v2/common"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

// const schemaVersion = "0.1.1"

// ServiceClient TODO: documentation
type ServiceClient struct {
	client *rest.Client
}

// NewService creates a new Service Client
// baseURL should look like this: "https://siz65484.live.dynatrace.com/api/config/v2"
// token is an API Token
func NewService(baseURL string, token string) *ServiceClient {
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ServiceClient{client: client}
}

type ListResponse struct {
	Items []struct {
		ObjectID string `json:"objectId"`
		Scope    string `json:"scope"`
		Value    struct {
			KeyRequestNames []string `json:"keyRequestNames"`
		} `json:"value"`
	} `json:"items"`
}

func (cs *ServiceClient) List(serviceID string) (string, *KeyRequest, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/settings/objects?schemaIds=builtin:settings.subscriptions.service&scopes=%s&fields=objectId,value", serviceID), 200); err != nil {
		return "", nil, err
	}
	var listResponse ListResponse
	if err = json.Unmarshal(bytes, &listResponse); err != nil {
		return "", nil, err
	}
	if len(listResponse.Items) == 0 {
		return "", nil, nil
	}
	result := listResponse.Items[0].Value.KeyRequestNames
	sort.Strings(result)
	return listResponse.Items[0].ObjectID, &KeyRequest{ServiceID: serviceID, Names: result}, nil
}

func (cs *ServiceClient) Get(ID string) (*KeyRequest, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/settings/objects/%s", ID), 200); err != nil {
		return nil, err
	}
	response := struct {
		ObjectID string `json:"objectId"`
		Scope    string `json:"scope"`
		Value    struct {
			KeyRequestNames []string `json:"keyRequestNames"`
		} `json:"value"`
	}{}
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	result := response.Value.KeyRequestNames
	sort.Strings(result)
	return &KeyRequest{ServiceID: response.Scope, Names: result}, nil
}

func (cs *ServiceClient) Update(keyRequest *KeyRequest) error {
	_, err := cs.Create(keyRequest)
	return err
}

func (cs *ServiceClient) Create(keyRequest *KeyRequest) (string, error) {
	payLoad := map[string]interface{}{
		"schemaVersion": "0.1.1",
		"schemaId":      "builtin:settings.subscriptions.service",
		"value": map[string][]string{
			"keyRequestNames": keyRequest.Names,
		},
		// "scope": "SERVICE-FBD5DB17596B3215",
		"scope": keyRequest.ServiceID,
	}

	post := cs.client.NewPOST("/settings/objects/", []interface{}{&payLoad}).Expect(200)
	if data, err := post.Send(); err == nil {
		var sor []v2common.SettingsObjectResponse
		if err := json.Unmarshal(data, &sor); err != nil {
			return "", err
		}
		return sor[0].ObjectID, nil
	} else if data != nil {
		var soer []v2common.SettingsObjectErrorResponse
		if err := json.Unmarshal(data, &soer); err == nil {
			od, _ := json.Marshal(soer[0])
			return "", errors.New(string(od))
		}
		return "", err
	} else {
		return "", err
	}
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the item to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/settings/objects/%s", id), 204); err != nil {
		return err
	}
	return nil
}

func (cs *ServiceClient) LIST(serviceID string) (string, interface{}, error) {
	return cs.List(serviceID)
}
