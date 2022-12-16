package enablement

import "github.com/dtcookie/hcl"

type Rum struct {
	CostAndTrafficControl int  `json:"costAndTrafficControl"` // Percentage of user sessions captured and analyzed
	Enabled               bool `json:"enabled"`               // Enable Real User Monitoring
}

func (me *Rum) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cost_and_traffic_control": {
			Type:        hcl.TypeInt,
			Description: "Percentage of user sessions captured and analyzed",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enable Real User Monitoring",
			Optional:    true,
		},
	}
}

func (me *Rum) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"cost_and_traffic_control": me.CostAndTrafficControl,
		"enabled":                  me.Enabled,
	})
}

func (me *Rum) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cost_and_traffic_control": &me.CostAndTrafficControl,
		"enabled":                  &me.Enabled,
	})
}
