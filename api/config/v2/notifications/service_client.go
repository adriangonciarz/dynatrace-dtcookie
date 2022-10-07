package notifications

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	v2common "github.com/dtcookie/dynatrace/api/config/v2/common"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

const schemaVersion = "1.2"

type SettingsObject struct {
	SchemaVersion string        `json:"schemaVersion"`
	SchemaID      string        `json:"schemaId"`
	Scope         string        `json:"scope"`
	Value         *Notification `json:"value"`
}

// ServiceClient TODO: documentation
type ServiceClient struct {
	client   *rest.Client
	schemaID string
}

type SettingsObjectCreate struct {
	SchemaVersion string        `json:"schemaVersion"`
	SchemaID      string        `json:"schemaId"`
	Scope         string        `json:"scope"`
	Value         *Notification `json:"value"`
}

// NewService creates a new Service Client
// baseURL should look like this: "https://######.live.dynatrace.com/api/v2"
// token is an API Token
func NewService(baseURL string, token string) *ServiceClient {
	credentials := credentials.New(token)
	config := rest.Config{}
	if !strings.HasSuffix(baseURL, "/api/v2") {
		if strings.HasSuffix(baseURL, "/") {
			baseURL = baseURL + "api/v2"
		} else {
			baseURL = baseURL + "/api/v2"
		}

	}
	client := rest.NewClient(&config, baseURL, credentials)
	return &ServiceClient{client: client, schemaID: "builtin:problem.notifications"}
}

// Create TODO: documentation
func (cs *ServiceClient) Create(config *Notification) (string, error) {
	payload := SettingsObjectCreate{
		Value:         config,
		SchemaID:      cs.schemaID,
		SchemaVersion: schemaVersion,
		Scope:         "environment",
	}

	post := cs.client.NewPOST("/settings/objects/", []*SettingsObjectCreate{&payload}).Expect(200)
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
func (cs *ServiceClient) Update(id string, config *Notification) error {
	payload := v2common.SettingsObjectUpdate{
		Value:         config,
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
func (cs *ServiceClient) Get(id string, target *Notification) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the config to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/settings/objects/%s", id), 200); err != nil {
		return err
	}
	var settingsObject SettingsObject
	settingsObject.Value = target
	if err = json.Unmarshal(bytes, &settingsObject); err != nil {
		return err
	}
	// settingsObject.Value.AnsibleTower.Enabled = settingsObject.Value.Enabled
	// settingsObject.Value.AnsibleTower.Name = settingsObject.Value.Name
	// settingsObject.Value.AnsibleTower.ProfileID = settingsObject.Value.AlertingProfile
	return nil
}

// List TODO: documentation
// schemaID builtin:problem.notifications
// pageSize 500
func (cs *ServiceClient) List(pageSize ...int) ([]string, error) {
	var err error
	var bytes []byte

	ps := 500
	if len(pageSize) > 0 {
		ps = pageSize[0]
	}

	if bytes, err = cs.client.GET(fmt.Sprintf("/settings/objects?schemaIds=%s&fields=objectId&pageSize=%d", url.QueryEscape(cs.schemaID), ps), 200); err != nil {
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
