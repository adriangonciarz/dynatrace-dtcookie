package locations_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/dtcookie/dynatrace/api/config/synthetic/locations"
	"github.com/dtcookie/dynatrace/rest"
)

func TestServiceclient(t *testing.T) {
	rest.Verbose = true
	envURL := os.Getenv("DYNATRACE_ENV_URL")
	apiToken := os.Getenv("DYNATRACE_API_TOKEN")
	client := locations.NewService(envURL+"/api/v1", apiToken)

	results, err := client.LIST()
	if err != nil {
		t.Error(err)
	}
	for idx, result := range results {
		t.Log(idx, result)
		record, err := client.Get(result)
		if err != nil {
			t.Error(err)
		}
		data, _ := json.MarshalIndent(record, "", "  ")
		t.Log(string(data))
	}

}
