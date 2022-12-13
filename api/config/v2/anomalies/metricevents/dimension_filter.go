package metricevents

import "github.com/dtcookie/hcl"

type DimensionFilters []*DimensionFilter

func (me *DimensionFilters) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"filter": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "A conditions for the metric usage",
			Elem:        &hcl.Resource{Schema: new(DimensionFilter).Schema()},
		},
	}
}

func (me DimensionFilters) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("filter", me)
}

func (me *DimensionFilters) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("filter", me)
}

type DimensionFilter struct {
	DimensionKey   string `json:"dimensionKey"`
	DimensionValue string `json:"dimensionValue"`
}

func (me *DimensionFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dimension_key": {
			Type:        hcl.TypeString,
			Description: "",
			Required:    true,
		},
		"dimension_value": {
			Type:        hcl.TypeString,
			Description: "",
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
	return decoder.DecodeAll(map[string]interface{}{
		"dimension_key":   &me.DimensionKey,
		"dimension_value": &me.DimensionValue,
	})
}
