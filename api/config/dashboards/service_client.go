package dashboards

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	api "github.com/dtcookie/dynatrace/api/config"
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

func (cs *ServiceClient) Validate(dashboard *Dashboard) []string {
	return cs.validate(dashboard, 0)
}

func (cs *ServiceClient) validate(dashboard *Dashboard, retryCnt int) []string {
	var err error

	if len(opt.String(dashboard.ID)) > 0 {
		return []string{"you MUST NOT provide an ID within the Dashboard payload upon creation"}
	}
	if _, err = cs.client.POST("/dashboards/validator", dashboard, 204); err != nil {
		if env, ok := err.(*rest.Error); ok {
			if len(env.ConstraintViolations) > 0 {
				errs := []string{}
				for _, violation := range env.ConstraintViolations {
					errs = append(errs, violation.Message)
				}
				return errs
			} else if env.Message == "Token is missing required scope. Use one of: WriteConfig (Write configuration)" {
				return []string{env.Message}
			} else {
				return []string{err.Error()}
			}
		} else {
			if strings.HasSuffix(err.Error(), "/validator\": EOF") {
				if retryCnt < 5 {
					return cs.validate(dashboard, retryCnt+1)
				}
			}
			return []string{err.Error()}
		}

	}
	return nil
}

// Create TODO: documentation
func (cs *ServiceClient) Create(dashboard *Dashboard) (*api.EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if len(opt.String(dashboard.ID)) > 0 {
		return nil, errors.New("you MUST NOT provide an ID within the Dashboard payload upon creation")
	}

	if bytes, err = cs.client.POST("/dashboards", dashboard, 201); err != nil {
		return nil, err
	}
	var stub api.EntityShortRepresentation
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(dashboard *Dashboard) error {
	if len(opt.String(dashboard.ID)) == 0 {
		return errors.New("the Dashboard doesn't contain an ID")
	}
	if _, err := cs.client.PUT(fmt.Sprintf("/dashboards/%s", opt.String(dashboard.ID)), dashboard, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the Dashboard to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/dashboards/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*Dashboard, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the Dashboard to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/dashboards/%s", id), 200); err != nil {
		return nil, err
	}
	var dashboard Dashboard
	if err = json.Unmarshal(bytes, &dashboard); err != nil {
		return nil, err
	}
	return &dashboard, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) ListAll() (*DashboardList, error) {
	return cs.List("")
}

// List TODO: documentation
func (cs *ServiceClient) List(owner string, tags ...string) (*DashboardList, error) {
	var err error
	var bytes []byte

	var qb querybuilder
	if len(owner) > 0 {
		qb.Append("owner", owner)
	}
	if (tags != nil) && (len(tags) > 0) {
		for _, tag := range tags {
			qb.Append("tags", tag)
		}
	}
	if bytes, err = cs.client.GET(qb.Build("/dashboards"), 200); err != nil {
		return nil, err
	}
	var dashboardList DashboardList
	if err = json.Unmarshal(bytes, &dashboardList); err != nil {
		return nil, err
	}
	return &dashboardList, nil
}

func (cs *ServiceClient) GET(id string) (interface{}, error) {
	config, err := cs.Get(id)
	config.ConfigurationMetadata = nil
	return config, err
}

func (cs *ServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if dashboardList, err := cs.ListAll(); err == nil {
		for _, dashboard := range dashboardList.Dashboards {
			ids = append(ids, dashboard.ID)
		}
	}
	return ids, err
}
