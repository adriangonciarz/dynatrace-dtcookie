package maintenancewindow

import "github.com/dtcookie/hcl"

type Schedule struct {
	DailyRecurrence   *DailyRecurrence   `json:"dailyRecurrence"`
	MonthlyRecurrence *MonthlyRecurrence `json:"monthlyRecurrence"`
	OnceRecurrence    *OnceRecurrence    `json:"onceRecurrence"`
	ScheduleType      ScheduleType       `json:"scheduleType"` // Defines the recurrence type of the maintenance window.. * **Once**: One time maintenance window with start and end date time.\n* **Daily**: Maintenance window occurs every day during the configured time window.\n* **Weekly**: Maintenance window occurs each week on one day during the configured time window.\n* **Monthly**: Maintenance window occurs each month on one day during the configured time window.
	WeeklyRecurrence  *WeeklyRecurrence  `json:"weeklyRecurrence"`
}

func (me *Schedule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"daily_recurrence": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DailyRecurrence).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"monthly_recurrence": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MonthlyRecurrence).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"once_recurrence": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(OnceRecurrence).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"schedule_type": {
			Type:        hcl.TypeString,
			Description: "Defines the recurrence type of the maintenance window.. * **Once**: One time maintenance window with start and end date time.\n* **Daily**: Maintenance window occurs every day during the configured time window.\n* **Weekly**: Maintenance window occurs each week on one day during the configured time window.\n* **Monthly**: Maintenance window occurs each month on one day during the configured time window.",
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

func (me *Schedule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"daily_recurrence":   me.DailyRecurrence,
		"monthly_recurrence": me.MonthlyRecurrence,
		"once_recurrence":    me.OnceRecurrence,
		"schedule_type":      me.ScheduleType,
		"weekly_recurrence":  me.WeeklyRecurrence,
	})
}

func (me *Schedule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"daily_recurrence":   &me.DailyRecurrence,
		"monthly_recurrence": &me.MonthlyRecurrence,
		"once_recurrence":    &me.OnceRecurrence,
		"schedule_type":      &me.ScheduleType,
		"weekly_recurrence":  &me.WeeklyRecurrence,
	})
}
