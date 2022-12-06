package metricevents

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

// Service TODO: documentation
type Service struct {
	client *rest.Client
}

// NewService TODO: documentation
// "https://#######.live.dynatrace.com/api/config/v1", "###########"
func NewService(baseURL string, token string) *Service {
	rest.Verbose = false
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &Service{client: client}
}

// Create TODO: documentation
func (cs *Service) Create(config *MetricEvent) (*api.EntityRef, error) {
	return cs.creationAttempt(config, false)
}

func (cs *Service) creationAttempt(config *MetricEvent, isRetry bool) (*api.EntityRef, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.POST("/anomalyDetection/metricEvents", config, 201); err != nil {
		if restErr, ok := err.(*rest.Error); ok {
			if len(restErr.ConstraintViolations) > 0 {
				for _, violation := range restErr.ConstraintViolations {
					violationMessage := violation.Message
					if strings.Contains(violationMessage, "Metric selectors could only be parsed in backward compatibility mode") {
						if strings.Contains(violationMessage, "Consider querying `") && strings.Contains(violationMessage, "` instead..") {
							metricSelector := violationMessage[strings.Index(violationMessage, "Consider querying `")+len("Consider querying `") : strings.Index(violationMessage, "` instead..")]
							if isRetry {
								log.Println("retry failed with another error: ", err.Error())
								return nil, err
							} else {
								log.Println("retrying with metricSelector `" + metricSelector + "`")
								config.MetricSelector = &metricSelector
								return cs.creationAttempt(config, true)
							}
						}
					}
				}
			}
		}
		return nil, err
	}
	var stub api.EntityRef
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *Service) Update(config *MetricEvent) error {
	if _, err := cs.client.PUT(fmt.Sprintf("/anomalyDetection/metricEvents/%s", *config.ID), config, 204); err != nil {
		return err
	}
	return nil
}

// Validate TODO: documentation
func (cs *Service) Validate(config *MetricEvent) error {
	if config.ID != nil {
		if _, err := cs.client.POST(fmt.Sprintf("/anomalyDetection/metricEvents/%s/validator", *config.ID), config, 204); err != nil {
			return err
		}
	} else {
		if _, err := cs.client.POST("/anomalyDetection/metricEvents/validator", config, 204); err != nil {
			return err
		}
	}
	return nil
}

// Delete TODO: documentation
func (cs *Service) Delete(id string) error {
	if _, err := cs.client.DELETE(fmt.Sprintf("/anomalyDetection/metricEvents/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *Service) Get(id string) (*MetricEvent, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/anomalyDetection/metricEvents/%s", url.QueryEscape(id)), 200); err != nil {
		return nil, err
	}
	var config MetricEvent
	if err = json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// List TODO: documentation
func (cs *Service) List() (*api.EntityRefs, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/anomalyDetection/metricEvents", 200); err != nil {
		return nil, err
	}
	var stubList api.EntityRefs
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}

func (cs *Service) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

func (cs *Service) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if stubList, err := cs.List(); err == nil {
		for _, stub := range stubList.Values {
			ids = append(ids, stub.ID)
		}
	}
	return ids, err
}
