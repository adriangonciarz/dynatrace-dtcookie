package maintenancewindow

import "github.com/dtcookie/hcl"

type GeneralProperties struct {
	Description                      string          `json:"description"`                      // A short description of the maintenance purpose.
	DisableSyntheticMonitorExecution bool            `json:"disableSyntheticMonitorExecution"` // Disables the execution of the synthetic monitors that are within [the scope of this maintenance window](https://dt-url.net/0e0341m).
	MaintenanceType                  MaintenanceType `json:"maintenanceType"`                  // Whether the maintenance is planned or unplanned.
	Name                             string          `json:"name"`
	Suppression                      SuppressionType `json:"suppression"` // Defines if alerting or problem generation is disabled.. * **Detect problems and alert**: Problems are generated and alerted.\n* **Detect problems but don't alert**: Problems are generated but no alerts are sent out.\n* **Disable problem detection during maintenance**: Neither problems are generated nor alerts are sent out.
}

func (me *GeneralProperties) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"description": {
			Type:        hcl.TypeString,
			Description: "A short description of the maintenance purpose.",
			Optional:    true,
		},
		"disable_synthetic_monitor_execution": {
			Type:        hcl.TypeBool,
			Description: "Disables the execution of the synthetic monitors that are within [the scope of this maintenance window](https://dt-url.net/0e0341m).",
			Optional:    true,
		},
		"maintenance_type": {
			Type:        hcl.TypeString,
			Description: "Whether the maintenance is planned or unplanned.",
			Required:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"suppression": {
			Type:        hcl.TypeString,
			Description: "Defines if alerting or problem generation is disabled.. * **Detect problems and alert**: Problems are generated and alerted.\n* **Detect problems but don't alert**: Problems are generated but no alerts are sent out.\n* **Disable problem detection during maintenance**: Neither problems are generated nor alerts are sent out.",
			Required:    true,
		},
	}
}

func (me *GeneralProperties) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"description":                         me.Description,
		"disable_synthetic_monitor_execution": me.DisableSyntheticMonitorExecution,
		"maintenance_type":                    me.MaintenanceType,
		"name":                                me.Name,
		"suppression":                         me.Suppression,
	})
}

func (me *GeneralProperties) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"description":                         &me.Description,
		"disable_synthetic_monitor_execution": &me.DisableSyntheticMonitorExecution,
		"maintenance_type":                    &me.MaintenanceType,
		"name":                                &me.Name,
		"suppression":                         &me.Suppression,
	})
}
