package performancethresholds

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled    bool             `json:"enabled"`    // Generate a problem and send an alert on performance threshold violations
	Thresholds ThresholdEntries `json:"thresholds"` // Performance thresholds
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Generate a problem and send an alert on performance threshold violations",
			Optional:    true,
		},
		"thresholds": {
			Type:        hcl.TypeList,
			Description: "Performance thresholds",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ThresholdEntries).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":    me.Enabled,
		"thresholds": me.Thresholds,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":    &me.Enabled,
		"thresholds": &me.Thresholds,
	})
}
