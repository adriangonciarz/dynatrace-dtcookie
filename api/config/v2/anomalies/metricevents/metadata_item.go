package metricevents

import "github.com/dtcookie/hcl"

type MetadataItem struct {
	MetadataKey   string `json:"metadataKey"`
	MetadataValue string `json:"metadataValue"`
}

func (me *MetadataItem) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"metadata_key": {
			Type:        hcl.TypeString,
			Description: "",
			Required:    true,
		},
		"metadata_value": {
			Type:        hcl.TypeString,
			Description: "",
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
