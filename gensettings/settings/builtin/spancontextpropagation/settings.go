package spancontextpropagation

import "github.com/dtcookie/hcl"

type Settings struct {
	ContextPropagationRule *SpanContextPropagationRule `json:"contextPropagationRule"` // Context Propagation Rule
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"context_propagation_rule": {
			Type:        hcl.TypeList,
			Description: "Context Propagation Rule",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SpanContextPropagationRule).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"context_propagation_rule": me.ContextPropagationRule,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"context_propagation_rule": &me.ContextPropagationRule,
	})
}
