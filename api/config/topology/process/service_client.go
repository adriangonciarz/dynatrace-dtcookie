package process

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

// ProcessClient holds REST API connection details
type ProcessClient struct {
	client *rest.Client
}

// NewService creates a new Process Client
// baseURL should look like this: "https://siz65484.live.dynatrace.com/api/v1"
// token is an API Token
func NewService(baseURL string, token string) *ProcessClient {
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ProcessClient{client: client}
}

// List function retrieves the list of all processes in the environment
func (cs *ProcessClient) List() (Processes, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/entity/infrastructure/processes", 200); err != nil {
		return nil, err
	}
	var processList Processes
	if err = json.Unmarshal(bytes, &processList); err != nil {
		return nil, err
	}

	return processList, nil
}
