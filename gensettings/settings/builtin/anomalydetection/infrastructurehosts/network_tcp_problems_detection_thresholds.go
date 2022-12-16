package infrastructurehosts

import "github.com/dtcookie/hcl"

type NetworkTcpProblemsDetectionThresholds struct {
	EventThresholds                  *EventThresholds `json:"eventThresholds"`
	FailedConnectionsNumberPerMinute int              `json:"failedConnectionsNumberPerMinute"` // Number of failed connections threshold
	NewConnectionFailuresPercentage  int              `json:"newConnectionFailuresPercentage"`  // New connection failure threshold
}

func (me *NetworkTcpProblemsDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"event_thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EventThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"failed_connections_number_per_minute": {
			Type:        hcl.TypeInt,
			Description: "Number of failed connections threshold",
			Required:    true,
		},
		"new_connection_failures_percentage": {
			Type:        hcl.TypeInt,
			Description: "New connection failure threshold",
			Required:    true,
		},
	}
}

func (me *NetworkTcpProblemsDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"event_thresholds":                     me.EventThresholds,
		"failed_connections_number_per_minute": me.FailedConnectionsNumberPerMinute,
		"new_connection_failures_percentage":   me.NewConnectionFailuresPercentage,
	})
}

func (me *NetworkTcpProblemsDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"event_thresholds":                     &me.EventThresholds,
		"failed_connections_number_per_minute": &me.FailedConnectionsNumberPerMinute,
		"new_connection_failures_percentage":   &me.NewConnectionFailuresPercentage,
	})
}
