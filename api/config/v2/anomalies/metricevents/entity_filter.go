package metricevents

import (
	"github.com/dtcookie/hcl"
)

type EntityFilter struct {
	DimensionKey *string                `json:"dimensionKey,omitempty"` // Dimension key of entity type to filter
	Conditions   EntityFilterConditions `json:"conditions,omitempty"`   // Conditions of entity type to filter
}

func (me *EntityFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dimension_key": {
			Type:        hcl.TypeString,
			Description: "Dimension key of entity type to filter",
			Optional:    true,
		},
		"conditions": {
			Type:        hcl.TypeList,
			Description: "Conditions of entity type to filter",
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(EntityFilterConditions).Schema()},
			Optional:    true,
		},
	}
}

func (me *EntityFilter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"dimension_key": me.DimensionKey,
		"conditions":    me.Conditions,
	})
}

func (me *EntityFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"dimension_key": &me.DimensionKey,
		"conditions":    &me.Conditions,
	})
}
