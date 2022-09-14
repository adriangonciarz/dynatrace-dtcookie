package alerting

import (
	"github.com/dtcookie/hcl"
)

type TextFilter struct {
	Operator      Operator `json:"operator"`      // Operator of the comparison
	Value         string   `json:"value"`         // The value to compare with
	Negate        bool     `json:"negate"`        // Negate the operator
	Enabled       bool     `json:"enabled"`       // Enable this filter
	CaseSensitive bool     `json:"caseSensitive"` // Case Sensitive comparison of text
}

func (me *TextFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "The filter is enabled (`true`) or disabled (`false`)",
			Optional:    true,
		},
		"negate": {
			Type:        hcl.TypeBool,
			Description: "Reverses the comparison **operator**. For example it turns the **begins with** into **does not begin with**",
			Optional:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "Operator of the comparison.   You can reverse it by setting **negate** to `true`. Possible values are `BEGINS_WITH`, `CONTAINS`, `CONTAINS_REGEX`, `ENDS_WITH` and `EQUALS`",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value to compare to",
			Required:    true,
		},
		"case_insensitive": {
			Type:        hcl.TypeBool,
			Description: "The condition is case sensitive (`false`) or case insensitive (`true`).   If not set, then `false` is used, making the condition case sensitive",
			Optional:    true,
		},
	}
}

func (me *TextFilter) MarshalHCL() (map[string]interface{}, error) {
	return map[string]interface{}{
		"enabled":          me.Enabled,
		"negate":           me.Negate,
		"operator":         string(me.Operator),
		"value":            me.Value,
		"case_insensitive": !me.CaseSensitive,
	}, nil
}

func (me *TextFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	if _, value := decoder.GetChange("enabled"); value != nil {
		me.Enabled = value.(bool)
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		me.Negate = value.(bool)
	}
	if _, value := decoder.GetChange("case_insensitive"); value != nil {
		me.CaseSensitive = !(value.(bool))
	}
	if value, ok := decoder.GetOk("operator"); ok {
		me.Operator = Operator(value.(string))
	}
	if value, ok := decoder.GetOk("value"); ok {
		me.Value = value.(string)
	}
	return nil
}
