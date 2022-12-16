package logdpprules

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled             bool                 `json:"enabled"`             // Active
	ProcessorDefinition *ProcessorDefinition `json:"ProcessorDefinition"` // ## Processor rule definition\nAdd a rule definition using our syntax. [In our documentation](https://dt-url.net/8k03xm2) you will find instructions and application [examples](https://dt-url.net/m24305t).
	Query               string               `json:"query"`               // Matcher
	RuleName            string               `json:"ruleName"`            // Processor name
	RuleTesting         *RuleTesting         `json:"RuleTesting"`         // ## Rule testing\n### 1. Select a log sample
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Active",
			Optional:    true,
		},
		"processor_definition": {
			Type:        hcl.TypeList,
			Description: "## Processor rule definition\nAdd a rule definition using our syntax. [In our documentation](https://dt-url.net/8k03xm2) you will find instructions and application [examples](https://dt-url.net/m24305t).",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ProcessorDefinition).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"query": {
			Type:        hcl.TypeString,
			Description: "Matcher",
			Required:    true,
		},
		"rule_name": {
			Type:        hcl.TypeString,
			Description: "Processor name",
			Required:    true,
		},
		"rule_testing": {
			Type:        hcl.TypeList,
			Description: "## Rule testing\n### 1. Select a log sample",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(RuleTesting).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":              me.Enabled,
		"processor_definition": me.ProcessorDefinition,
		"query":                me.Query,
		"rule_name":            me.RuleName,
		"rule_testing":         me.RuleTesting,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":              &me.Enabled,
		"processor_definition": &me.ProcessorDefinition,
		"query":                &me.Query,
		"rule_name":            &me.RuleName,
		"rule_testing":         &me.RuleTesting,
	})
}
