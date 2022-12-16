package profile

import "github.com/dtcookie/hcl"

type TextFilter struct {
	CaseSensitive bool               `json:"caseSensitive"` // Case sensitive
	Enabled       bool               `json:"enabled"`
	Negate        bool               `json:"negate"`
	Operator      ComparisonOperator `json:"operator"` // Operator of the comparison
	Value         string             `json:"value"`
}

func (me *TextFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"case_sensitive": {
			Type:        hcl.TypeBool,
			Description: "Case sensitive",
			Optional:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "no documentation available",
			Optional:    true,
		},
		"negate": {
			Type:        hcl.TypeBool,
			Description: "no documentation available",
			Optional:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "Operator of the comparison",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *TextFilter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"case_sensitive": me.CaseSensitive,
		"enabled":        me.Enabled,
		"negate":         me.Negate,
		"operator":       me.Operator,
		"value":          me.Value,
	})
}

func (me *TextFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"case_sensitive": &me.CaseSensitive,
		"enabled":        &me.Enabled,
		"negate":         &me.Negate,
		"operator":       &me.Operator,
		"value":          &me.Value,
	})
}
