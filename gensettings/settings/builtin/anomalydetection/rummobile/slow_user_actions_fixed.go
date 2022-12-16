package rummobile

import "github.com/dtcookie/hcl"

type SlowUserActionsFixed struct {
	DurationAvoidOveralerting *SlowUserActionsAvoidOveralerting `json:"durationAvoidOveralerting"` // To avoid over-alerting do not alert for low traffic applications with less than
	DurationThresholdAllFixed *SlowUserActionsManualAll         `json:"durationThresholdAllFixed"` // Alert if the action duration of all user actions degrades beyond the absolute threshold:
	DurationThresholdSlowest  *SlowUserActionsManualSlowest     `json:"durationThresholdSlowest"`  // Alert if the action duration of the slowest 10% of user actions degrades beyond the absolute threshold:
	Sensitivity               Sensitivity                       `json:"sensitivity"`               // Detection sensitivity
}

func (me *SlowUserActionsFixed) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"duration_avoid_overalerting": {
			Type:        hcl.TypeList,
			Description: "To avoid over-alerting do not alert for low traffic applications with less than",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SlowUserActionsAvoidOveralerting).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"duration_threshold_all_fixed": {
			Type:        hcl.TypeList,
			Description: "Alert if the action duration of all user actions degrades beyond the absolute threshold:",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SlowUserActionsManualAll).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"duration_threshold_slowest": {
			Type:        hcl.TypeList,
			Description: "Alert if the action duration of the slowest 10% of user actions degrades beyond the absolute threshold:",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SlowUserActionsManualSlowest).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"sensitivity": {
			Type:        hcl.TypeString,
			Description: "Detection sensitivity",
			Required:    true,
		},
	}
}

func (me *SlowUserActionsFixed) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"duration_avoid_overalerting":  me.DurationAvoidOveralerting,
		"duration_threshold_all_fixed": me.DurationThresholdAllFixed,
		"duration_threshold_slowest":   me.DurationThresholdSlowest,
		"sensitivity":                  me.Sensitivity,
	})
}

func (me *SlowUserActionsFixed) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"duration_avoid_overalerting":  &me.DurationAvoidOveralerting,
		"duration_threshold_all_fixed": &me.DurationThresholdAllFixed,
		"duration_threshold_slowest":   &me.DurationThresholdSlowest,
		"sensitivity":                  &me.Sensitivity,
	})
}
