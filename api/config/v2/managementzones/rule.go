package managementzones

import "github.com/dtcookie/hcl"

type Rules []*Rule

func (me *Rules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rule": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "A management zone rule",
			Elem:        &hcl.Resource{Schema: new(Rule).Schema()},
		},
	}
}

func (me Rules) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("rule", me)
}

func (me *Rules) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("rule", me); err != nil {
		return err
	}
	return nil
}

// No documentation available
type Rule struct {
	Enabled        bool                         `json:"enabled" hcl:"enabled"`                // Enabled
	Type           RuleType                     `json:"type" hcl:"type"`                      // Rule type
	AttributeRule  *ManagementZoneAttributeRule `json:"attributeRule" hcl:"attribute_rule"`   // No documentation available
	DimensionRule  *DimensionRule               `json:"dimensionRule" hcl:"dimension_rule"`   // No documentation available
	EntitySelector string                       `json:"entitySelector" hcl:"entity_selector"` // Entity selector. The documentation of the entity selector can be found [here](https://dt-url.net/apientityselector).
}

func (me *Rule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
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
			Required:    true,
		},
		"dimension_rule": {
			Type:        hcl.TypeList,
			Description: "No documentation available",
			MaxItems:    1,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(DimensionRule).Schema()},
			Required:    true,
		},
		"entity_selector": {
			Type:        hcl.TypeString,
			Description: "Entity selector. The documentation of the entity selector can be found [here](https://dt-url.net/apientityselector).",
			Required:    true,
		},
	}
}
