package logdpprules

import "github.com/dtcookie/hcl"

type ProcessorDefinition struct {
	Rule string `json:"rule"` // Processor definition
}

func (me *ProcessorDefinition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rule": {
			Type:        hcl.TypeString,
			Description: "Processor definition",
			Required:    true,
		},
	}
}

func (me *ProcessorDefinition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"rule": me.Rule,
	})
}

func (me *ProcessorDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"rule": &me.Rule,
	})
}
