package managementzones

import (
	"encoding/json"
	"errors"
	"fmt"

	v2common "github.com/dtcookie/dynatrace/api/config/v2/common"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

const schemaVersion = "1.0.3"

type SettingsObject struct {
	SchemaVersion string    `json:"schemaVersion"`
	SchemaID      string    `json:"schemaId"`
	Scope         string    `json:"scope"`
	Value         *Settings `json:"value"`
}

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

// Create TODO: documentation
func (cs *ServiceClient) Create(item *Settings) (string, error) {
	payload := v2common.SettingsObjectCreate{
		Value:         item,
		SchemaID:      "builtin:management-zones",
		SchemaVersion: schemaVersion,
		Scope:         "environment",
	}

	post := cs.client.NewPOST("/settings/objects/", []*v2common.SettingsObjectCreate{&payload}).Expect(200)
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

// Update TODO: documentation
func (cs *ServiceClient) Update(id string, item *Settings) error {
	payload := v2common.SettingsObjectUpdate{
		Value:         item,
		SchemaVersion: schemaVersion,
	}
	if _, err := cs.client.PUT(fmt.Sprintf("/settings/objects/%s", id), &payload, 200); err != nil {
		return err
	}
	return nil
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

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*Settings, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the config to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/settings/objects/%s", id), 200); err != nil {
		return nil, err
	}
	var settingsObject SettingsObject
	if err = json.Unmarshal(bytes, &settingsObject); err != nil {
		return nil, err
	}
	return settingsObject.Value, nil
}

type SettingsObjectList struct {
	Items []*SettingsObjectListItem `json:"items"`
}

type SettingsObjectListItem struct {
	ObjectID string `json:"objectId"`
	Value    struct {
		Name string `json:"name"`
	} `json:"value"`
}

type Stub struct {
	ID   string
	Name string
}

// List TODO: documentation
func (cs *ServiceClient) List() ([]Stub, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/settings/objects?schemaIds=builtin%3Amanagement-zones&fields=objectId,value&pageSize=500", 200); err != nil {
		return nil, err
	}
	var sol SettingsObjectList
	if err = json.Unmarshal(bytes, &sol); err != nil {
		return nil, err
	}
	stubs := []Stub{}
	for _, stub := range sol.Items {
		stubs = append(stubs, Stub{ID: stub.ObjectID, Name: stub.Value.Name})
	}

	return stubs, nil
}

func (cs *ServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

func (cs *ServiceClient) LIST() ([]string, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/settings/objects?schemaIds=builtin%3Amanagement-zones&fields=objectId&pageSize=500", 200); err != nil {
		return nil, err
	}
	var sol v2common.SettingsObjectList
	if err = json.Unmarshal(bytes, &sol); err != nil {
		return nil, err
	}
	ids := []string{}
	for _, stub := range sol.Items {
		ids = append(ids, stub.ObjectID)
	}

	return ids, nil
}
