package infrastructurehosts

import "github.com/dtcookie/hcl"

type NetworkTcpProblemsDetection struct {
	CustomThresholds *NetworkTcpProblemsDetectionThresholds `json:"customThresholds"` // Alert if the percentage of new connection failures is higher than the specified threshold **and** the number of failed connections is higher than the defined threshold for the defined amount of samples
	DetectionMode    DetectionMode                          `json:"detectionMode"`    // Detection mode for TCP connectivity problems
	Enabled          bool                                   `json:"enabled"`          // Detect TCP connectivity problems for process
}

func (me *NetworkTcpProblemsDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if the percentage of new connection failures is higher than the specified threshold **and** the number of failed connections is higher than the defined threshold for the defined amount of samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(NetworkTcpProblemsDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for TCP connectivity problems",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect TCP connectivity problems for process",
			Optional:    true,
		},
	}
}

func (me *NetworkTcpProblemsDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *NetworkTcpProblemsDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
