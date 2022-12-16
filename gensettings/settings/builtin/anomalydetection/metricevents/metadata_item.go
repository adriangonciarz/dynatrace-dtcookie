package metricevents

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type MetadataItems []*MetadataItem

func (me *MetadataItems) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"item": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(MetadataItem).Schema()},
		},
	}
}

func (me MetadataItems) MarshalHCL() (map[string]interface{}, error) {
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
		result["item"] = entries
	}
	return result, nil
}

func (me *MetadataItems) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("item"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(MetadataItem)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "item", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type MetadataItem struct {
	MetadataKey   string `json:"metadataKey"` // Type 'dt.' for key hints.
	MetadataValue string `json:"metadataValue"`
}

func (me *MetadataItem) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"metadata_key": {
			Type:        hcl.TypeString,
			Description: "Type 'dt.' for key hints.",
			Required:    true,
		},
		"metadata_value": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *MetadataItem) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"metadata_key":   me.MetadataKey,
		"metadata_value": me.MetadataValue,
	})
}

func (me *MetadataItem) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"metadata_key":   &me.MetadataKey,
		"metadata_value": &me.MetadataValue,
	})
}
