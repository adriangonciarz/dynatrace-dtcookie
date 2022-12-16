package privacypreferences

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type BlockListRules []*BlockListRule

func (me *BlockListRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"recordingMaskingBlockListRule": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(BlockListRule).Schema()},
		},
	}
}

func (me BlockListRules) MarshalHCL() (map[string]interface{}, error) {
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
		result["recordingMaskingBlockListRule"] = entries
	}
	return result, nil
}

func (me *BlockListRules) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("recordingMaskingBlockListRule"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(BlockListRule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "recordingMaskingBlockListRule", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type BlockListRule struct {
	AttributeExpression string            `json:"attributeExpression"` // Attribute masking can be applied to web applications that store data within attributes, typically data-NAME attributes in HTML5. When you define attributes, their values are masked while recording but not removed.
	CssExpression       string            `json:"cssExpression"`       // Content masking can be applied to webpages where personal data is displayed. When content masking is applied to parent elements, all child elements are masked by default.
	HideUserInteraction bool              `json:"hideUserInteraction"` // Hide user interactions with these elements, including clicks that expand elements, highlighting that results from hovering a cursor over an option, and selection of specific form options.
	Target              MaskingTargetType `json:"target"`              // Choose the masking rule target type
}

func (me *BlockListRule) Schema() map[string]*hcl.Schema {
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
		"hide_user_interaction": {
			Type:        hcl.TypeBool,
			Description: "Hide user interactions with these elements, including clicks that expand elements, highlighting that results from hovering a cursor over an option, and selection of specific form options.",
			Optional:    true,
		},
		"target": {
			Type:        hcl.TypeString,
			Description: "Choose the masking rule target type",
			Required:    true,
		},
	}
}

func (me *BlockListRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"attribute_expression":  me.AttributeExpression,
		"css_expression":        me.CssExpression,
		"hide_user_interaction": me.HideUserInteraction,
		"target":                me.Target,
	})
}

func (me *BlockListRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute_expression":  &me.AttributeExpression,
		"css_expression":        &me.CssExpression,
		"hide_user_interaction": &me.HideUserInteraction,
		"target":                &me.Target,
	})
}
