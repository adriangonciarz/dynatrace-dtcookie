package managementzones

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Rules []*Rule

func (me *Rules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rule": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(Rule).Schema()},
		},
	}
}

func (me Rules) MarshalHCL() (map[string]interface{}, error) {
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
		result["rule"] = entries
	}
	return result, nil
}

func (me *Rules) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("rule"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(Rule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "rule", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type Rule struct {
	AttributeRule  *ManagementZoneAttributeRule `json:"attributeRule"`
	DimensionRule  *DimensionRule               `json:"dimensionRule"`
	Enabled        bool                         `json:"enabled"`
	EntitySelector string                       `json:"entitySelector"` // The documentation of the entity selector can be found [here](https://dt-url.net/apientityselector).
	Type           RuleType                     `json:"type"`           // Rule type
}

func (me *Rule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"attribute_rule": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ManagementZoneAttributeRule).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"dimension_rule": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DimensionRule).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "no documentation available",
			Optional:    true,
		},
		"entity_selector": {
			Type:        hcl.TypeString,
			Description: "The documentation of the entity selector can be found [here](https://dt-url.net/apientityselector).",
			Required:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "Rule type",
			Required:    true,
		},
	}
}

func (me *Rule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"attribute_rule":  me.AttributeRule,
		"dimension_rule":  me.DimensionRule,
		"enabled":         me.Enabled,
		"entity_selector": me.EntitySelector,
		"type":            me.Type,
	})
}

func (me *Rule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute_rule":  &me.AttributeRule,
		"dimension_rule":  &me.DimensionRule,
		"enabled":         &me.Enabled,
		"entity_selector": &me.EntitySelector,
		"type":            &me.Type,
	})
}
