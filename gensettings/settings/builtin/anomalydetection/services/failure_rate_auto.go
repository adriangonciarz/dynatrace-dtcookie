package services

import "github.com/dtcookie/hcl"

type FailureRateAuto struct {
	AbsoluteIncrease       float64                 `json:"absoluteIncrease"`       // Absolute threshold
	OverAlertingProtection *OverAlertingProtection `json:"overAlertingProtection"` // Avoid over-alerting
	RelativeIncrease       float64                 `json:"relativeIncrease"`       // Relative threshold
}

func (me *FailureRateAuto) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"absolute_increase": {
			Type:        hcl.TypeFloat,
			Description: "Absolute threshold",
			Required:    true,
		},
		"over_alerting_protection": {
			Type:        hcl.TypeList,
			Description: "Avoid over-alerting",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(OverAlertingProtection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"relative_increase": {
			Type:        hcl.TypeFloat,
			Description: "Relative threshold",
			Required:    true,
		},
	}
}

func (me *FailureRateAuto) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"absolute_increase":        me.AbsoluteIncrease,
		"over_alerting_protection": me.OverAlertingProtection,
		"relative_increase":        me.RelativeIncrease,
	})
}

func (me *FailureRateAuto) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"absolute_increase":        &me.AbsoluteIncrease,
		"over_alerting_protection": &me.OverAlertingProtection,
		"relative_increase":        &me.RelativeIncrease,
	})
}
