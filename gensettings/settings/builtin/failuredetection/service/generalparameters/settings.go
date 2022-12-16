package generalparameters

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled        bool            `json:"enabled"`        // Override global failure detection settings
	ExceptionRules *ExceptionRules `json:"exceptionRules"` // Customize failure detection for specific exceptions and errors
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Override global failure detection settings",
			Optional:    true,
		},
		"exception_rules": {
			Type:        hcl.TypeList,
			Description: "Customize failure detection for specific exceptions and errors",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ExceptionRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":         me.Enabled,
		"exception_rules": me.ExceptionRules,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":         &me.Enabled,
		"exception_rules": &me.ExceptionRules,
	})
}
