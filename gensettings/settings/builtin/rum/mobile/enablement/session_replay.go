package enablement

import "github.com/dtcookie/hcl"

type SessionReplay struct {
	Enabled bool `json:"enabled"` // Before enabling, Dynatrace checks your system against the [prerequisites for Session Replay](https://dt-url.net/t23s0ppi).
}

func (me *SessionReplay) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Before enabling, Dynatrace checks your system against the [prerequisites for Session Replay](https://dt-url.net/t23s0ppi).",
			Optional:    true,
		},
	}
}

func (me *SessionReplay) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled": me.Enabled,
	})
}

func (me *SessionReplay) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled": &me.Enabled,
	})
}
