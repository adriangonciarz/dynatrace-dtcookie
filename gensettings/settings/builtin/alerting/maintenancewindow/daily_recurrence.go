package maintenancewindow

import "github.com/dtcookie/hcl"

type DailyRecurrence struct {
	RecurrenceRange *RecurrenceRange `json:"recurrenceRange"` // Recurrence range
	TimeWindow      *TimeWindow      `json:"timeWindow"`      // Time window
}

func (me *DailyRecurrence) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"recurrence_range": {
			Type:        hcl.TypeList,
			Description: "Recurrence range",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(RecurrenceRange).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"time_window": {
			Type:        hcl.TypeList,
			Description: "Time window",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(TimeWindow).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *DailyRecurrence) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"recurrence_range": me.RecurrenceRange,
		"time_window":      me.TimeWindow,
	})
}

func (me *DailyRecurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"recurrence_range": &me.RecurrenceRange,
		"time_window":      &me.TimeWindow,
	})
}
