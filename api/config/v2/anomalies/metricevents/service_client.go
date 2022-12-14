package metricevents

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	v2common "github.com/dtcookie/dynatrace/api/config/v2/common"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

const schemaVersion = "1.0.3"

type SettingsObject struct {
	SchemaVersion string        `json:"schemaVersion"`
	SchemaID      string        `json:"schemaId"`
	Scope         string        `json:"scope"`
	Value         *MetricEvents `json:"value"`
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
func (cs *ServiceClient) Create(config *MetricEvents) (string, error) {
	return cs.creationAttempt(config, false)
}

func (cs *ServiceClient) creationAttempt(config *MetricEvents, isRetry bool) (string, error) {
	payload := v2common.SettingsObjectCreate{
		Value:         config,
		SchemaID:      "builtin:anomaly-detection.metric-events",
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
			if len(soer[0].Error.ConstraintViolations) > 0 {
				for _, violation := range soer[0].Error.ConstraintViolations {
					violationMessage := *violation.Message
					if strings.Contains(violationMessage, "Metric selectors could only be parsed in backward compatibility mode") {
						if strings.Contains(violationMessage, "Consider querying `") && strings.Contains(violationMessage, "` instead..") {
							metricSelector := violationMessage[strings.Index(violationMessage, "Consider querying `")+len("Consider querying `") : strings.Index(violationMessage, "` instead..")]
							if isRetry {
								log.Println("retry failed with another error: ", err.Error())
								return "", err
							} else {
								log.Println("retrying with metricSelector `" + metricSelector + "`")
								config.QueryDefinition.MetricSelector = &metricSelector
								return cs.creationAttempt(config, true)
							}
						}
					}
				}
			}
			od, _ := json.Marshal(soer[0])
			return "", errors.New(string(od))
		}
		return "", err
	} else {
		return "", err
	}
}

// Update TODO: documentation
func (cs *ServiceClient) Update(id string, item *MetricEvents) error {
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
func (cs *ServiceClient) Get(id string) (*MetricEvents, error) {
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

// List TODO: documentation
func (cs *ServiceClient) List() ([]string, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/settings/objects?schemaIds=builtin%3Aanomaly-detection.metric-events&fields=objectId&pageSize=500", 200); err != nil {
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

func (cs *ServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

func (cs *ServiceClient) LIST() ([]string, error) {
	return cs.List()
}