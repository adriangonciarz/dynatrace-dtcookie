package databases

import "github.com/dtcookie/hcl"

type DatabaseConnections struct {
	Enabled           bool `json:"enabled"`           // Detect failed database connects
	MaxFailedConnects int  `json:"maxFailedConnects"` // Threshold
	TimePeriod        int  `json:"timePeriod"`        // Time span
}

func (me *DatabaseConnections) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect failed database connects",
			Optional:    true,
		},
		"max_failed_connects": {
			Type:        hcl.TypeInt,
			Description: "Threshold",
			Required:    true,
		},
		"time_period": {
			Type:        hcl.TypeInt,
			Description: "Time span",
			Required:    true,
		},
	}
}

func (me *DatabaseConnections) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":             me.Enabled,
		"max_failed_connects": me.MaxFailedConnects,
		"time_period":         me.TimePeriod,
	})
}

func (me *DatabaseConnections) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":             &me.Enabled,
		"max_failed_connects": &me.MaxFailedConnects,
		"time_period":         &me.TimePeriod,
	})
}
