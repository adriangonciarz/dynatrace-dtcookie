package rumweb

import "github.com/dtcookie/hcl"

type AppTrafficSpikes struct {
	Enabled       bool           `json:"enabled"`       // Detect traffic spikes
	TrafficSpikes *TrafficSpikes `json:"trafficSpikes"` // Dynatrace learns your typical application traffic over an observation period of one week.\n\nDepending on this expected value Dynatrace detects abnormal traffic spikes within your application.
}

func (me *AppTrafficSpikes) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect traffic spikes",
			Optional:    true,
		},
		"traffic_spikes": {
			Type:        hcl.TypeList,
			Description: "Dynatrace learns your typical application traffic over an observation period of one week.\n\nDepending on this expected value Dynatrace detects abnormal traffic spikes within your application.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(TrafficSpikes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *AppTrafficSpikes) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":        me.Enabled,
		"traffic_spikes": me.TrafficSpikes,
	})
}

func (me *AppTrafficSpikes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":        &me.Enabled,
		"traffic_spikes": &me.TrafficSpikes,
	})
}
