package custommetrics

import "github.com/dtcookie/hcl"

type MetricValue struct {
	FieldName string    `json:"fieldName"` // Field name
	Type      ValueType `json:"type"`
}

func (me *MetricValue) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"field_name": {
			Type:        hcl.TypeString,
			Description: "Field name",
			Required:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *MetricValue) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"field_name": me.FieldName,
		"type":       me.Type,
	})
}

func (me *MetricValue) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"field_name": &me.FieldName,
		"type":       &me.Type,
	})
}
