package maintenance

import "github.com/dtcookie/hcl"

type WeeklyRecurrence struct {
	DayOfWeek       DayOfWeek        `json:"dayOfWeek"`       // The day of the week for weekly maintenance.  The format is the full name of the day in upper case, for example `THURSDAY`.
	TimeWindow      *TimeWindow      `json:"timeWindow"`      // The time window of the maintenance window
	RecurrenceRange *RecurrenceRange `json:"recurrenceRange"` // The recurrence date range of the maintenance window
}

func (me *WeeklyRecurrence) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"day_of_week": {
			Type:        hcl.TypeString,
			Description: "The day of the week for weekly maintenance.  The format is the full name of the day in upper case, for example `THURSDAY`",
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

func (me *WeeklyRecurrence) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"day_of_week":      me.DayOfWeek,
		"time_window":      me.TimeWindow,
		"recurrence_range": me.RecurrenceRange,
	})
}

func (me *WeeklyRecurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"day_of_week":      &me.DayOfWeek,
		"time_window":      &me.TimeWindow,
		"recurrence_range": &me.RecurrenceRange,
	})
}

// DayOfWeek The day of the week for weekly maintenance.
// The format is the full name of the day in upper case, for example `THURSDAY`.
type DayOfWeek string

// DayOfWeeks offers the known enum values
var DayOfWeeks = struct {
	Friday    DayOfWeek
	Monday    DayOfWeek
	Saturday  DayOfWeek
	Sunday    DayOfWeek
	Thursday  DayOfWeek
	Tuesday   DayOfWeek
	Wednesday DayOfWeek
}{
	"FRIDAY",
	"MONDAY",
	"SATURDAY",
	"SUNDAY",
	"THURSDAY",
	"TUESDAY",
	"WEDNESDAY",
}
