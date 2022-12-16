package relation

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SourceFilters []*SourceFilter

func (me *SourceFilters) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"source": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(SourceFilter).Schema()},
		},
	}
}

func (me SourceFilters) MarshalHCL() (map[string]interface{}, error) {
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
		result["source"] = entries
	}
	return result, nil
}

func (me *SourceFilters) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("source"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(SourceFilter)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "source", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

// Source filter. The source filter determines based on which data the relationship should be created. This way a subset of a specified data source can be used for creating the type.
type SourceFilter struct {
	Condition    string           `json:"condition"`    // Specify a filter that needs to match in order for the extraction to happen.. Two different filters are supported: `$eq(value)` will ensure that the source matches exactly 'value', while `$prefix(value)` will ensure that the source begins with exactly 'value'.\nIf your value contains the characters '(', ')' or '\\~', you need to escape them by adding a '\\~' in front of them.
	MappingRules MappingRules     `json:"mappingRules"` // Specify all properties which should be compared. If all mapping rules match a relationship between entities will be created.
	SourceType   IngestDataSource `json:"sourceType"`   // Specify the source type of the filter to identify which data source should be evaluated.
}

func (me *SourceFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeString,
			Description: "Specify a filter that needs to match in order for the extraction to happen.. Two different filters are supported: `$eq(value)` will ensure that the source matches exactly 'value', while `$prefix(value)` will ensure that the source begins with exactly 'value'.\nIf your value contains the characters '(', ')' or '\\~', you need to escape them by adding a '\\~' in front of them.",
			Required:    true,
		},
		"mapping_rules": {
			Type:        hcl.TypeList,
			Description: "Specify all properties which should be compared. If all mapping rules match a relationship between entities will be created.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MappingRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"source_type": {
			Type:        hcl.TypeString,
			Description: "Specify the source type of the filter to identify which data source should be evaluated.",
			Required:    true,
		},
	}
}

func (me *SourceFilter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"condition":     me.Condition,
		"mapping_rules": me.MappingRules,
		"source_type":   me.SourceType,
	})
}

func (me *SourceFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"condition":     &me.Condition,
		"mapping_rules": &me.MappingRules,
		"source_type":   &me.SourceType,
	})
}
