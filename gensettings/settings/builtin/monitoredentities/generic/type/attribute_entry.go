package generictype

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AttributeEntries []*AttributeEntry

func (me *AttributeEntries) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"attribute": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(AttributeEntry).Schema()},
		},
	}
}

func (me AttributeEntries) MarshalHCL() (map[string]interface{}, error) {
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
		result["attribute"] = entries
	}
	return result, nil
}

func (me *AttributeEntries) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("attribute"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(AttributeEntry)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "attribute", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

// Attribute entry. Describe how an attribute is extracted from ingest data.
type AttributeEntry struct {
	DisplayName string `json:"displayName"` // The human readable attribute name for this extraction rule. Leave blank to use the key as the display name.
	Key         string `json:"key"`         // The attribute key is the unique name of the attribute.
	Pattern     string `json:"pattern"`     // Pattern for specifying the value for the extracted attribute. Can be a static value, placeholders or a combination of both.
}

func (me *AttributeEntry) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"display_name": {
			Type:        hcl.TypeString,
			Description: "The human readable attribute name for this extraction rule. Leave blank to use the key as the display name.",
			Optional:    true,
		},
		"key": {
			Type:        hcl.TypeString,
			Description: "The attribute key is the unique name of the attribute.",
			Required:    true,
		},
		"pattern": {
			Type:        hcl.TypeString,
			Description: "Pattern for specifying the value for the extracted attribute. Can be a static value, placeholders or a combination of both.",
			Required:    true,
		},
	}
}

func (me *AttributeEntry) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"display_name": me.DisplayName,
		"key":          me.Key,
		"pattern":      me.Pattern,
	})
}

func (me *AttributeEntry) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"display_name": &me.DisplayName,
		"key":          &me.Key,
		"pattern":      &me.Pattern,
	})
}
