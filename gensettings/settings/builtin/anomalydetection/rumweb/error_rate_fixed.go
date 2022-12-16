package rumweb

import "github.com/dtcookie/hcl"

type ErrorRateFixed struct {
	ErrorRateReqPerMin     float64     `json:"errorRateReqPerMin"`     // To avoid over-alerting for low traffic applications
	ErrorRateSensitivity   Sensitivity `json:"errorRateSensitivity"`   // Sensitivity
	MaxFailureRateIncrease float64     `json:"maxFailureRateIncrease"` // Alert if this custom error rate threshold is exceeded during any 5-minute-period
	MinutesAbnormalState   float64     `json:"minutesAbnormalState"`   // Amount of minutes the observed traffic has to stay in abnormal state before alert
}

func (me *ErrorRateFixed) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"error_rate_req_per_min": {
			Type:        hcl.TypeFloat,
			Description: "To avoid over-alerting for low traffic applications",
			Required:    true,
		},
		"error_rate_sensitivity": {
			Type:        hcl.TypeString,
			Description: "Sensitivity",
			Required:    true,
		},
		"max_failure_rate_increase": {
			Type:        hcl.TypeFloat,
			Description: "Alert if this custom error rate threshold is exceeded during any 5-minute-period",
			Required:    true,
		},
		"minutes_abnormal_state": {
			Type:        hcl.TypeFloat,
			Description: "Amount of minutes the observed traffic has to stay in abnormal state before alert",
			Required:    true,
		},
	}
}

func (me *ErrorRateFixed) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"error_rate_req_per_min":    me.ErrorRateReqPerMin,
		"error_rate_sensitivity":    me.ErrorRateSensitivity,
		"max_failure_rate_increase": me.MaxFailureRateIncrease,
		"minutes_abnormal_state":    me.MinutesAbnormalState,
	})
}

func (me *ErrorRateFixed) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"error_rate_req_per_min":    &me.ErrorRateReqPerMin,
		"error_rate_sensitivity":    &me.ErrorRateSensitivity,
		"max_failure_rate_increase": &me.MaxFailureRateIncrease,
		"minutes_abnormal_state":    &me.MinutesAbnormalState,
	})
}
