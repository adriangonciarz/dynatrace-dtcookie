package services

import "github.com/dtcookie/hcl"

type OverAlertingProtection struct {
	MinutesAbnormalState int     `json:"minutesAbnormalState"` // Only alert if the abnormal state remains for at least
	RequestsPerMinute    float64 `json:"requestsPerMinute"`    // Only alert if there are at least
}

func (me *OverAlertingProtection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"minutes_abnormal_state": {
			Type:        hcl.TypeInt,
			Description: "Only alert if the abnormal state remains for at least",
			Required:    true,
		},
		"requests_per_minute": {
			Type:        hcl.TypeFloat,
			Description: "Only alert if there are at least",
			Required:    true,
		},
	}
}

func (me *OverAlertingProtection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"minutes_abnormal_state": me.MinutesAbnormalState,
		"requests_per_minute":    me.RequestsPerMinute,
	})
}

func (me *OverAlertingProtection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"minutes_abnormal_state": &me.MinutesAbnormalState,
		"requests_per_minute":    &me.RequestsPerMinute,
	})
}
