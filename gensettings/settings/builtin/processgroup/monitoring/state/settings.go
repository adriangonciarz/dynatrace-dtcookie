package state

import "github.com/dtcookie/hcl"

type Settings struct {
	MonitoringState ProcessGroupMonitoringMode `json:"MonitoringState"` // Monitoring state
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"monitoring_state": {
			Type:        hcl.TypeString,
			Description: "Monitoring state",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"monitoring_state": me.MonitoringState,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"monitoring_state": &me.MonitoringState,
	})
}
