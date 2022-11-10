package alerting

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

const schemaVersion = "7.0.6"

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
func (cs *ServiceClient) Create(item *Profile) (string, error) {
	payload := SettingsObjectCreate{
		Value:         item,
		SchemaID:      "builtin:alerting.profile",
		SchemaVersion: schemaVersion,
		Scope:         "environment",
	}

	post := cs.client.NewPOST("/settings/objects/", []*SettingsObjectCreate{&payload}).Expect(200)
	if data, err := post.Send(); err == nil {
		var sor []SettingsObjectResponse
		if err := json.Unmarshal(data, &sor); err != nil {
			return "", err
		}
		return sor[0].ObjectID, nil
	} else if data != nil {
		var soer []SettingsObjectErrorResponse
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
func (cs *ServiceClient) Update(id string, item *Profile) error {
	payload := SettingsObjectUpdate{
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
func (cs *ServiceClient) Get(id string) (*Profile, error) {
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
	settingsObject.Value.ID = id

	return settingsObject.Value, nil
}

// List TODO: documentation
func (cs *ServiceClient) List() ([]string, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/settings/objects?schemaIds=builtin%3Aalerting.profile&fields=objectId&pageSize=500", 200); err != nil {
		return nil, err
	}
	var sol SettingsObjectList
	if err = json.Unmarshal(bytes, &sol); err != nil {
		return nil, err
	}
	ids := []string{}
	for _, stub := range sol.Items {
		ids = append(ids, stub.ObjectID)
	}

	return ids, nil
}

func (cs *ServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

func (cs *ServiceClient) LIST() ([]string, error) {
	return cs.List()
}

func (cs *ServiceClient) ListInterface() (interface{}, error) {
	var ids []string
	var err error
	if ids, err = cs.List(); err != nil {
		return nil, err
	}

	var profile *Profile
	var profiles []*Profile
	for _, id := range ids {
		if profile, err = cs.Get(id); err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}
	return profiles, err
}

type SettingsObjectList struct {
	Items []*SettingsObjectListItem `json:"items"`
}

type SettingsObjectListItem struct {
	ObjectID string `json:"objectId"`
}

type SettingsObject struct {
	SchemaVersion string   `json:"schemaVersion"`
	SchemaID      string   `json:"schemaId"`
	Scope         string   `json:"scope"`
	Value         *Profile `json:"value"`
}

type SettingsObjectUpdate struct {
	SchemaVersion string      `json:"schemaVersion"`
	Value         interface{} `json:"value"`
}

type SettingsObjectCreate struct {
	SchemaVersion string      `json:"schemaVersion"`
	SchemaID      string      `json:"schemaId"`
	Scope         string      `json:"scope"`
	Value         interface{} `json:"value"`
}

type SettingsObjectResponse struct {
	ObjectID string `json:"objectId"` // The ID of the settings object
	Code     int32  // The HTTP status code for the object
}

type SettingsObjectErrorResponse struct {
	InvalidValue map[string]interface{} `json:"invalidValue,omitempty"` // The value of the setting. \n\n It defines the actual values of settings' parameters. \n\nThe actual content depends on the object's schema.
	Error        *Error                 `json:"error,omitempty"`        // Error details
	Code         *int32                 `json:"code,omitempty"`         // The HTTP status code for the object
}

type Error struct {
	ConstraintViolations []*ConstraintViolation `json:"constraintViolations,omitempty"` // A list of constraint violations
	Message              string                 `json:"message,omitempty"`              // The error message
	Code                 int32                  `json:"code,omitempty"`                 // The HTTP status code
}

type ConstraintViolation struct {
	ParmeterLocation *ParameterLocation `json:"parameterLocation,omitempty"`
	Location         *string            `json:"location,omitempty"`
	Message          *string            `json:"message,omitempty"`
	Path             *string            `json:"path,omitempty"`
}

type ParameterLocation string

var ParameterLocations = struct {
	Path        ParameterLocation
	PayloadBody ParameterLocation
	Query       ParameterLocation
}{
	Path:        ParameterLocation("PATH"),
	PayloadBody: ParameterLocation("PAYLOAD_BODY"),
	Query:       ParameterLocation("QUERY"),
}
