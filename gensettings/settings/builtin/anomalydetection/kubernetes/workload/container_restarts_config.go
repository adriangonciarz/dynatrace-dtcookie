package workload

import "github.com/dtcookie/hcl"

type ContainerRestartsConfig struct {
	ObservationPeriodInMinutes int `json:"observationPeriodInMinutes"` // within the last
	SamplePeriodInMinutes      int `json:"samplePeriodInMinutes"`      // each minute for at least
	Threshold                  int `json:"threshold"`                  // there were at least
}

func (me *ContainerRestartsConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"observation_period_in_minutes": {
			Type:        hcl.TypeInt,
			Description: "within the last",
			Required:    true,
		},
		"sample_period_in_minutes": {
			Type:        hcl.TypeInt,
			Description: "each minute for at least",
			Required:    true,
		},
		"threshold": {
			Type:        hcl.TypeInt,
			Description: "there were at least",
			Required:    true,
		},
	}
}

func (me *ContainerRestartsConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"observation_period_in_minutes": me.ObservationPeriodInMinutes,
		"sample_period_in_minutes":      me.SamplePeriodInMinutes,
		"threshold":                     me.Threshold,
	})
}

func (me *ContainerRestartsConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"observation_period_in_minutes": &me.ObservationPeriodInMinutes,
		"sample_period_in_minutes":      &me.SamplePeriodInMinutes,
		"threshold":                     &me.Threshold,
	})
}
