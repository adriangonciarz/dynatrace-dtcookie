package rule

import "github.com/dtcookie/hcl"

type Settings struct {
	Dimensions       []string `json:"dimensions"`
	Enabled          bool     `json:"enabled"`          // Enabled
	Key              string   `json:"key"`              // Key
	Matcher          string   `json:"matcher"`          // [See our documentation](https://dt-url.net/bp234rv)
	Measure          Measure  `json:"measure"`          // Measure
	MeasureAttribute string   `json:"measureAttribute"` // Attribute
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
		"matcher": {
			Type:        hcl.TypeString,
			Description: "[See our documentation](https://dt-url.net/bp234rv)",
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
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"dimensions":        me.Dimensions,
		"enabled":           me.Enabled,
		"key":               me.Key,
		"matcher":           me.Matcher,
		"measure":           me.Measure,
		"measure_attribute": me.MeasureAttribute,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"dimensions":        &me.Dimensions,
		"enabled":           &me.Enabled,
		"key":               &me.Key,
		"matcher":           &me.Matcher,
		"measure":           &me.Measure,
		"measure_attribute": &me.MeasureAttribute,
	})
}
