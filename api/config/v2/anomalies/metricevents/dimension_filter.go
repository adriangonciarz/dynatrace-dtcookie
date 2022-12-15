package metricevents

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DimensionFilters []*DimensionFilter // Dimension filter definitions

func (me *DimensionFilters) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"filter": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "Dimension filter definitions",
			Elem:        &hcl.Resource{Schema: new(DimensionFilter).Schema()},
		},
	}
}

func (me DimensionFilters) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me) > 0 {
		entries := []interface{}{}
		for _, entry := range me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["filter"] = entries
	}
	return result, nil
}

func (me *DimensionFilters) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("filter"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(DimensionFilter)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "filter", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type DimensionFilter struct {
	DimensionKey   string `json:"dimensionKey"`   // The key of the dimension filter
	DimensionValue string `json:"dimensionValue"` // The value of the dimension filter
}

func (me *DimensionFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dimension_key": {
			Type:        hcl.TypeString,
			Description: "The key of the dimension filter",
			Required:    true,
		},
		"dimension_value": {
			Type:        hcl.TypeString,
			Description: "The value of the dimension filter",
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
