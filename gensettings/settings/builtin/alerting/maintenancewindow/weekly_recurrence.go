package maintenancewindow

import "github.com/dtcookie/hcl"

type WeeklyRecurrence struct {
	DayOfWeek       DayOfWeekType    `json:"dayOfWeek"`       // Day of week
	RecurrenceRange *RecurrenceRange `json:"recurrenceRange"` // Recurrence range
	TimeWindow      *TimeWindow      `json:"timeWindow"`      // Time window
}

func (me *WeeklyRecurrence) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"day_of_week": {
			Type:        hcl.TypeString,
			Description: "Day of week",
			Required:    true,
		},
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

func (me *WeeklyRecurrence) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"day_of_week":      me.DayOfWeek,
		"recurrence_range": me.RecurrenceRange,
		"time_window":      me.TimeWindow,
	})
}

func (me *WeeklyRecurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"day_of_week":      &me.DayOfWeek,
		"recurrence_range": &me.RecurrenceRange,
		"time_window":      &me.TimeWindow,
	})
}
