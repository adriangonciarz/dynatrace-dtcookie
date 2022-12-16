package updatewindows

import "github.com/dtcookie/hcl"

type DailyRecurrence struct {
	Every           int              `json:"every"`           // Every **X** days:\n* `1` = every day,\n* `2` = every two days,\n* `3` = every three days,\n* etc.
	RecurrenceRange *RecurrenceRange `json:"recurrenceRange"` // Recurrence range
	UpdateTime      *UpdateTime      `json:"updateTime"`      // Update time
}

func (me *DailyRecurrence) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"every": {
			Type:        hcl.TypeInt,
			Description: "Every **X** days:\n* `1` = every day,\n* `2` = every two days,\n* `3` = every three days,\n* etc.",
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
		"update_time": {
			Type:        hcl.TypeList,
			Description: "Update time",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(UpdateTime).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *DailyRecurrence) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"every":            me.Every,
		"recurrence_range": me.RecurrenceRange,
		"update_time":      me.UpdateTime,
	})
}

func (me *DailyRecurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"every":            &me.Every,
		"recurrence_range": &me.RecurrenceRange,
		"update_time":      &me.UpdateTime,
	})
}
