package presets

import (
	"github.com/dtcookie/hcl"
)

type DashboardPresetsList []*DashboardPresets

func (me *DashboardPresetsList) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dashboardPresets": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(DashboardPresets).Schema()},
		},
	}
}

func (me DashboardPresetsList) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("dashboardPresets", me)
}

func (me *DashboardPresetsList) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("dashboardPresets", me)
}

type DashboardPresets struct {
	DashboardPreset string `json:"DashboardPreset"` // Dashboard preset to limit visibility for
	UserGroup       string `json:"UserGroup"`       // User group to show selected dashboard preset to
}

func (me *DashboardPresets) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dashboard_preset": {
			Type:        hcl.TypeString,
			Description: "Dashboard preset to limit visibility for",
			Required:    true,
		},
		"user_group": {
			Type:        hcl.TypeString,
			Description: "User group to show selected dashboard preset to",
			Required:    true,
		},
	}
}

func (me *DashboardPresets) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"dashboard_preset": me.DashboardPreset,
		"user_group":       me.UserGroup,
	})
}

func (me *DashboardPresets) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"dashboard_preset": &me.DashboardPreset,
		"user_group":       &me.UserGroup,
	})
}
