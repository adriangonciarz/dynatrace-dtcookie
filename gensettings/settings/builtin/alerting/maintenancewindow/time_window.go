package maintenancewindow

import "github.com/dtcookie/hcl"

// Time window. The time window when the maintenance will take place.
type TimeWindow struct {
	EndTime   string `json:"endTime"`   // End time
	StartTime string `json:"startTime"` // Start time
	TimeZone  string `json:"timeZone"`
}

func (me *TimeWindow) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"end_time": {
			Type:        hcl.TypeString,
			Description: "End time",
			Required:    true,
		},
		"start_time": {
			Type:        hcl.TypeString,
			Description: "Start time",
			Required:    true,
		},
		"time_zone": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *TimeWindow) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"end_time":   me.EndTime,
		"start_time": me.StartTime,
		"time_zone":  me.TimeZone,
	})
}

func (me *TimeWindow) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"end_time":   &me.EndTime,
		"start_time": &me.StartTime,
		"time_zone":  &me.TimeZone,
	})
}
