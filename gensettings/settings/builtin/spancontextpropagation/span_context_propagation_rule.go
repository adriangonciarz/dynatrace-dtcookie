package spancontextpropagation

import "github.com/dtcookie/hcl"

type SpanContextPropagationRule struct {
	Matchers   SpanMatchers                 `json:"matchers"`
	RuleAction SpanContextPropagationAction `json:"ruleAction"` // Rule action
	RuleName   string                       `json:"ruleName"`   // Rule name
}

func (me *SpanContextPropagationRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"matchers": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SpanMatchers).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"rule_action": {
			Type:        hcl.TypeString,
			Description: "Rule action",
			Required:    true,
		},
		"rule_name": {
			Type:        hcl.TypeString,
			Description: "Rule name",
			Required:    true,
		},
	}
}

func (me *SpanContextPropagationRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"matchers":    me.Matchers,
		"rule_action": me.RuleAction,
		"rule_name":   me.RuleName,
	})
}

func (me *SpanContextPropagationRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"matchers":    &me.Matchers,
		"rule_action": &me.RuleAction,
		"rule_name":   &me.RuleName,
	})
}
