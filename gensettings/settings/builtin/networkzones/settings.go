package networkzones

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled bool `json:"enabled"` // For details, see [Network zones](https://www.dynatrace.com/support/help/shortlink/network-zones). \n\n⚠ Warning: You may experience network imbalance if you suddenly introduce network zones into an environment that has a high number of OneAgents. To avoid this and to ensure smooth adoption of network zones, follow best practices for planning and proper naming, as explained in [Network zones](https://www.dynatrace.com/support/help/shortlink/network-zones).
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "For details, see [Network zones](https://www.dynatrace.com/support/help/shortlink/network-zones). \n\n⚠ Warning: You may experience network imbalance if you suddenly introduce network zones into an environment that has a high number of OneAgents. To avoid this and to ensure smooth adoption of network zones, follow best practices for planning and proper naming, as explained in [Network zones](https://www.dynatrace.com/support/help/shortlink/network-zones).",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled": me.Enabled,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled": &me.Enabled,
	})
}
