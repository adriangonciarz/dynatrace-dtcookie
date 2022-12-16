package services

import "github.com/dtcookie/hcl"

type FailureRateFixed struct {
	OverAlertingProtection *OverAlertingProtection `json:"overAlertingProtection"` // Avoid over-alerting
	Sensitivity            Sensitivity             `json:"sensitivity"`
	Threshold              float64                 `json:"threshold"`
}

func (me *FailureRateFixed) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"over_alerting_protection": {
			Type:        hcl.TypeList,
			Description: "Avoid over-alerting",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(OverAlertingProtection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"sensitivity": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"threshold": {
			Type:        hcl.TypeFloat,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *FailureRateFixed) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"over_alerting_protection": me.OverAlertingProtection,
		"sensitivity":              me.Sensitivity,
		"threshold":                me.Threshold,
	})
}

func (me *FailureRateFixed) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"over_alerting_protection": &me.OverAlertingProtection,
		"sensitivity":              &me.Sensitivity,
		"threshold":                &me.Threshold,
	})
}
