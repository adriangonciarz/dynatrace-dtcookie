package metricevents

import "github.com/dtcookie/hcl"

type EntityFilter struct {
	Conditions   EntityFilterConditions `json:"conditions"`
	DimensionKey string                 `json:"dimensionKey"` // Dimension key of entity type to filter
}

func (me *EntityFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"conditions": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EntityFilterConditions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"dimension_key": {
			Type:        hcl.TypeString,
			Description: "Dimension key of entity type to filter",
			Optional:    true,
		},
	}
}

func (me *EntityFilter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"conditions":    me.Conditions,
		"dimension_key": me.DimensionKey,
	})
}

func (me *EntityFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"conditions":    &me.Conditions,
		"dimension_key": &me.DimensionKey,
	})
}
