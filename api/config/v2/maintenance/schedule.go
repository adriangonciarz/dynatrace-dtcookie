package maintenance

import "github.com/dtcookie/hcl"

type Schedule struct {
	Type              ScheduleType       `json:"scheduleType"`
	OnceRecurrence    *OnceRecurrence    `json:"onceRecurrence,omitempty"`
	DailyRecurrence   *DailyRecurrence   `json:"dailyRecurrence,omitempty"`
	WeeklyRecurrence  *WeeklyRecurrence  `json:"weeklyRecurrence,omitempty"`
	MonthlyRecurrence *MonthlyRecurrence `json:"monthlyRecurrence,omitempty"`
}

func (me *Schedule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The time window of the maintenance window",
		},
		"once_recurrence": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The configuration for maintenance windows occuring once",
			Elem: &hcl.Resource{
				Schema: new(OnceRecurrence).Schema(),
			},
		},
		"daily_recurrence": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The configuration for maintenance windows occuring daily",
			Elem: &hcl.Resource{
				Schema: new(DailyRecurrence).Schema(),
			},
		},
		"weekly_recurrence": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The configuration for maintenance windows occuring weekly",
			Elem: &hcl.Resource{
				Schema: new(WeeklyRecurrence).Schema(),
			},
		},
		"monthly_recurrence": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The configuration for maintenance windows occuring monthly",
			Elem: &hcl.Resource{
				Schema: new(MonthlyRecurrence).Schema(),
			},
		},
	}
}

func (me *Schedule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"type":               me.Type,
		"once_recurrence":    me.OnceRecurrence,
		"daily_recurrence":   me.DailyRecurrence,
		"weekly_recurrence":  me.WeeklyRecurrence,
		"monthly_recurrence": me.MonthlyRecurrence,
	})
}

func (me *Schedule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"type":               &me.Type,
		"once_recurrence":    &me.OnceRecurrence,
		"daily_recurrence":   &me.DailyRecurrence,
		"weekly_recurrence":  &me.WeeklyRecurrence,
		"monthly_recurrence": &me.MonthlyRecurrence,
	})
}

// ScheduleType The type of the maintenance: planned or unplanned.
type ScheduleType string

// ScheduleTypes offers the known enum values
var ScheduleTypes = struct {
	Once    ScheduleType
	Daily   ScheduleType
	Weekly  ScheduleType
	Monthly ScheduleType
}{
	"ONCE",
	"DAILY",
	"WEEKLY",
	"MONTHLY",
}
