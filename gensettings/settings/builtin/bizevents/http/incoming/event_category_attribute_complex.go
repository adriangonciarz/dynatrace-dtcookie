package incoming

import "github.com/dtcookie/hcl"

type EventCategoryAttributeComplex struct {
	Path       string                         `json:"path"`       // [See our documentation](https://dt-url.net/ei034bx)
	Source     string                         `json:"source"`     // Fixed value
	SourceType DataSourceWithStaticStringEnum `json:"sourceType"` // Data source
}

func (me *EventCategoryAttributeComplex) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"path": {
			Type:        hcl.TypeString,
			Description: "[See our documentation](https://dt-url.net/ei034bx)",
			Required:    true,
		},
		"source": {
			Type:        hcl.TypeString,
			Description: "Fixed value",
			Required:    true,
		},
		"source_type": {
			Type:        hcl.TypeString,
			Description: "Data source",
			Required:    true,
		},
	}
}

func (me *EventCategoryAttributeComplex) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"path":        me.Path,
		"source":      me.Source,
		"source_type": me.SourceType,
	})
}

func (me *EventCategoryAttributeComplex) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"path":        &me.Path,
		"source":      &me.Source,
		"source_type": &me.SourceType,
	})
}
