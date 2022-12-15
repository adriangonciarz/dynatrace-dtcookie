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
			Optional:    true,
			MinItems:    1,
			Description: "A management zone rule",
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

// No documentation available
type Rule struct {
	Enabled        bool                         `json:"enabled"`                  // Enabled
	Type           RuleType                     `json:"type"`                     // Rule type
	AttributeRule  *ManagementZoneAttributeRule `json:"attributeRule,omitempty"`  // No documentation available
	DimensionRule  *DimensionRule               `json:"dimensionRule,omitempty"`  // No documentation available
	EntitySelector string                       `json:"entitySelector,omitempty"` // Entity selector. The documentation of the entity selector can be found [here](https://dt-url.net/apientityselector).
}

func (me *Rule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Required:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "Rule type",
			Required:    true,
		},
		"attribute_rule": {
			Type:        hcl.TypeList,
			Description: "No documentation available",
			MaxItems:    1,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(ManagementZoneAttributeRule).Schema()},
			Optional:    true,
		},
		"dimension_rule": {
			Type:        hcl.TypeList,
			Description: "No documentation available",
			MaxItems:    1,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(DimensionRule).Schema()},
			Optional:    true,
		},
		"entity_selector": {
			Type:        hcl.TypeString,
			Description: "Entity selector. The documentation of the entity selector can be found [here](https://dt-url.net/apientityselector).",
			Optional:    true,
		},
	}
}

func (me *Rule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"enabled":         me.Enabled,
		"type":            me.Type,
		"attribute_rule":  me.AttributeRule,
		"dimension_rule":  me.DimensionRule,
		"entity_selector": me.EntitySelector,
	})
}

func (me *Rule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"enabled":         &me.Enabled,
		"type":            &me.Type,
		"attribute_rule":  &me.AttributeRule,
		"dimension_rule":  &me.DimensionRule,
		"entity_selector": &me.EntitySelector,
	})
}
