package node

import "github.com/dtcookie/hcl"

type ReadinessIssuesConfig struct {
	ObservationPeriodInMinutes int `json:"observationPeriodInMinutes"` // within the last
	SamplePeriodInMinutes      int `json:"samplePeriodInMinutes"`      // Node has not been ready for at least
}

func (me *ReadinessIssuesConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"observation_period_in_minutes": {
			Type:        hcl.TypeInt,
			Description: "within the last",
			Required:    true,
		},
		"sample_period_in_minutes": {
			Type:        hcl.TypeInt,
			Description: "Node has not been ready for at least",
			Required:    true,
		},
	}
}

func (me *ReadinessIssuesConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"observation_period_in_minutes": me.ObservationPeriodInMinutes,
		"sample_period_in_minutes":      me.SamplePeriodInMinutes,
	})
}

func (me *ReadinessIssuesConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"observation_period_in_minutes": &me.ObservationPeriodInMinutes,
		"sample_period_in_minutes":      &me.SamplePeriodInMinutes,
	})
}
