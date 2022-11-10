package gc

import "github.com/dtcookie/hcl"

// Thresholds Custom thresholds for high GC activity. If not set, automatic mode is used.
//  Meeting **any** of these conditions triggers an alert.
type Thresholds struct {
	GcSuspensionPercentage int32 `json:"gcSuspensionPercentage"` // GC suspension is higher than *X*% in 3 out of 5 samples.
	GcTimePercentage       int32 `json:"gcTimePercentage"`       // GC time is higher than *X*% in 3 out of 5 samples.
}

func (me *Thresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"suspension_percentage": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "GC suspension is higher than *X*% in 3 out of 5 samples",
		},
		"time_percentage": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "GC time is higher than *X*% in 3 out of 5 samples",
		},
	}
}

func (me *Thresholds) MarshalHCL() (map[string]interface{}, error) {
	return map[string]interface{}{
		"time_percentage":       int(me.GcSuspensionPercentage),
		"suspension_percentage": int(me.GcTimePercentage),
	}, nil
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("time_percentage"); ok {
		me.GcSuspensionPercentage = int32(value.(int))
	}
	if value, ok := decoder.GetOk("suspension_percentage"); ok {
		me.GcTimePercentage = int32(value.(int))
	}
	return nil
}
