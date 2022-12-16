package rule

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled              bool                 `json:"enabled"`              // Enabled
	Matcher              string               `json:"matcher"`              // [See our documentation](https://dt-url.net/bp234rv)
	RuleName             string               `json:"ruleName"`             // Rule name
	Script               string               `json:"script"`               // [See our documentation](https://dt-url.net/8k03xm2)
	TransformationFields TransformationFields `json:"transformationFields"` // Transformation fields
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"matcher": {
			Type:        hcl.TypeString,
			Description: "[See our documentation](https://dt-url.net/bp234rv)",
			Required:    true,
		},
		"rule_name": {
			Type:        hcl.TypeString,
			Description: "Rule name",
			Required:    true,
		},
		"script": {
			Type:        hcl.TypeString,
			Description: "[See our documentation](https://dt-url.net/8k03xm2)",
			Required:    true,
		},
		"transformation_fields": {
			Type:        hcl.TypeList,
			Description: "Transformation fields",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(TransformationFields).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":               me.Enabled,
		"matcher":               me.Matcher,
		"rule_name":             me.RuleName,
		"script":                me.Script,
		"transformation_fields": me.TransformationFields,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":               &me.Enabled,
		"matcher":               &me.Matcher,
		"rule_name":             &me.RuleName,
		"script":                &me.Script,
		"transformation_fields": &me.TransformationFields,
	})
}
