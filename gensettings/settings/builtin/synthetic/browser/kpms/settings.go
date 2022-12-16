package kpms

import "github.com/dtcookie/hcl"

type Settings struct {
	LoadActionKpm LoadKpm `json:"loadActionKpm"` // Load action key performance metric
	XhrActionKpm  XhrKpm  `json:"xhrActionKpm"`  // XHR action key performance metric
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"load_action_kpm": {
			Type:        hcl.TypeString,
			Description: "Load action key performance metric",
			Required:    true,
		},
		"xhr_action_kpm": {
			Type:        hcl.TypeString,
			Description: "XHR action key performance metric",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"load_action_kpm": me.LoadActionKpm,
		"xhr_action_kpm":  me.XhrActionKpm,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"load_action_kpm": &me.LoadActionKpm,
		"xhr_action_kpm":  &me.XhrActionKpm,
	})
}
