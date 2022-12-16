package maintenancewindow

import "github.com/dtcookie/hcl"

// Maintenance window. The date time window when the maintenance will take place once.
type OnceRecurrence struct {
	EndTime   string `json:"endTime"`   // End time
	StartTime string `json:"startTime"` // Start time
	TimeZone  string `json:"timeZone"`
}

func (me *OnceRecurrence) Schema() map[string]*hcl.Schema {
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

func (me *OnceRecurrence) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"end_time":   me.EndTime,
		"start_time": me.StartTime,
		"time_zone":  me.TimeZone,
	})
}

func (me *OnceRecurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"end_time":   &me.EndTime,
		"start_time": &me.StartTime,
		"time_zone":  &me.TimeZone,
	})
}
