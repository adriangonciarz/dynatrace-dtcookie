package metricevents

import "github.com/dtcookie/hcl"

type MetadataItem struct {
	MetadataKey   string `json:"metadataKey"`   // The key of the metadata item
	MetadataValue string `json:"metadataValue"` // The value of the metadata item
}

func (me *MetadataItem) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"metadata_key": {
			Type:        hcl.TypeString,
			Description: "The key of the metadata item",
			Required:    true,
		},
		"metadata_value": {
			Type:        hcl.TypeString,
			Description: "The value of the metadata item",
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
	return decoder.DecodeAll(map[string]interface{}{
		"metadata_key":   &me.MetadataKey,
		"metadata_value": &me.MetadataValue,
	})
}
