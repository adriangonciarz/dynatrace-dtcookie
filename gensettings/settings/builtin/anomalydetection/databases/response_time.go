package databases

import "github.com/dtcookie/hcl"

type ResponseTime struct {
	AutoDetection  *ResponseTimeAuto  `json:"autoDetection"`
	DetectionMode  DetectionMode      `json:"detectionMode"` // Detection mode for response time degradations
	Enabled        bool               `json:"enabled"`       // Detect response time degradations
	FixedDetection *ResponseTimeFixed `json:"fixedDetection"`
}

func (me *ResponseTime) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"auto_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ResponseTimeAuto).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for response time degradations",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect response time degradations",
			Optional:    true,
		},
		"fixed_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ResponseTimeFixed).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *ResponseTime) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"auto_detection":  me.AutoDetection,
		"detection_mode":  me.DetectionMode,
		"enabled":         me.Enabled,
		"fixed_detection": me.FixedDetection,
	})
}

func (me *ResponseTime) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"auto_detection":  &me.AutoDetection,
		"detection_mode":  &me.DetectionMode,
		"enabled":         &me.Enabled,
		"fixed_detection": &me.FixedDetection,
	})
}
