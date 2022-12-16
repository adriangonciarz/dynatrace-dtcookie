package schemalesslogmetric

import "github.com/dtcookie/hcl"

type Settings struct {
	Dimensions       []string `json:"dimensions"`
	Enabled          bool     `json:"enabled"`          // Enabled
	Key              string   `json:"key"`              // Key
	Measure          Measure  `json:"measure"`          // Measure
	MeasureAttribute string   `json:"measureAttribute"` // Attribute
	Query            string   `json:"query"`            // Query
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dimensions": {
			Type:        hcl.TypeSet,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"key": {
			Type:        hcl.TypeString,
			Description: "Key",
			Required:    true,
		},
		"measure": {
			Type:        hcl.TypeString,
			Description: "Measure",
			Required:    true,
		},
		"measure_attribute": {
			Type:        hcl.TypeString,
			Description: "Attribute",
			Required:    true,
		},
		"query": {
			Type:        hcl.TypeString,
			Description: "Query",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"dimensions":        me.Dimensions,
		"enabled":           me.Enabled,
		"key":               me.Key,
		"measure":           me.Measure,
		"measure_attribute": me.MeasureAttribute,
		"query":             me.Query,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"dimensions":        &me.Dimensions,
		"enabled":           &me.Enabled,
		"key":               &me.Key,
		"measure":           &me.Measure,
		"measure_attribute": &me.MeasureAttribute,
		"query":             &me.Query,
	})
}
