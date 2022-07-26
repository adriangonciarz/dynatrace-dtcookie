package processgroup

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

// ProcessGroupClient holds REST API connection details
type ProcessGroupClient struct {
	client *rest.Client
}

// NewService creates a new Process Group Client
// baseURL should look like this: "https://siz65484.live.dynatrace.com/api/v1"
// token is an API Token
func NewService(baseURL string, token string) *ProcessGroupClient {
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ProcessGroupClient{client: client}
}

// List function retrieves the list of all process groups in the environment
func (cs *ProcessGroupClient) List() (ProcessGroups, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/entity/infrastructure/process-groups", 200); err != nil {
		return nil, err
	}
	var processGroupList ProcessGroups
	if err = json.Unmarshal(bytes, &processGroupList); err != nil {
		return nil, err
	}

	return processGroupList, nil
}
