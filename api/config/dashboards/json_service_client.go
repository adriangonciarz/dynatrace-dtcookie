package dashboards

import (
	"encoding/json"
	"fmt"
	"strings"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

type JSONServiceClient struct {
	client *rest.Client
}

func NewJSONService(baseURL string, token string) *JSONServiceClient {
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &JSONServiceClient{client: client}
}

func (cs *JSONServiceClient) Validate(dashboard *JSONDashboard) []string {
	return cs.validate(dashboard, 0)
}

func (cs *JSONServiceClient) validate(dashboard *JSONDashboard, retryCnt int) []string {
	var err error

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

func (cs *JSONServiceClient) Create(dashboard *JSONDashboard) (*api.EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.POST("/dashboards", dashboard, 201); err != nil {
		return nil, err
	}
	var stub api.EntityShortRepresentation
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

func (cs *JSONServiceClient) Update(dashboard *JSONDashboard, id string) error {
	_, err := cs.client.PUT(fmt.Sprintf("/dashboards/%s", id), dashboard, 204)
	return err
}

func (cs *JSONServiceClient) Delete(id string) error {
	_, err := cs.client.DELETE(fmt.Sprintf("/dashboards/%s", id), 204)
	return err
}

func (cs *JSONServiceClient) Get(id string) (*JSONDashboard, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/dashboards/%s", id), 200); err != nil {
		return nil, err
	}
	var dashboard JSONDashboard
	if err = json.Unmarshal(bytes, &dashboard); err != nil {
		return nil, err
	}
	return &dashboard, nil
}

func (cs *JSONServiceClient) ListAll() (*DashboardList, error) {
	return cs.List("")
}

func (cs *JSONServiceClient) List(owner string, tags ...string) (*DashboardList, error) {
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

func (cs *JSONServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

func (cs *JSONServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if dashboardList, err := cs.ListAll(); err == nil {
		for _, dashboard := range dashboardList.Dashboards {
			ids = append(ids, dashboard.ID)
		}
	}
	return ids, err
}
