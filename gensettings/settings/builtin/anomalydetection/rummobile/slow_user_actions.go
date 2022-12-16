package rummobile

import "github.com/dtcookie/hcl"

type SlowUserActions struct {
	DetectionMode        DetectionMode         `json:"detectionMode"` // Detection strategy for slow user actions
	Enabled              bool                  `json:"enabled"`       // Detect slow user actions
	SlowUserActionsAuto  *SlowUserActionsAuto  `json:"slowUserActionsAuto"`
	SlowUserActionsFixed *SlowUserActionsFixed `json:"slowUserActionsFixed"`
}

func (me *SlowUserActions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection strategy for slow user actions",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect slow user actions",
			Optional:    true,
		},
		"slow_user_actions_auto": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SlowUserActionsAuto).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"slow_user_actions_fixed": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SlowUserActionsFixed).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *SlowUserActions) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"detection_mode":          me.DetectionMode,
		"enabled":                 me.Enabled,
		"slow_user_actions_auto":  me.SlowUserActionsAuto,
		"slow_user_actions_fixed": me.SlowUserActionsFixed,
	})
}

func (me *SlowUserActions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"detection_mode":          &me.DetectionMode,
		"enabled":                 &me.Enabled,
		"slow_user_actions_auto":  &me.SlowUserActionsAuto,
		"slow_user_actions_fixed": &me.SlowUserActionsFixed,
	})
}
