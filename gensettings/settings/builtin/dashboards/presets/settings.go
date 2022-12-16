package presets

import "github.com/dtcookie/hcl"

type Settings struct {
	DashboardPresetsList   DashboardPresetsList `json:"dashboardPresetsList"`   // Show selected preset to respective user group only.
	EnableDashboardPresets bool                 `json:"enableDashboardPresets"` // Dashboard presets are visible to all users by default. For a pristine environment you may disable them entirely or opt to manually limit visibility to selected user groups.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dashboard_presets_list": {
			Type:        hcl.TypeList,
			Description: "Show selected preset to respective user group only.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DashboardPresetsList).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enable_dashboard_presets": {
			Type:        hcl.TypeBool,
			Description: "Dashboard presets are visible to all users by default. For a pristine environment you may disable them entirely or opt to manually limit visibility to selected user groups.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"dashboard_presets_list":   me.DashboardPresetsList,
		"enable_dashboard_presets": me.EnableDashboardPresets,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"dashboard_presets_list":   &me.DashboardPresetsList,
		"enable_dashboard_presets": &me.EnableDashboardPresets,
	})
}
