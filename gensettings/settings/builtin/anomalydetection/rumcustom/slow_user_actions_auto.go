package rumcustom

import "github.com/dtcookie/hcl"

type SlowUserActionsAuto struct {
	DurationAvoidOveralerting *SlowUserActionsAvoidOveralerting `json:"durationAvoidOveralerting"` // To avoid over-alerting do not alert for low traffic applications with less than
	DurationThresholdAll      *SlowUserActionsAutoAll           `json:"durationThresholdAll"`      // Alert if the action duration of all user actions degrades beyond **both** the absolute and relative threshold:
	DurationThresholdSlowest  *SlowUserActionsAutoSlowest       `json:"durationThresholdSlowest"`  // Alert if the action duration of the slowest 10% of user actions degrades beyond **both** the absolute and relative threshold:
}

func (me *SlowUserActionsAuto) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"duration_avoid_overalerting": {
			Type:        hcl.TypeList,
			Description: "To avoid over-alerting do not alert for low traffic applications with less than",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SlowUserActionsAvoidOveralerting).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"duration_threshold_all": {
			Type:        hcl.TypeList,
			Description: "Alert if the action duration of all user actions degrades beyond **both** the absolute and relative threshold:",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SlowUserActionsAutoAll).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"duration_threshold_slowest": {
			Type:        hcl.TypeList,
			Description: "Alert if the action duration of the slowest 10% of user actions degrades beyond **both** the absolute and relative threshold:",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SlowUserActionsAutoSlowest).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *SlowUserActionsAuto) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"duration_avoid_overalerting": me.DurationAvoidOveralerting,
		"duration_threshold_all":      me.DurationThresholdAll,
		"duration_threshold_slowest":  me.DurationThresholdSlowest,
	})
}

func (me *SlowUserActionsAuto) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"duration_avoid_overalerting": &me.DurationAvoidOveralerting,
		"duration_threshold_all":      &me.DurationThresholdAll,
		"duration_threshold_slowest":  &me.DurationThresholdSlowest,
	})
}
