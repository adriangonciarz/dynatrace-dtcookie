package rummobile

import "github.com/dtcookie/hcl"

type Settings struct {
	ErrorRateIncrease  *ErrorRateIncrease  `json:"errorRateIncrease"`  // Error rate increase
	SlowUserActions    *SlowUserActions    `json:"slowUserActions"`    // Slow user actions
	UnexpectedHighLoad *UnexpectedHighLoad `json:"unexpectedHighLoad"` // Unexpected high load
	UnexpectedLowLoad  *UnexpectedLowLoad  `json:"unexpectedLowLoad"`  // Unexpected low load
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"error_rate_increase": {
			Type:        hcl.TypeList,
			Description: "Error rate increase",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ErrorRateIncrease).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"slow_user_actions": {
			Type:        hcl.TypeList,
			Description: "Slow user actions",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SlowUserActions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"unexpected_high_load": {
			Type:        hcl.TypeList,
			Description: "Unexpected high load",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(UnexpectedHighLoad).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"unexpected_low_load": {
			Type:        hcl.TypeList,
			Description: "Unexpected low load",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(UnexpectedLowLoad).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"error_rate_increase":  me.ErrorRateIncrease,
		"slow_user_actions":    me.SlowUserActions,
		"unexpected_high_load": me.UnexpectedHighLoad,
		"unexpected_low_load":  me.UnexpectedLowLoad,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"error_rate_increase":  &me.ErrorRateIncrease,
		"slow_user_actions":    &me.SlowUserActions,
		"unexpected_high_load": &me.UnexpectedHighLoad,
		"unexpected_low_load":  &me.UnexpectedLowLoad,
	})
}
