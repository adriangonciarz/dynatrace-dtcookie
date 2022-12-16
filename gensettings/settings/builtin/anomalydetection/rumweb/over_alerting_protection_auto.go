package rumweb

import "github.com/dtcookie/hcl"

type OverAlertingProtectionAuto struct {
	ActionsPerMinute     float64 `json:"actionsPerMinute"`     // Only alert if there are at least
	MinutesAbnormalState float64 `json:"minutesAbnormalState"` // Only alert if the abnormal state remains for at least
}

func (me *OverAlertingProtectionAuto) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"actions_per_minute": {
			Type:        hcl.TypeFloat,
			Description: "Only alert if there are at least",
			Required:    true,
		},
		"minutes_abnormal_state": {
			Type:        hcl.TypeFloat,
			Description: "Only alert if the abnormal state remains for at least",
			Required:    true,
		},
	}
}

func (me *OverAlertingProtectionAuto) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"actions_per_minute":     me.ActionsPerMinute,
		"minutes_abnormal_state": me.MinutesAbnormalState,
	})
}

func (me *OverAlertingProtectionAuto) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"actions_per_minute":     &me.ActionsPerMinute,
		"minutes_abnormal_state": &me.MinutesAbnormalState,
	})
}
