package maintenance

import "github.com/dtcookie/hcl"

type MonthlyRecurrence struct {
	DayOfMonth      *int32           `json:"dayOfMonth"`      // The day of the month for monthly maintenance.  The value of `31` is treated as the last day of the month for months that don't have a 31st day. The value of `30` is also treated as the last day of the month for February.	TimeWindow      TimeWindow      `json:"timeWindow"`
	TimeWindow      *TimeWindow      `json:"timeWindow"`      // The time window of the maintenance window
	RecurrenceRange *RecurrenceRange `json:"recurrenceRange"` // The recurrence date range of the maintenance window
}

func (me *MonthlyRecurrence) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"day_of_month": {
			Type:        hcl.TypeInt,
			Description: "The day of the month for monthly maintenance.  The value of `31` is treated as the last day of the month for months that don't have a 31st day. The value of `30` is also treated as the last day of the month for February",
			Required:    true,
		},
		"time_window": {
			Type:        hcl.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "The time window of the maintenance window",
			Elem: &hcl.Resource{
				Schema: new(TimeWindow).Schema(),
			},
		},
		"recurrence_range": {
			Type:        hcl.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "The recurrence date range of the maintenance window",
			Elem: &hcl.Resource{
				Schema: new(RecurrenceRange).Schema(),
			},
		},
	}
}

func (me *MonthlyRecurrence) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"day_of_month":     me.DayOfMonth,
		"time_window":      me.TimeWindow,
		"recurrence_range": me.RecurrenceRange,
	})
}

func (me *MonthlyRecurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"day_of_month":     &me.DayOfMonth,
		"time_window":      &me.TimeWindow,
		"recurrence_range": &me.RecurrenceRange,
	})
}
