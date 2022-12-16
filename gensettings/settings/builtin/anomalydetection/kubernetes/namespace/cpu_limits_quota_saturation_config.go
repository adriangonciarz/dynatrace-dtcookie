package namespace

import "github.com/dtcookie/hcl"

type CpuLimitsQuotaSaturationConfig struct {
	ObservationPeriodInMinutes int `json:"observationPeriodInMinutes"` // within the last
	SamplePeriodInMinutes      int `json:"samplePeriodInMinutes"`      // for at least
	Threshold                  int `json:"threshold"`                  // Utilized CPU limits quota has been higher than
}

func (me *CpuLimitsQuotaSaturationConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"observation_period_in_minutes": {
			Type:        hcl.TypeInt,
			Description: "within the last",
			Required:    true,
		},
		"sample_period_in_minutes": {
			Type:        hcl.TypeInt,
			Description: "for at least",
			Required:    true,
		},
		"threshold": {
			Type:        hcl.TypeInt,
			Description: "Utilized CPU limits quota has been higher than",
			Required:    true,
		},
	}
}

func (me *CpuLimitsQuotaSaturationConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"observation_period_in_minutes": me.ObservationPeriodInMinutes,
		"sample_period_in_minutes":      me.SamplePeriodInMinutes,
		"threshold":                     me.Threshold,
	})
}

func (me *CpuLimitsQuotaSaturationConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"observation_period_in_minutes": &me.ObservationPeriodInMinutes,
		"sample_period_in_minutes":      &me.SamplePeriodInMinutes,
		"threshold":                     &me.Threshold,
	})
}
