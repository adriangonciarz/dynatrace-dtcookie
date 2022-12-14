package metricevents

import "github.com/dtcookie/hcl"

type EntityFilterConditions []*EntityFilterCondition // Entity filter conditions

func (me *EntityFilterConditions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "Entity filter conditions",
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
	Type     EntityFilterType     `json:"type"`
	Operator EntityFilterOperator `json:"operator"`
	Value    string               `json:"value"`
}

func (me *EntityFilterCondition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "",
			Required:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "",
			Required:    true,
		},
	}
}

func (me *EntityFilterCondition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"type":     me.Type,
		"operator": me.Operator,
		"value":    me.Value,
	})
}

func (me *EntityFilterCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"type":     &me.Type,
		"operator": &me.Operator,
		"value":    &me.Value,
	})
}
