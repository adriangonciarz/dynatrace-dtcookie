package metadata

import "github.com/dtcookie/hcl"

type Settings struct {
	Description       *string            `json:"description"`       // Description
	Dimensions        Dimensions         `json:"dimensions"`        // Define metadata per metric dimension.
	DisplayName       *string            `json:"displayName"`       // Display name
	MetricProperties  *MetricProperties  `json:"metricProperties"`  // Metric properties
	SourceEntityType  *string            `json:"sourceEntityType"`  // Specifies which entity dimension should be used as the primary dimension. The property can only be configured for metrics ingested with the Metrics API.
	Tags              []string           `json:"tags"`              // Tags
	Unit              string             `json:"unit"`              // Unit
	UnitDisplayFormat *UnitDisplayFormat `json:"unitDisplayFormat"` // The raw value is stored in bits or bytes. The user interface can display it in these numeral systems:\n\nBinary: 1 MiB = 1024 KiB = 1,048,576 bytes\n\nDecimal: 1 MB = 1000 kB = 1,000,000 bytes\n\nIf not set, the decimal system is used.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"description": {
			Type:        hcl.TypeString,
			Description: "Description",
			Optional:    true,
		},
		"dimensions": {
			Type:        hcl.TypeList,
			Description: "Define metadata per metric dimension.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Dimensions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"display_name": {
			Type:        hcl.TypeString,
			Description: "Display name",
			Optional:    true,
		},
		"metric_properties": {
			Type:        hcl.TypeList,
			Description: "Metric properties",
			Optional:    true,
			Elem:        &hcl.Resource{Schema: new(MetricProperties).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"source_entity_type": {
			Type:        hcl.TypeString,
			Description: "Specifies which entity dimension should be used as the primary dimension. The property can only be configured for metrics ingested with the Metrics API.",
			Optional:    true,
		},
		"tags": {
			Type:        hcl.TypeSet,
			Description: "Tags",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"unit": {
			Type:        hcl.TypeString,
			Description: "Unit",
			Required:    true,
		},
		"unit_display_format": {
			Type:        hcl.TypeString,
			Description: "The raw value is stored in bits or bytes. The user interface can display it in these numeral systems:\n\nBinary: 1 MiB = 1024 KiB = 1,048,576 bytes\n\nDecimal: 1 MB = 1000 kB = 1,000,000 bytes\n\nIf not set, the decimal system is used.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"description":         me.Description,
		"dimensions":          me.Dimensions,
		"display_name":        me.DisplayName,
		"metric_properties":   me.MetricProperties,
		"source_entity_type":  me.SourceEntityType,
		"tags":                me.Tags,
		"unit":                me.Unit,
		"unit_display_format": me.UnitDisplayFormat,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"description":         &me.Description,
		"dimensions":          &me.Dimensions,
		"display_name":        &me.DisplayName,
		"metric_properties":   &me.MetricProperties,
		"source_entity_type":  &me.SourceEntityType,
		"tags":                &me.Tags,
		"unit":                &me.Unit,
		"unit_display_format": &me.UnitDisplayFormat,
	})
}
