package applicationdetectionrules

import (
	"encoding/json"
	"fmt"

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
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &Service{client: client}
}

// Create TODO: documentation
func (cs *Service) Create(rule *ApplicationDetectionRule) (*api.EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.POST("/applicationDetectionRules", rule, 201); err != nil {
		return nil, err
	}
	var stub api.EntityShortRepresentation
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *Service) Update(rule *ApplicationDetectionRule) error {
	if _, err := cs.client.PUT(fmt.Sprintf("/applicationDetectionRules/%s", *rule.ID), rule, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *Service) Delete(id string) error {
	if _, err := cs.client.DELETE(fmt.Sprintf("/applicationDetectionRules/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *Service) Get(id string) (*ApplicationDetectionRule, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/applicationDetectionRules/%s", id), 200); err != nil {
		return nil, err
	}
	var rule ApplicationDetectionRule
	if err = json.Unmarshal(bytes, &rule); err != nil {
		return nil, err
	}
	return &rule, nil
}

// List TODO: documentation
func (cs *Service) List() (*api.StubList, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/applicationDetectionRules", 200); err != nil {
		return nil, err
	}
	var stubList api.StubList
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
