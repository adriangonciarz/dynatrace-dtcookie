package maintenance

import "github.com/dtcookie/hcl"

type OnceRecurrence struct {
	StartTime string `json:"startTime"` // The start time of the maintenance window validity period in YYYY-MM-DDThh:mm:ss format (for example, `2022-01-01T08:00:00`)
	EndTime   string `json:"endTime"`   // The end time of the maintenance window validity period in YYYY-MM-DDThh:mm:ss format (for example, `2022-01-01T08:00:00`)
	TimeZone  string `json:"timeZone"`  // The time zone of the start and end time. Default time zone is UTC. You can use either UTC offset `UTC+01:00` format or the IANA Time Zone Database format (for example, `Europe/Vienna`)
}

func (me *OnceRecurrence) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"start_time": {
			Type:        hcl.TypeString,
			Description: "The start time of the maintenance window validity period in YYYY-MM-DDThh:mm:ss format (for example, `2022-01-01T08:00:00`)",
			Required:    true,
		},
		"end_time": {
			Type:        hcl.TypeString,
			Description: "The end time of the maintenance window validity period in YYYY-MM-DDThh:mm:ss format (for example, `2022-01-01T08:00:00`)",
			Required:    true,
		},
		"time_zone": {
			Type:        hcl.TypeString,
			Description: "The time zone of the start and end time. Default time zone is UTC. You can use either UTC offset `UTC+01:00` format or the IANA Time Zone Database format (for example, `Europe/Vienna`)",
			Required:    true,
		},
	}
}

func (me *OnceRecurrence) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"start_time": me.StartTime,
		"end_time":   me.EndTime,
		"time_zone":  me.TimeZone,
	})
}

func (me *OnceRecurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"start_time": &me.StartTime,
		"end_time":   &me.EndTime,
		"time_zone":  &me.TimeZone,
	})
}
