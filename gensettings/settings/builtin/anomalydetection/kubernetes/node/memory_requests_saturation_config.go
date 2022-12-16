package node

import "github.com/dtcookie/hcl"

type MemoryRequestsSaturationConfig struct {
	ObservationPeriodInMinutes int `json:"observationPeriodInMinutes"` // within the last
	SamplePeriodInMinutes      int `json:"samplePeriodInMinutes"`      // of total node memory capacities for at least
	Threshold                  int `json:"threshold"`                  // Requested memory capacities have been higher than
}

func (me *MemoryRequestsSaturationConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"observation_period_in_minutes": {
			Type:        hcl.TypeInt,
			Description: "within the last",
			Required:    true,
		},
		"sample_period_in_minutes": {
			Type:        hcl.TypeInt,
			Description: "of total node memory capacities for at least",
			Required:    true,
		},
		"threshold": {
			Type:        hcl.TypeInt,
			Description: "Requested memory capacities have been higher than",
			Required:    true,
		},
	}
}

func (me *MemoryRequestsSaturationConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"observation_period_in_minutes": me.ObservationPeriodInMinutes,
		"sample_period_in_minutes":      me.SamplePeriodInMinutes,
		"threshold":                     me.Threshold,
	})
}

func (me *MemoryRequestsSaturationConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"observation_period_in_minutes": &me.ObservationPeriodInMinutes,
		"sample_period_in_minutes":      &me.SamplePeriodInMinutes,
		"threshold":                     &me.Threshold,
	})
}
