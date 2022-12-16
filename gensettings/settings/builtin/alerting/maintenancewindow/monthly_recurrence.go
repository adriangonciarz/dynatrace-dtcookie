package maintenancewindow

import "github.com/dtcookie/hcl"

type MonthlyRecurrence struct {
	DayOfMonth      int              `json:"dayOfMonth"`      // If the selected day does not fall within the month, the maintenance window will be active on the last day of the month.
	RecurrenceRange *RecurrenceRange `json:"recurrenceRange"` // Recurrence range
	TimeWindow      *TimeWindow      `json:"timeWindow"`      // Time window
}

func (me *MonthlyRecurrence) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"day_of_month": {
			Type:        hcl.TypeInt,
			Description: "If the selected day does not fall within the month, the maintenance window will be active on the last day of the month.",
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

func (me *MonthlyRecurrence) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"day_of_month":     me.DayOfMonth,
		"recurrence_range": me.RecurrenceRange,
		"time_window":      me.TimeWindow,
	})
}

func (me *MonthlyRecurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"day_of_month":     &me.DayOfMonth,
		"recurrence_range": &me.RecurrenceRange,
		"time_window":      &me.TimeWindow,
	})
}
