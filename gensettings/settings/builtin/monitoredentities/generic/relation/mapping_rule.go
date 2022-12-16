package relation

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type MappingRules []*MappingRule

func (me *MappingRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"mappingRule": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(MappingRule).Schema()},
		},
	}
}

func (me MappingRules) MarshalHCL() (map[string]interface{}, error) {
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
		result["mappingRule"] = entries
	}
	return result, nil
}

func (me *MappingRules) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("mappingRule"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(MappingRule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "mappingRule", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type MappingRule struct {
	DestinationProperty       string        `json:"destinationProperty"`       // The case-sensitive name of a property of the destination type.
	DestinationTransformation Normalization `json:"destinationTransformation"` // Normalize text or leave it as-is?
	SourceProperty            string        `json:"sourceProperty"`            // The case-sensitive name of a property of the source type.
	SourceTransformation      Normalization `json:"sourceTransformation"`      // Normalize text or leave it as-is?
}

func (me *MappingRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"destination_property": {
			Type:        hcl.TypeString,
			Description: "The case-sensitive name of a property of the destination type.",
			Required:    true,
		},
		"destination_transformation": {
			Type:        hcl.TypeString,
			Description: "Normalize text or leave it as-is?",
			Required:    true,
		},
		"source_property": {
			Type:        hcl.TypeString,
			Description: "The case-sensitive name of a property of the source type.",
			Required:    true,
		},
		"source_transformation": {
			Type:        hcl.TypeString,
			Description: "Normalize text or leave it as-is?",
			Required:    true,
		},
	}
}

func (me *MappingRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"destination_property":       me.DestinationProperty,
		"destination_transformation": me.DestinationTransformation,
		"source_property":            me.SourceProperty,
		"source_transformation":      me.SourceTransformation,
	})
}

func (me *MappingRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"destination_property":       &me.DestinationProperty,
		"destination_transformation": &me.DestinationTransformation,
		"source_property":            &me.SourceProperty,
		"source_transformation":      &me.SourceTransformation,
	})
}
