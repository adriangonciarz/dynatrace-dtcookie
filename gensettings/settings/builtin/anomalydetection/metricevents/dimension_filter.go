package metricevents

import (
	"github.com/dtcookie/hcl"
)

type DimensionFilters []*DimensionFilter

func (me *DimensionFilters) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dimension_filter": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(DimensionFilter).Schema()},
		},
	}
}

func (me DimensionFilters) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("dimension_filter", me)
}

func (me *DimensionFilters) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("dimension_filter", me)
}

type DimensionFilter struct {
	DimensionKey   string `json:"dimensionKey"`   // Dimension key
	DimensionValue string `json:"dimensionValue"` // Dimension value
}

func (me *DimensionFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dimension_key": {
			Type:        hcl.TypeString,
			Description: "Dimension key",
			Required:    true,
		},
		"dimension_value": {
			Type:        hcl.TypeString,
			Description: "Dimension value",
			Required:    true,
		},
	}
}

func (me *DimensionFilter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"dimension_key":   me.DimensionKey,
		"dimension_value": me.DimensionValue,
	})
}

func (me *DimensionFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"dimension_key":   &me.DimensionKey,
		"dimension_value": &me.DimensionValue,
	})
}
