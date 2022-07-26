package host

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

// HostClient holds REST API connection details
type HostClient struct {
	client *rest.Client
}

// NewService creates a new Host Client
// baseURL should look like this: "https://siz65484.live.dynatrace.com/api/v1"
// token is an API Token
func NewService(baseURL string, token string) *HostClient {
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &HostClient{client: client}
}

// List function retrieves the list of all hosts in the environment
func (cs *HostClient) List() (Hosts, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/entity/infrastructure/hosts", 200); err != nil {
		return nil, err
	}
	var hostList Hosts
	if err = json.Unmarshal(bytes, &hostList); err != nil {
		return nil, err
	}

	return hostList, nil
}
