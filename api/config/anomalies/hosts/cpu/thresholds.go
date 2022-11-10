package cpu

import (
	"github.com/dtcookie/hcl"
)

// Thresholds Custom thresholds for high CPU saturation. If not set then the automatic mode is used.
type Thresholds struct {
	CPUSaturation int32 `json:"cpuSaturation"` // Alert if CPU usage is higher than *X*% in 3 out of 5 samples.
}

func (me *Thresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"saturation": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Alert if CPU usage is higher than *X*% in 3 out of 5 samples",
		},
	}
}

func (me *Thresholds) MarshalHCL() (map[string]interface{}, error) {
	return map[string]interface{}{
		"saturation": int(me.CPUSaturation),
	}, nil
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("saturation"); ok {
		me.CPUSaturation = int32(value.(int))
	}
	return nil
}
