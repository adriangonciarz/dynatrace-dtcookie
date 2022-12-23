package consumption

import (
	"github.com/dtcookie/hcl"
)

type DDUPoolConfig struct {
	LimitEnabled bool   `json:"limitEnabled"`
	LimitType    string `json:"limitType"`
	LimitValue   int    `json:"limitValue"`
}

func (me *DDUPoolConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"limit_enabled": {
			Type:        hcl.TypeBool,
			Required:    true,
			Description: "Is the limit configuration enabled",
		},
		"limit_type": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "Type of the limit applied: MONTHLY or ANNUAL",
		},
		"limit_value": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Value of the DDU limit applied for provided timerange",
		},
	}
}

func (me *DDUPoolConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"limit_enabled": me.LimitEnabled,
		"limit_type":    me.LimitType,
		"limit_value":   me.LimitValue,
	})
}

func (me *DDUPoolConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"limit_enabled": &me.LimitEnabled,
		"limit_type":    &me.LimitType,
		"limit_value":   &me.LimitValue,
	})
}
