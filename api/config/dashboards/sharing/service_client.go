package sharing

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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
func (cs *ServiceClient) Create(ctx context.Context, settings *DashboardSharing) (string, error) {
	if err := cs.Update(ctx, settings); err != nil {
		return "", err
	}
	return settings.DashboardID + "-sharing", nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(ctx context.Context, settings *DashboardSharing) error {
	// debugFlag := os.Getenv("TF_LOG_PROVIDER_DYNATRACE")
	// if strings.ToUpper(debugFlag) == "DEBUG" {
	// data, _ := json.Marshal(settings)
	// tflog.Debug(ctx, fmt.Sprintf("\nPUT /dashboards/%s/shareSettings\n    %s\n", settings.DashboardID, string(data)))
	// }
	_, err := cs.client.PUT(fmt.Sprintf("/dashboards/%s/shareSettings", settings.DashboardID), settings, 201)
	if err != nil && strings.HasPrefix(err.Error(), "No Content (PUT)") {
		return nil
	}
	return err
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(ctx context.Context, id string) error {
	id = strings.TrimSuffix(id, "-sharing")
	settings := DashboardSharing{
		DashboardID: id,
		Enabled:     false,
		Preset:      false,
		Permissions: []*SharePermission{
			{
				Type:       PermissionTypes.All,
				Permission: Permissions.View,
			},
		},
		PublicAccess: &AnonymousAccess{
			ManagementZoneIDs: []string{},
			URLs:              map[string]string{},
		},
	}
	return cs.Update(ctx, &settings)
}

// Get TODO: documentation
func (cs *ServiceClient) Get(ctx context.Context, id string) (*DashboardSharing, error) {
	id = strings.TrimSuffix(id, "-sharing")
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/dashboards/%s/shareSettings", id), 200); err != nil {
		return nil, err
	}
	var settings DashboardSharing
	if err = json.Unmarshal(bytes, &settings); err != nil {
		return nil, err
	}
	return &settings, nil
}

func (cs *ServiceClient) GET(ctx context.Context, id string) (interface{}, error) {
	return cs.Get(ctx, id)
}
