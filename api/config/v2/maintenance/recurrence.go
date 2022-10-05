package maintenance

import "github.com/dtcookie/hcl"

type RecurrenceRange struct {
	StartDate string `json:"scheduleStartDate"` // The start date of the recurrence range in YYYY-MM-DD format
	EndDate   string `json:"scheduleEndDate"`   // The end date of the recurrence range in YYYY-MM-DD format
}

func (me *RecurrenceRange) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"start_date": {
			Type:        hcl.TypeString,
			Description: "The start date of the recurrence range in YYYY-MM-DD format",
			Required:    true,
		},
		"end_date": {
			Type:        hcl.TypeString,
			Description: "The end date of the recurrence range in YYYY-MM-DD format",
			Required:    true,
		},
	}
}

func (me *RecurrenceRange) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"start_date": me.StartDate,
		"end_date":   me.EndDate,
	})
}

func (me *RecurrenceRange) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"start_date": &me.StartDate,
		"end_date":   &me.EndDate,
	})
}
