package rumcustomcrashrateincrease

import "github.com/dtcookie/hcl"

type CrashRateIncrease struct {
	CrashRateIncreaseAuto  *CrashRateIncreaseAuto  `json:"crashRateIncreaseAuto"`  // Alert crash rate increases when auto-detected baseline is exceeded by a certain number of users
	CrashRateIncreaseFixed *CrashRateIncreaseFixed `json:"crashRateIncreaseFixed"` // Alert crash rate increases when the defined threshold is exceeded by a certain number of users
	DetectionMode          DetectionMode           `json:"detectionMode"`          // Detection strategy for crash rate increases
	Enabled                bool                    `json:"enabled"`                // Detect crash rate increase
}

func (me *CrashRateIncrease) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"crash_rate_increase_auto": {
			Type:        hcl.TypeList,
			Description: "Alert crash rate increases when auto-detected baseline is exceeded by a certain number of users",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CrashRateIncreaseAuto).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"crash_rate_increase_fixed": {
			Type:        hcl.TypeList,
			Description: "Alert crash rate increases when the defined threshold is exceeded by a certain number of users",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CrashRateIncreaseFixed).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection strategy for crash rate increases",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect crash rate increase",
			Optional:    true,
		},
	}
}

func (me *CrashRateIncrease) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"crash_rate_increase_auto":  me.CrashRateIncreaseAuto,
		"crash_rate_increase_fixed": me.CrashRateIncreaseFixed,
		"detection_mode":            me.DetectionMode,
		"enabled":                   me.Enabled,
	})
}

func (me *CrashRateIncrease) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"crash_rate_increase_auto":  &me.CrashRateIncreaseAuto,
		"crash_rate_increase_fixed": &me.CrashRateIncreaseFixed,
		"detection_mode":            &me.DetectionMode,
		"enabled":                   &me.Enabled,
	})
}
