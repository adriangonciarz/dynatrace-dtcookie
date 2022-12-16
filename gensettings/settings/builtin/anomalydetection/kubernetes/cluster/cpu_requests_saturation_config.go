package cluster

import "github.com/dtcookie/hcl"

type CpuRequestsSaturationConfig struct {
	ObservationPeriodInMinutes int `json:"observationPeriodInMinutes"` // within the last
	SamplePeriodInMinutes      int `json:"samplePeriodInMinutes"`      // of total cluster CPU capacities for at least
	Threshold                  int `json:"threshold"`                  // Requested CPU capacities have been higher than
}

func (me *CpuRequestsSaturationConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"observation_period_in_minutes": {
			Type:        hcl.TypeInt,
			Description: "within the last",
			Required:    true,
		},
		"sample_period_in_minutes": {
			Type:        hcl.TypeInt,
			Description: "of total cluster CPU capacities for at least",
			Required:    true,
		},
		"threshold": {
			Type:        hcl.TypeInt,
			Description: "Requested CPU capacities have been higher than",
			Required:    true,
		},
	}
}

func (me *CpuRequestsSaturationConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"observation_period_in_minutes": me.ObservationPeriodInMinutes,
		"sample_period_in_minutes":      me.SamplePeriodInMinutes,
		"threshold":                     me.Threshold,
	})
}

func (me *CpuRequestsSaturationConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"observation_period_in_minutes": &me.ObservationPeriodInMinutes,
		"sample_period_in_minutes":      &me.SamplePeriodInMinutes,
		"threshold":                     &me.Threshold,
	})
}
