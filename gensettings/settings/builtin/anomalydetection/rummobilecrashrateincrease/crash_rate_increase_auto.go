package rummobilecrashrateincrease

import "github.com/dtcookie/hcl"

type CrashRateIncreaseAuto struct {
	BaselineViolationPercentage float64     `json:"baselineViolationPercentage"` // Dynatrace learns the typical crash rate for all app versions and will create an alert if the baseline is violated by more than a specified threshold. Analysis happens based on a sliding window of 10 minutes.
	ConcurrentUsers             float64     `json:"concurrentUsers"`             // Amount of users
	Sensitivity                 Sensitivity `json:"sensitivity"`                 // Detection sensitivity
}

func (me *CrashRateIncreaseAuto) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"baseline_violation_percentage": {
			Type:        hcl.TypeFloat,
			Description: "Dynatrace learns the typical crash rate for all app versions and will create an alert if the baseline is violated by more than a specified threshold. Analysis happens based on a sliding window of 10 minutes.",
			Required:    true,
		},
		"concurrent_users": {
			Type:        hcl.TypeFloat,
			Description: "Amount of users",
			Required:    true,
		},
		"sensitivity": {
			Type:        hcl.TypeString,
			Description: "Detection sensitivity",
			Required:    true,
		},
	}
}

func (me *CrashRateIncreaseAuto) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"baseline_violation_percentage": me.BaselineViolationPercentage,
		"concurrent_users":              me.ConcurrentUsers,
		"sensitivity":                   me.Sensitivity,
	})
}

func (me *CrashRateIncreaseAuto) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"baseline_violation_percentage": &me.BaselineViolationPercentage,
		"concurrent_users":              &me.ConcurrentUsers,
		"sensitivity":                   &me.Sensitivity,
	})
}
