package osservicesmonitoring

import "github.com/dtcookie/hcl"

type Settings struct {
	AlertActivationDuration    int                        `json:"alertActivationDuration"`    // The number of **10-second measurement cycles** before alerting is triggered
	Alerting                   bool                       `json:"alerting"`                   // Toggle the switch in order to enable or disable alerting for this policy
	DetectionConditionsLinux   LinuxDetectionConditions   `json:"detectionConditionsLinux"`   // Detection rules
	DetectionConditionsWindows WindowsDetectionConditions `json:"detectionConditionsWindows"` // Detection rules
	Enabled                    bool                       `json:"enabled"`                    // Enabled
	Metadata                   MetadataItems              `json:"metadata"`                   // Set of additional key-value properties to be attached to the triggered event.
	Monitoring                 bool                       `json:"monitoring"`                 // Toggle the switch in order to enable or disable availability metric monitoring for this policy. Availability metrics consume custom metrics (DDUs). Refer to [documentation](https://dt-url.net/vl03xzk) for DDU consumption examples. Each monitored service consumes one custom metric.
	Name                       string                     `json:"name"`                       // Rule name
	NotInstalledAlerting       bool                       `json:"notInstalledAlerting"`       // By default, Dynatrace does not alert if the service is not installed. Toggle the switch to enable or disable this feature
	StatusConditionLinux       string                     `json:"statusConditionLinux"`       // This string has to match a required format. See [OS services monitoring](https://dt-url.net/vl03xzk).\n\n- `$eq(failed)` – Matches services that are in failed state.\n\nAvailable logic operations:\n- `$not($eq(active))` – Matches services with state different from active.\n- `$or($eq(inactive),$eq(failed))` – Matches services that are either in inactive or failed state.\n\nUse one of the following values as a parameter for this condition:\n\n- `reloading`\n- `activating`\n- `deactivating`\n- `failed`\n- `inactive`\n- `active`
	StatusConditionWindows     string                     `json:"statusConditionWindows"`     // This string has to match a required format. See [OS services monitoring](https://dt-url.net/vl03xzk).\n\n- `$eq(paused)` – Matches services that are in paused state.\n\nAvailable logic operations:\n- `$not($eq(paused))` – Matches services that are in state different from paused.\n- `$or($eq(paused),$eq(running))` – Matches services that are either in paused or running state.\n\nUse one of the following values as a parameter for this condition:\n\n- `running`\n- `stopped`\n- `start_pending`\n- `stop_pending`\n- `continue_pending`\n- `pause_pending`\n- `paused`
	System                     System                     `json:"system"`                     // System
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"alert_activation_duration": {
			Type:        hcl.TypeInt,
			Description: "The number of **10-second measurement cycles** before alerting is triggered",
			Required:    true,
		},
		"alerting": {
			Type:        hcl.TypeBool,
			Description: "Toggle the switch in order to enable or disable alerting for this policy",
			Optional:    true,
		},
		"detection_conditions_linux": {
			Type:        hcl.TypeList,
			Description: "Detection rules",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(LinuxDetectionConditions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_conditions_windows": {
			Type:        hcl.TypeList,
			Description: "Detection rules",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(WindowsDetectionConditions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"metadata": {
			Type:        hcl.TypeList,
			Description: "Set of additional key-value properties to be attached to the triggered event.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MetadataItems).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"monitoring": {
			Type:        hcl.TypeBool,
			Description: "Toggle the switch in order to enable or disable availability metric monitoring for this policy. Availability metrics consume custom metrics (DDUs). Refer to [documentation](https://dt-url.net/vl03xzk) for DDU consumption examples. Each monitored service consumes one custom metric.",
			Optional:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "Rule name",
			Required:    true,
		},
		"not_installed_alerting": {
			Type:        hcl.TypeBool,
			Description: "By default, Dynatrace does not alert if the service is not installed. Toggle the switch to enable or disable this feature",
			Optional:    true,
		},
		"status_condition_linux": {
			Type:        hcl.TypeString,
			Description: "This string has to match a required format. See [OS services monitoring](https://dt-url.net/vl03xzk).\n\n- `$eq(failed)` – Matches services that are in failed state.\n\nAvailable logic operations:\n- `$not($eq(active))` – Matches services with state different from active.\n- `$or($eq(inactive),$eq(failed))` – Matches services that are either in inactive or failed state.\n\nUse one of the following values as a parameter for this condition:\n\n- `reloading`\n- `activating`\n- `deactivating`\n- `failed`\n- `inactive`\n- `active`",
			Required:    true,
		},
		"status_condition_windows": {
			Type:        hcl.TypeString,
			Description: "This string has to match a required format. See [OS services monitoring](https://dt-url.net/vl03xzk).\n\n- `$eq(paused)` – Matches services that are in paused state.\n\nAvailable logic operations:\n- `$not($eq(paused))` – Matches services that are in state different from paused.\n- `$or($eq(paused),$eq(running))` – Matches services that are either in paused or running state.\n\nUse one of the following values as a parameter for this condition:\n\n- `running`\n- `stopped`\n- `start_pending`\n- `stop_pending`\n- `continue_pending`\n- `pause_pending`\n- `paused`",
			Required:    true,
		},
		"system": {
			Type:        hcl.TypeString,
			Description: "System",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"alert_activation_duration":    me.AlertActivationDuration,
		"alerting":                     me.Alerting,
		"detection_conditions_linux":   me.DetectionConditionsLinux,
		"detection_conditions_windows": me.DetectionConditionsWindows,
		"enabled":                      me.Enabled,
		"metadata":                     me.Metadata,
		"monitoring":                   me.Monitoring,
		"name":                         me.Name,
		"not_installed_alerting":       me.NotInstalledAlerting,
		"status_condition_linux":       me.StatusConditionLinux,
		"status_condition_windows":     me.StatusConditionWindows,
		"system":                       me.System,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"alert_activation_duration":    &me.AlertActivationDuration,
		"alerting":                     &me.Alerting,
		"detection_conditions_linux":   &me.DetectionConditionsLinux,
		"detection_conditions_windows": &me.DetectionConditionsWindows,
		"enabled":                      &me.Enabled,
		"metadata":                     &me.Metadata,
		"monitoring":                   &me.Monitoring,
		"name":                         &me.Name,
		"not_installed_alerting":       &me.NotInstalledAlerting,
		"status_condition_linux":       &me.StatusConditionLinux,
		"status_condition_windows":     &me.StatusConditionWindows,
		"system":                       &me.System,
	})
}
