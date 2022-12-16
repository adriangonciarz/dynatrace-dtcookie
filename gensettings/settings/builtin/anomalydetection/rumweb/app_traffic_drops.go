package rumweb

import "github.com/dtcookie/hcl"

type AppTrafficDrops struct {
	Enabled      bool          `json:"enabled"`      // Detect traffic drops
	TrafficDrops *TrafficDrops `json:"trafficDrops"` // Dynatrace learns your typical application traffic over an observation period of one week.\n\nDepending on this expected value Dynatrace detects abnormal traffic drops within your application.
}

func (me *AppTrafficDrops) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect traffic drops",
			Optional:    true,
		},
		"traffic_drops": {
			Type:        hcl.TypeList,
			Description: "Dynatrace learns your typical application traffic over an observation period of one week.\n\nDepending on this expected value Dynatrace detects abnormal traffic drops within your application.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(TrafficDrops).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *AppTrafficDrops) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":       me.Enabled,
		"traffic_drops": me.TrafficDrops,
	})
}

func (me *AppTrafficDrops) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":       &me.Enabled,
		"traffic_drops": &me.TrafficDrops,
	})
}
