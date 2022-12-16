package connectivityalerts

import "github.com/dtcookie/hcl"

type Settings struct {
	ConnectivityAlerts bool `json:"connectivityAlerts"` // TCP connectivity problems
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"connectivity_alerts": {
			Type:        hcl.TypeBool,
			Description: "TCP connectivity problems",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"connectivity_alerts": me.ConnectivityAlerts,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"connectivity_alerts": &me.ConnectivityAlerts,
	})
}
