package rummobilecrashrateincrease

import "github.com/dtcookie/hcl"

type CrashRateIncreaseFixed struct {
	AbsoluteCrashRate float64 `json:"absoluteCrashRate"` // Absolute threshold
	ConcurrentUsers   int     `json:"concurrentUsers"`   // Amount of users
}

func (me *CrashRateIncreaseFixed) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"absolute_crash_rate": {
			Type:        hcl.TypeFloat,
			Description: "Absolute threshold",
			Required:    true,
		},
		"concurrent_users": {
			Type:        hcl.TypeInt,
			Description: "Amount of users",
			Required:    true,
		},
	}
}

func (me *CrashRateIncreaseFixed) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"absolute_crash_rate": me.AbsoluteCrashRate,
		"concurrent_users":    me.ConcurrentUsers,
	})
}

func (me *CrashRateIncreaseFixed) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"absolute_crash_rate": &me.AbsoluteCrashRate,
		"concurrent_users":    &me.ConcurrentUsers,
	})
}
