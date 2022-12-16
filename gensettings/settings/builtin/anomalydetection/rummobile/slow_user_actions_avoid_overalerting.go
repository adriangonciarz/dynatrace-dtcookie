package rummobile

import "github.com/dtcookie/hcl"

type SlowUserActionsAvoidOveralerting struct {
	MinActionRate int `json:"minActionRate"`
}

func (me *SlowUserActionsAvoidOveralerting) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"min_action_rate": {
			Type:        hcl.TypeInt,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *SlowUserActionsAvoidOveralerting) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"min_action_rate": me.MinActionRate,
	})
}

func (me *SlowUserActionsAvoidOveralerting) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"min_action_rate": &me.MinActionRate,
	})
}
