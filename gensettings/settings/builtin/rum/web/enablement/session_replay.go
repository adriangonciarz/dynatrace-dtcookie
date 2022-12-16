package enablement

import "github.com/dtcookie/hcl"

type SessionReplay struct {
	CostAndTrafficControl int  `json:"costAndTrafficControl"` // [Percentage of user sessions recorded with Session Replay](https://dt-url.net/sr-cost-traffic-control). For example, if you have 50% for RUM and 50% for Session Replay, it results in 25% of sessions recorded with Session Replay.
	Enabled               bool `json:"enabled"`               // Before enabling, Dynatrace checks your system against the prerequisites for [Session Replay](https://dt-url.net/ma3m0psf).
}

func (me *SessionReplay) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cost_and_traffic_control": {
			Type:        hcl.TypeInt,
			Description: "[Percentage of user sessions recorded with Session Replay](https://dt-url.net/sr-cost-traffic-control). For example, if you have 50% for RUM and 50% for Session Replay, it results in 25% of sessions recorded with Session Replay.",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Before enabling, Dynatrace checks your system against the prerequisites for [Session Replay](https://dt-url.net/ma3m0psf).",
			Optional:    true,
		},
	}
}

func (me *SessionReplay) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"cost_and_traffic_control": me.CostAndTrafficControl,
		"enabled":                  me.Enabled,
	})
}

func (me *SessionReplay) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cost_and_traffic_control": &me.CostAndTrafficControl,
		"enabled":                  &me.Enabled,
	})
}
