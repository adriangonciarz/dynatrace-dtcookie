package maintenance

import "github.com/dtcookie/hcl"

type DailyRecurrence struct {
	TimeWindow      *TimeWindow      `json:"timeWindow"`      // The time window of the maintenance window
	RecurrenceRange *RecurrenceRange `json:"recurrenceRange"` // The recurrence date range of the maintenance window
}

func (me *DailyRecurrence) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
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

func (me *DailyRecurrence) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"time_window":      me.TimeWindow,
		"recurrence_range": me.RecurrenceRange,
	})
}

func (me *DailyRecurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"time_window":      &me.TimeWindow,
		"recurrence_range": &me.RecurrenceRange,
	})
}
