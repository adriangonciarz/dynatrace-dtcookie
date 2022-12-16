package infrastructurehosts

import "github.com/dtcookie/hcl"

type HighGcActivityDetectionThresholds struct {
	EventThresholds        *EventThresholds `json:"eventThresholds"`
	GcSuspensionPercentage int              `json:"gcSuspensionPercentage"` // Alert if the GC suspension is higher than this threshold
	GcTimePercentage       int              `json:"gcTimePercentage"`       // Alert if GC time is higher than this threshold
}

func (me *HighGcActivityDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"event_thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EventThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"gc_suspension_percentage": {
			Type:        hcl.TypeInt,
			Description: "Alert if the GC suspension is higher than this threshold",
			Required:    true,
		},
		"gc_time_percentage": {
			Type:        hcl.TypeInt,
			Description: "Alert if GC time is higher than this threshold",
			Required:    true,
		},
	}
}

func (me *HighGcActivityDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"event_thresholds":         me.EventThresholds,
		"gc_suspension_percentage": me.GcSuspensionPercentage,
		"gc_time_percentage":       me.GcTimePercentage,
	})
}

func (me *HighGcActivityDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"event_thresholds":         &me.EventThresholds,
		"gc_suspension_percentage": &me.GcSuspensionPercentage,
		"gc_time_percentage":       &me.GcTimePercentage,
	})
}
