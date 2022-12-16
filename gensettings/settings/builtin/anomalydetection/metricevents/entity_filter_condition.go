package metricevents

import (
	"github.com/dtcookie/hcl"
)

type EntityFilterConditions []*EntityFilterCondition

func (me *EntityFilterConditions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(EntityFilterCondition).Schema()},
		},
	}
}

func (me EntityFilterConditions) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("condition", me)
}

func (me *EntityFilterConditions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("condition", me)
}

type EntityFilterCondition struct {
	Operator EntityFilterOperator `json:"operator"`
	Type     EntityFilterType     `json:"type"` // Filter type
	Value    string               `json:"value"`
}

func (me *EntityFilterCondition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"operator": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "Filter type",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *EntityFilterCondition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"operator": me.Operator,
		"type":     me.Type,
		"value":    me.Value,
	})
}

func (me *EntityFilterCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"operator": &me.Operator,
		"type":     &me.Type,
		"value":    &me.Value,
	})
}
