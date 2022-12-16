package privacypreferences

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AllowListRules []*AllowListRule

func (me *AllowListRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"playbackMaskingAllowListRule": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(AllowListRule).Schema()},
		},
	}
}

func (me AllowListRules) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me) > 0 {
		entries := []interface{}{}
		for _, entry := range me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["playbackMaskingAllowListRule"] = entries
	}
	return result, nil
}

func (me *AllowListRules) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("playbackMaskingAllowListRule"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(AllowListRule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "playbackMaskingAllowListRule", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type AllowListRule struct {
	AttributeExpression string            `json:"attributeExpression"` // Attribute masking can be applied to web applications that store data within attributes, typically data-NAME attributes in HTML5. When you define attributes, their values are masked while recording but not removed.
	CssExpression       string            `json:"cssExpression"`       // Content masking can be applied to webpages where personal data is displayed. When content masking is applied to parent elements, all child elements are masked by default.
	Target              MaskingTargetType `json:"target"`              // Choose the masking rule target type
}

func (me *AllowListRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"attribute_expression": {
			Type:        hcl.TypeString,
			Description: "Attribute masking can be applied to web applications that store data within attributes, typically data-NAME attributes in HTML5. When you define attributes, their values are masked while recording but not removed.",
			Required:    true,
		},
		"css_expression": {
			Type:        hcl.TypeString,
			Description: "Content masking can be applied to webpages where personal data is displayed. When content masking is applied to parent elements, all child elements are masked by default.",
			Required:    true,
		},
		"target": {
			Type:        hcl.TypeString,
			Description: "Choose the masking rule target type",
			Required:    true,
		},
	}
}

func (me *AllowListRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"attribute_expression": me.AttributeExpression,
		"css_expression":       me.CssExpression,
		"target":               me.Target,
	})
}

func (me *AllowListRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute_expression": &me.AttributeExpression,
		"css_expression":       &me.CssExpression,
		"target":               &me.Target,
	})
}
