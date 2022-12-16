package infrastructureaws

import "github.com/dtcookie/hcl"

type Ec2CandidateHighCpuDetectionConfig struct {
	CustomThresholds *Ec2CandidateHighCpuDetectionThresholds `json:"customThresholds"` // Alert if the condition is met in 3 out of 5 samples
	DetectionMode    DetectionMode                           `json:"detectionMode"`    // Detection mode
	Enabled          bool                                    `json:"enabled"`          // Detect high CPU saturation on EC2 monitoring candidate
}

func (me *Ec2CandidateHighCpuDetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if the condition is met in 3 out of 5 samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Ec2CandidateHighCpuDetectionThresholds).Schema()},
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
			Description: "Detect high CPU saturation on EC2 monitoring candidate",
			Optional:    true,
		},
	}
}

func (me *Ec2CandidateHighCpuDetectionConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *Ec2CandidateHighCpuDetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
