package maintenance

import "github.com/dtcookie/hcl"

type TimeWindow struct {
	StartTime string `json:"startTime"` // The start time of the maintenance window validity period in hh:mm:ss format
	EndTime   string `json:"endTime"`   // The end time of the maintenance window validity period in hh:mm:ss format
	TimeZone  string `json:"timeZone"`  // The time zone of the start and end time. Default time zone is UTC. You can use either UTC offset `UTC+01:00` format or the IANA Time Zone Database format (for example, `Europe/Vienna`)
}

func (me *TimeWindow) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"start_time": {
			Type:        hcl.TypeString,
			Description: "The start time of the maintenance window validity period in hh:mm:ss format",
			Required:    true,
		},
		"end_time": {
			Type:        hcl.TypeString,
			Description: "The end time of the maintenance window validity period in hh:mm:ss format",
			Required:    true,
		},
		"time_zone": {
			Type:        hcl.TypeString,
			Description: "The time zone of the start and end time. Default time zone is UTC. You can use either UTC offset `UTC+01:00` format or the IANA Time Zone Database format (for example, `Europe/Vienna`)",
			Required:    true,
		},
	}
}

func (me *TimeWindow) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"start_time": me.StartTime,
		"end_time":   me.EndTime,
		"time_zone":  me.TimeZone,
	})
}

func (me *TimeWindow) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"start_time": &me.StartTime,
		"end_time":   &me.EndTime,
		"time_zone":  &me.TimeZone,
	})
}
