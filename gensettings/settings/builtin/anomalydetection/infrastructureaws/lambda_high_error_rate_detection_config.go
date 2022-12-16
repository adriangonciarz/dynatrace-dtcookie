package infrastructureaws

import "github.com/dtcookie/hcl"

type LambdaHighErrorRateDetectionConfig struct {
	CustomThresholds *LambdaHighErrorRateDetectionThresholds `json:"customThresholds"` // Alert if the condition is met in 3 out of 5 samples
	DetectionMode    DetectionMode                           `json:"detectionMode"`    // Detection mode
	Enabled          bool                                    `json:"enabled"`          // Detect AWS Lambda high error rate
}

func (me *LambdaHighErrorRateDetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if the condition is met in 3 out of 5 samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(LambdaHighErrorRateDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect AWS Lambda high error rate",
			Optional:    true,
		},
	}
}

func (me *LambdaHighErrorRateDetectionConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *LambdaHighErrorRateDetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
