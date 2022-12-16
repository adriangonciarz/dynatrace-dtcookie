package generictype

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DimensionFilters []*DimensionFilter

func (me *DimensionFilters) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"requiredDimension": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
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
		result["requiredDimension"] = entries
	}
	return result, nil
}

func (me *DimensionFilters) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("requiredDimension"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(DimensionFilter)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "requiredDimension", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

// Ingest dimension filter. A dimension describes a property key which is present in the ingest data.
type DimensionFilter struct {
	Key          string `json:"key"`          // A dimension key which needs to exist in the ingest data to match this filter.
	ValuePattern string `json:"valuePattern"` // A dimension value pattern which needs to exist in the ingest data to match this filter.
}

func (me *DimensionFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key": {
			Type:        hcl.TypeString,
			Description: "A dimension key which needs to exist in the ingest data to match this filter.",
			Required:    true,
		},
		"value_pattern": {
			Type:        hcl.TypeString,
			Description: "A dimension value pattern which needs to exist in the ingest data to match this filter.",
			Optional:    true,
		},
	}
}

func (me *DimensionFilter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"key":           me.Key,
		"value_pattern": me.ValuePattern,
	})
}

func (me *DimensionFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"key":           &me.Key,
		"value_pattern": &me.ValuePattern,
	})
}
