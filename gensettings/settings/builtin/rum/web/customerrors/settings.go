package customerrors

import "github.com/dtcookie/hcl"

type Settings struct {
	ErrorRules                           CustomErrorRules `json:"errorRules"`
	IgnoreCustomErrorsInApdexCalculation bool             `json:"ignoreCustomErrorsInApdexCalculation"` // This setting overrides Apdex settings for individual rules listed below
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"error_rules": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CustomErrorRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"ignore_custom_errors_in_apdex_calculation": {
			Type:        hcl.TypeBool,
			Description: "This setting overrides Apdex settings for individual rules listed below",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"error_rules": me.ErrorRules,
		"ignore_custom_errors_in_apdex_calculation": me.IgnoreCustomErrorsInApdexCalculation,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"error_rules": &me.ErrorRules,
		"ignore_custom_errors_in_apdex_calculation": &me.IgnoreCustomErrorsInApdexCalculation,
	})
}
