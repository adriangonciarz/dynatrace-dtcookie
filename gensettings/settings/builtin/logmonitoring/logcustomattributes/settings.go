package logcustomattributes

import "github.com/dtcookie/hcl"

type Settings struct {
	AggregableAttribute bool   `json:"aggregableAttribute"` // Change applies only to newly ingested log events. Any log events ingested before this option was toggled on will not be searchable by this attribute.
	Key                 string `json:"key"`                 // The attribute key is case insensitive in log data ingestion.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"aggregable_attribute": {
			Type:        hcl.TypeBool,
			Description: "Change applies only to newly ingested log events. Any log events ingested before this option was toggled on will not be searchable by this attribute.",
			Optional:    true,
		},
		"key": {
			Type:        hcl.TypeString,
			Description: "The attribute key is case insensitive in log data ingestion.",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"aggregable_attribute": me.AggregableAttribute,
		"key":                  me.Key,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"aggregable_attribute": &me.AggregableAttribute,
		"key":                  &me.Key,
	})
}
