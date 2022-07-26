package application

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

// ApplicationClient holds REST API connection details
type ApplicationClient struct {
	client *rest.Client
}

// NewService creates a new Application Client
// baseURL should look like this: "https://siz65484.live.dynatrace.com/api/v1"
// token is an API Token
func NewService(baseURL string, token string) *ApplicationClient {
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ApplicationClient{client: client}
}

// List function retrieves the list of all applications in the environment
func (cs *ApplicationClient) List() (Applications, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/entity/applications", 200); err != nil {
		return nil, err
	}
	var applicationList Applications
	if err = json.Unmarshal(bytes, &applicationList); err != nil {
		return nil, err
	}

	return applicationList, nil
}
