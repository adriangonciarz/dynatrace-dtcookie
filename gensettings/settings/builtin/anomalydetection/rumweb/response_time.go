package rumweb

import "github.com/dtcookie/hcl"

type ResponseTime struct {
	DetectionMode     DetectionMode      `json:"detectionMode"` // Detection strategy for key performance metric degradations
	Enabled           bool               `json:"enabled"`       // Detect key performance metric degradations
	ResponseTimeAuto  *ResponseTimeAuto  `json:"responseTimeAuto"`
	ResponseTimeFixed *ResponseTimeFixed `json:"responseTimeFixed"`
}

func (me *ResponseTime) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection strategy for key performance metric degradations",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect key performance metric degradations",
			Optional:    true,
		},
		"response_time_auto": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ResponseTimeAuto).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"response_time_fixed": {
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
		"detection_mode":      me.DetectionMode,
		"enabled":             me.Enabled,
		"response_time_auto":  me.ResponseTimeAuto,
		"response_time_fixed": me.ResponseTimeFixed,
	})
}

func (me *ResponseTime) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"detection_mode":      &me.DetectionMode,
		"enabled":             &me.Enabled,
		"response_time_auto":  &me.ResponseTimeAuto,
		"response_time_fixed": &me.ResponseTimeFixed,
	})
}
