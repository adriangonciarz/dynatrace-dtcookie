package managementzones

import "github.com/dtcookie/hcl"

// No documentation available
type DimensionRule struct {
	AppliesTo  DimensionType         `json:"appliesTo" hcl:"applies_to"`  // Type
	Conditions []*DimensionCondition `json:"conditions" hcl:"conditions"` // Conditions
}

func (me *DimensionRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"applies_to": {
			Type:        hcl.TypeString,
			Description: "Type",
			Required:    true,
		},
		"conditions": {
			Type:        hcl.TypeList,
			Description: "Conditions",
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(DimensionCondition).Schema()},
			Required:    true,
		},
	}
}
