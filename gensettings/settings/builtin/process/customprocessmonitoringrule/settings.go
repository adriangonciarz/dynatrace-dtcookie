package customprocessmonitoringrule

import "github.com/dtcookie/hcl"

type Settings struct {
	Condition *Condition     `json:"condition"` // Condition
	Enabled   bool           `json:"enabled"`   // Enabled
	Mode      MonitoringMode `json:"mode"`      // Mode
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeList,
			Description: "Condition",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Condition).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"mode": {
			Type:        hcl.TypeString,
			Description: "Mode",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"condition": me.Condition,
		"enabled":   me.Enabled,
		"mode":      me.Mode,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"condition": &me.Condition,
		"enabled":   &me.Enabled,
		"mode":      &me.Mode,
	})
}
