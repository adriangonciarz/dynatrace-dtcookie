package spanentrypoints

import "github.com/dtcookie/hcl"

type Settings struct {
	EntryPointRule *SpanEntrypointRule `json:"entryPointRule"` // Entry Point Rule
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"entry_point_rule": {
			Type:        hcl.TypeList,
			Description: "Entry Point Rule",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SpanEntrypointRule).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"entry_point_rule": me.EntryPointRule,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"entry_point_rule": &me.EntryPointRule,
	})
}
