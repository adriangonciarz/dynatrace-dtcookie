package maintenancewindow

import "github.com/dtcookie/hcl"

// Recurrence range. The date range in which maintenance is activated during the specified time window.
type RecurrenceRange struct {
	ScheduleEndDate   string `json:"scheduleEndDate"`   // End date
	ScheduleStartDate string `json:"scheduleStartDate"` // Start date
}

func (me *RecurrenceRange) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"schedule_end_date": {
			Type:        hcl.TypeString,
			Description: "End date",
			Required:    true,
		},
		"schedule_start_date": {
			Type:        hcl.TypeString,
			Description: "Start date",
			Required:    true,
		},
	}
}

func (me *RecurrenceRange) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"schedule_end_date":   me.ScheduleEndDate,
		"schedule_start_date": me.ScheduleStartDate,
	})
}

func (me *RecurrenceRange) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"schedule_end_date":   &me.ScheduleEndDate,
		"schedule_start_date": &me.ScheduleStartDate,
	})
}
