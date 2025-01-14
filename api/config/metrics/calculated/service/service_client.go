package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
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
func (cs *ServiceClient) Create(config *CalculatedServiceMetric) (*api.EntityRef, error) {
	var err error
	var bytes []byte

	retry := true
	maxAttempts := 10
	attempts := 0

	for retry {
		attempts = attempts + 1
		if bytes, err = cs.client.POST("/calculatedMetrics/service", config, 201); err != nil {
			if !strings.Contains(err.Error(), "Metric definition must specify a known request attribute") {
				return nil, err
			}
			log.Println(".... request attribute is not fully known yet to cluster - retrying")
			if attempts < maxAttempts {
				time.Sleep(2 * time.Second)
			} else {
				log.Println(".... giving up")
			}
		} else {
			retry = false
		}
	}

	var stub api.EntityRef
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(config *CalculatedServiceMetric) error {
	if _, err := cs.client.PUT(fmt.Sprintf("/calculatedMetrics/service/%s", config.TsmMetricKey), config, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the Notification to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/calculatedMetrics/service/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*CalculatedServiceMetric, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the Notification to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/calculatedMetrics/service/%s", id), 200); err != nil {
		return nil, err
	}
	var record CalculatedServiceMetric
	if err = json.Unmarshal(bytes, &record); err != nil {
		return nil, err
	}
	return &record, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) ListAll() (*api.EntityRefs, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/calculatedMetrics/service", 200); err != nil {
		return nil, err
	}
	var stubList api.EntityRefs
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}

func (cs *ServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

func (cs *ServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if stubList, err := cs.ListAll(); err == nil {
		for _, stub := range stubList.Values {
			ids = append(ids, stub.ID)
		}
	}
	return ids, err
}

func (cs *ServiceClient) ListInterface() (interface{}, error) {
	return cs.ListAll()
}
