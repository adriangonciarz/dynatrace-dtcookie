package keyperformancemetrics

import "github.com/dtcookie/hcl"

type Thresholds struct {
	FrustratingThresholdSeconds float64 `json:"frustratingThresholdSeconds"` // If the action duration is above this value, the Apdex is considered to be **Frustrating**.
	TolerableThresholdSeconds   float64 `json:"tolerableThresholdSeconds"`   // If the action duration is below this value, the Apdex is considered to be **Satisfactory**.
}

func (me *Thresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"frustrating_threshold_seconds": {
			Type:        hcl.TypeFloat,
			Description: "If the action duration is above this value, the Apdex is considered to be **Frustrating**.",
			Required:    true,
		},
		"tolerable_threshold_seconds": {
			Type:        hcl.TypeFloat,
			Description: "If the action duration is below this value, the Apdex is considered to be **Satisfactory**.",
			Required:    true,
		},
	}
}

func (me *Thresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"frustrating_threshold_seconds": me.FrustratingThresholdSeconds,
		"tolerable_threshold_seconds":   me.TolerableThresholdSeconds,
	})
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"frustrating_threshold_seconds": &me.FrustratingThresholdSeconds,
		"tolerable_threshold_seconds":   &me.TolerableThresholdSeconds,
	})
}
