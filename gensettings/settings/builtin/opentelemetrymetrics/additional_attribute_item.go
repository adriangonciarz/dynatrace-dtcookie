package opentelemetrymetrics

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AdditionalAttributeItems []*AdditionalAttributeItem

func (me *AdditionalAttributeItems) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"additionalAttribute": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(AdditionalAttributeItem).Schema()},
		},
	}
}

func (me AdditionalAttributeItems) MarshalHCL() (map[string]interface{}, error) {
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
		result["additionalAttribute"] = entries
	}
	return result, nil
}

func (me *AdditionalAttributeItems) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("additionalAttribute"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(AdditionalAttributeItem)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "additionalAttribute", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type AdditionalAttributeItem struct {
	AttributeKey string `json:"attributeKey"` // Attribute key
	Enabled      bool   `json:"enabled"`      // When enabled, the attribute will be automatically added as a dimension to ingested metrics if present in the OpenTelemetry resource or in the instrumentation scope.
}

func (me *AdditionalAttributeItem) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"attribute_key": {
			Type:        hcl.TypeString,
			Description: "Attribute key",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "When enabled, the attribute will be automatically added as a dimension to ingested metrics if present in the OpenTelemetry resource or in the instrumentation scope.",
			Optional:    true,
		},
	}
}

func (me *AdditionalAttributeItem) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"attribute_key": me.AttributeKey,
		"enabled":       me.Enabled,
	})
}

func (me *AdditionalAttributeItem) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute_key": &me.AttributeKey,
		"enabled":       &me.Enabled,
	})
}
