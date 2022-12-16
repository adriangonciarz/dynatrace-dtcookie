package syntheticavailabilitysettings

import "github.com/dtcookie/hcl"

type Settings struct {
	ExcludeMaintenanceWindows bool `json:"excludeMaintenanceWindows"` // Exclude periods with maintenance windows from availability calculation
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"exclude_maintenance_windows": {
			Type:        hcl.TypeBool,
			Description: "Exclude periods with maintenance windows from availability calculation",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"exclude_maintenance_windows": me.ExcludeMaintenanceWindows,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"exclude_maintenance_windows": &me.ExcludeMaintenanceWindows,
	})
}
