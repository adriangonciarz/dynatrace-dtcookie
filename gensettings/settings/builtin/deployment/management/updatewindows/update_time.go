package updatewindows

import "github.com/dtcookie/hcl"

type UpdateTime struct {
	Duration  int          `json:"duration"`  // Duration (minutes)
	StartTime string       `json:"startTime"` // Start time (24-hour clock)
	TimeZone  TimezoneEnum `json:"timeZone"`  // Time zone
}

func (me *UpdateTime) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"duration": {
			Type:        hcl.TypeInt,
			Description: "Duration (minutes)",
			Required:    true,
		},
		"start_time": {
			Type:        hcl.TypeString,
			Description: "Start time (24-hour clock)",
			Required:    true,
		},
		"time_zone": {
			Type:        hcl.TypeString,
			Description: "Time zone",
			Required:    true,
		},
	}
}

func (me *UpdateTime) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"duration":   me.Duration,
		"start_time": me.StartTime,
		"time_zone":  me.TimeZone,
	})
}

func (me *UpdateTime) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"duration":   &me.Duration,
		"start_time": &me.StartTime,
		"time_zone":  &me.TimeZone,
	})
}
