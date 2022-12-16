package updatewindows

import "github.com/dtcookie/hcl"

type WeeklyRecurrence struct {
	Every            int               `json:"every"`            // Every **X** weeks:\n* `1` = every week,\n* `2` = every two weeks,\n* `3` = every three weeks,\n* etc.
	RecurrenceRange  *RecurrenceRange  `json:"recurrenceRange"`  // Recurrence range
	SelectedWeekDays *SelectedWeekDays `json:"selectedWeekDays"` // Day of the week
	UpdateTime       *UpdateTime       `json:"updateTime"`       // Update time
}

func (me *WeeklyRecurrence) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"every": {
			Type:        hcl.TypeInt,
			Description: "Every **X** weeks:\n* `1` = every week,\n* `2` = every two weeks,\n* `3` = every three weeks,\n* etc.",
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
		"selected_week_days": {
			Type:        hcl.TypeList,
			Description: "Day of the week",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SelectedWeekDays).Schema()},
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

func (me *WeeklyRecurrence) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"every":              me.Every,
		"recurrence_range":   me.RecurrenceRange,
		"selected_week_days": me.SelectedWeekDays,
		"update_time":        me.UpdateTime,
	})
}

func (me *WeeklyRecurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"every":              &me.Every,
		"recurrence_range":   &me.RecurrenceRange,
		"selected_week_days": &me.SelectedWeekDays,
		"update_time":        &me.UpdateTime,
	})
}
