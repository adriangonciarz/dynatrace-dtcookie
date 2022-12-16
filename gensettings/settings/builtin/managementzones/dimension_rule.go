package managementzones

import "github.com/dtcookie/hcl"

type DimensionRule struct {
	AppliesTo  DimensionType       `json:"appliesTo"` // Type
	Conditions DimensionConditions `json:"conditions"`
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
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DimensionConditions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *DimensionRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"applies_to": me.AppliesTo,
		"conditions": me.Conditions,
	})
}

func (me *DimensionRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"applies_to": &me.AppliesTo,
		"conditions": &me.Conditions,
	})
}
