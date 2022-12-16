package updatewindows

import "github.com/dtcookie/hcl"

type Settings struct {
	DailyRecurrence   *DailyRecurrence   `json:"dailyRecurrence"`
	Enabled           bool               `json:"enabled"` // On/Off
	MonthlyRecurrence *MonthlyRecurrence `json:"monthlyRecurrence"`
	Name              string             `json:"name"` // Name
	OnceRecurrence    *OnceRecurrence    `json:"onceRecurrence"`
	Recurrence        RecurrenceEnum     `json:"recurrence"` // Recurrence
	WeeklyRecurrence  *WeeklyRecurrence  `json:"weeklyRecurrence"`
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"daily_recurrence": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DailyRecurrence).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "On/Off",
			Optional:    true,
		},
		"monthly_recurrence": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MonthlyRecurrence).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "Name",
			Required:    true,
		},
		"once_recurrence": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(OnceRecurrence).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"recurrence": {
			Type:        hcl.TypeString,
			Description: "Recurrence",
			Required:    true,
		},
		"weekly_recurrence": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(WeeklyRecurrence).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"daily_recurrence":   me.DailyRecurrence,
		"enabled":            me.Enabled,
		"monthly_recurrence": me.MonthlyRecurrence,
		"name":               me.Name,
		"once_recurrence":    me.OnceRecurrence,
		"recurrence":         me.Recurrence,
		"weekly_recurrence":  me.WeeklyRecurrence,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"daily_recurrence":   &me.DailyRecurrence,
		"enabled":            &me.Enabled,
		"monthly_recurrence": &me.MonthlyRecurrence,
		"name":               &me.Name,
		"once_recurrence":    &me.OnceRecurrence,
		"recurrence":         &me.Recurrence,
		"weekly_recurrence":  &me.WeeklyRecurrence,
	})
}
