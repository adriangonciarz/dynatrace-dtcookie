package resourcecleanuprules

import "github.com/dtcookie/hcl"

type Settings struct {
	Name              string `json:"name"`              // For example: *Mask journeyId*
	RegularExpression string `json:"regularExpression"` // For example: `(.*)(journeyId=)-?\\d+(.*)`
	ReplaceWith       string `json:"replaceWith"`       // For example: `$1$2\\*$3`
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "For example: *Mask journeyId*",
			Required:    true,
		},
		"regular_expression": {
			Type:        hcl.TypeString,
			Description: "For example: `(.*)(journeyId=)-?\\d+(.*)`",
			Required:    true,
		},
		"replace_with": {
			Type:        hcl.TypeString,
			Description: "For example: `$1$2\\*$3`",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"name":               me.Name,
		"regular_expression": me.RegularExpression,
		"replace_with":       me.ReplaceWith,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"name":               &me.Name,
		"regular_expression": &me.RegularExpression,
		"replace_with":       &me.ReplaceWith,
	})
}
