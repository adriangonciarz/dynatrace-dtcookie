package limit

import "github.com/dtcookie/hcl"

type Limit struct {
	LimitEnabled bool      `json:"limitEnabled"` // Enable limit
	LimitType    LimitType `json:"limitType"`    // Time frame of limit
	LimitValue   int       `json:"limitValue"`   // The upper limit for the defined time frame
}

func (me *Limit) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"limit_enabled": {
			Type:        hcl.TypeBool,
			Description: "Enable limit",
			Optional:    true,
		},
		"limit_type": {
			Type:        hcl.TypeString,
			Description: "Time frame of limit",
			Required:    true,
		},
		"limit_value": {
			Type:        hcl.TypeInt,
			Description: "The upper limit for the defined time frame",
			Required:    true,
		},
	}
}

func (me *Limit) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"limit_enabled": me.LimitEnabled,
		"limit_type":    me.LimitType,
		"limit_value":   me.LimitValue,
	})
}

func (me *Limit) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"limit_enabled": &me.LimitEnabled,
		"limit_type":    &me.LimitType,
		"limit_value":   &me.LimitValue,
	})
}
