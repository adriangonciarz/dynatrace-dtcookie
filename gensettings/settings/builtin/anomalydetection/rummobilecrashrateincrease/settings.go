package rummobilecrashrateincrease

import "github.com/dtcookie/hcl"

type Settings struct {
	CrashRateIncrease *CrashRateIncrease `json:"crashRateIncrease"` // Crash rate increase
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"crash_rate_increase": {
			Type:        hcl.TypeList,
			Description: "Crash rate increase",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CrashRateIncrease).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"crash_rate_increase": me.CrashRateIncrease,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"crash_rate_increase": &me.CrashRateIncrease,
	})
}
