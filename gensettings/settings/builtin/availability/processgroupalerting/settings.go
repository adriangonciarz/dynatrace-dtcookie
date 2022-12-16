package processgroupalerting

import "github.com/dtcookie/hcl"

type Settings struct {
	AlertingMode             AlertingMode `json:"alertingMode"`             // **if any process becomes unavailable:**\nDynatrace will open a new problem if a single process in this group shuts down or crashes. \n\n**if minimum threshold is not met:**\nDynatrace tracks the number of process instances that comprise this process group and treats the group as a cluster. This setting enables you to define a minimum number of process instances that must be available. A problem will be opened if this process group has fewer than the minimum number of required process instances. \n\n Details of the related impact on service requests will be included in the problem summary.\n\n**Note:** If a process is intentionally shutdown or retired while this setting is active, you'll need to manually close the problem.
	Enabled                  bool         `json:"enabled"`                  // Enable process group availability monitoring
	MinimumInstanceThreshold int          `json:"minimumInstanceThreshold"` // Open a new problem if the number of active process instances in the group is fewer than:
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"alerting_mode": {
			Type:        hcl.TypeString,
			Description: "**if any process becomes unavailable:**\nDynatrace will open a new problem if a single process in this group shuts down or crashes. \n\n**if minimum threshold is not met:**\nDynatrace tracks the number of process instances that comprise this process group and treats the group as a cluster. This setting enables you to define a minimum number of process instances that must be available. A problem will be opened if this process group has fewer than the minimum number of required process instances. \n\n Details of the related impact on service requests will be included in the problem summary.\n\n**Note:** If a process is intentionally shutdown or retired while this setting is active, you'll need to manually close the problem.",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enable process group availability monitoring",
			Optional:    true,
		},
		"minimum_instance_threshold": {
			Type:        hcl.TypeInt,
			Description: "Open a new problem if the number of active process instances in the group is fewer than:",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"alerting_mode":              me.AlertingMode,
		"enabled":                    me.Enabled,
		"minimum_instance_threshold": me.MinimumInstanceThreshold,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"alerting_mode":              &me.AlertingMode,
		"enabled":                    &me.Enabled,
		"minimum_instance_threshold": &me.MinimumInstanceThreshold,
	})
}
