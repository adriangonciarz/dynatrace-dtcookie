package rumweb

import "github.com/dtcookie/hcl"

type TrafficSpikes struct {
	MinutesAbnormalState   float64 `json:"minutesAbnormalState"`   // Minutes an application has to stay in abnormal state before alert
	TrafficSpikePercentage float64 `json:"trafficSpikePercentage"` // Alert if the observed traffic is more than this percentage of the expected value
}

func (me *TrafficSpikes) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"minutes_abnormal_state": {
			Type:        hcl.TypeFloat,
			Description: "Minutes an application has to stay in abnormal state before alert",
			Required:    true,
		},
		"traffic_spike_percentage": {
			Type:        hcl.TypeFloat,
			Description: "Alert if the observed traffic is more than this percentage of the expected value",
			Required:    true,
		},
	}
}

func (me *TrafficSpikes) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"minutes_abnormal_state":   me.MinutesAbnormalState,
		"traffic_spike_percentage": me.TrafficSpikePercentage,
	})
}

func (me *TrafficSpikes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"minutes_abnormal_state":   &me.MinutesAbnormalState,
		"traffic_spike_percentage": &me.TrafficSpikePercentage,
	})
}
