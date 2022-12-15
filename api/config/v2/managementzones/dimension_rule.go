package managementzones

import "github.com/dtcookie/hcl"

// No documentation available
type DimensionRule struct {
	AppliesTo  DimensionType       `json:"appliesTo"`            // Type
	Conditions DimensionConditions `json:"conditions,omitempty"` // Conditions
}

func (me *DimensionRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"applies_to": {
			Type:        hcl.TypeString,
			Description: "Type",
			Required:    true,
		},
		"dimension_conditions": {
			Type:        hcl.TypeList,
			Description: "Conditions",
			MaxItems:    1,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(DimensionConditions).Schema()},
			Optional:    true,
		},
	}
}

func (me *DimensionRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"applies_to":           me.AppliesTo,
		"dimension_conditions": me.Conditions,
	})
}

func (me *DimensionRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"applies_to":           &me.AppliesTo,
		"dimension_conditions": &me.Conditions,
	})
}
