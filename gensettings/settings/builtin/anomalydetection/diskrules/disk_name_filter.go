package diskrules

import "github.com/dtcookie/hcl"

type DiskNameFilter struct {
	Operator DiskNameFilterOperator `json:"operator"`
	Value    string                 `json:"value"` // Matching text
}

func (me *DiskNameFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"operator": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "Matching text",
			Optional:    true,
		},
	}
}

func (me *DiskNameFilter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"operator": me.Operator,
		"value":    me.Value,
	})
}

func (me *DiskNameFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"operator": &me.Operator,
		"value":    &me.Value,
	})
}
