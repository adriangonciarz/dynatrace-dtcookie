package updatewindows

import "github.com/dtcookie/hcl"

type MonthlyRecurrence struct {
	Every            int              `json:"every"`            // Every **X** months:\n* `1` = every month,\n* `2` = every two months,\n* `3` = every three months,\n* etc.
	RecurrenceRange  *RecurrenceRange `json:"recurrenceRange"`  // Recurrence range
	SelectedMonthDay int              `json:"selectedMonthDay"` // Day of the month
	UpdateTime       *UpdateTime      `json:"updateTime"`       // Update time
}

func (me *MonthlyRecurrence) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"every": {
			Type:        hcl.TypeInt,
			Description: "Every **X** months:\n* `1` = every month,\n* `2` = every two months,\n* `3` = every three months,\n* etc.",
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
		"selected_month_day": {
			Type:        hcl.TypeInt,
			Description: "Day of the month",
			Required:    true,
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

func (me *MonthlyRecurrence) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"every":              me.Every,
		"recurrence_range":   me.RecurrenceRange,
		"selected_month_day": me.SelectedMonthDay,
		"update_time":        me.UpdateTime,
	})
}

func (me *MonthlyRecurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"every":              &me.Every,
		"recurrence_range":   &me.RecurrenceRange,
		"selected_month_day": &me.SelectedMonthDay,
		"update_time":        &me.UpdateTime,
	})
}
