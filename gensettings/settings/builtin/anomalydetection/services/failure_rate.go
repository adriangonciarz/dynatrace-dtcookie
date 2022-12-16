package services

import "github.com/dtcookie/hcl"

type FailureRate struct {
	AutoDetection  *FailureRateAuto  `json:"autoDetection"`  // Alert if the percentage of failing service calls increases by **both** the absolute and relative thresholds:
	DetectionMode  DetectionMode     `json:"detectionMode"`  // Detection mode for increases in failure rate
	Enabled        bool              `json:"enabled"`        // Detect increases in failure rate
	FixedDetection *FailureRateFixed `json:"fixedDetection"` // Alert if a given failure rate is exceeded during any 5-minute-period
}

func (me *FailureRate) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"auto_detection": {
			Type:        hcl.TypeList,
			Description: "Alert if the percentage of failing service calls increases by **both** the absolute and relative thresholds:",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(FailureRateAuto).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for increases in failure rate",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect increases in failure rate",
			Optional:    true,
		},
		"fixed_detection": {
			Type:        hcl.TypeList,
			Description: "Alert if a given failure rate is exceeded during any 5-minute-period",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(FailureRateFixed).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *FailureRate) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"auto_detection":  me.AutoDetection,
		"detection_mode":  me.DetectionMode,
		"enabled":         me.Enabled,
		"fixed_detection": me.FixedDetection,
	})
}

func (me *FailureRate) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"auto_detection":  &me.AutoDetection,
		"detection_mode":  &me.DetectionMode,
		"enabled":         &me.Enabled,
		"fixed_detection": &me.FixedDetection,
	})
}
