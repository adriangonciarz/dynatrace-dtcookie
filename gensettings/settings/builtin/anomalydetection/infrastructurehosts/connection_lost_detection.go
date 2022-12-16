package infrastructurehosts

import "github.com/dtcookie/hcl"

type ConnectionLostDetection struct {
	Enabled             bool                               `json:"enabled"`             // Detect host or monitoring connection lost problems
	OnGracefulShutdowns ConnectionLostDetectionSensitivity `json:"onGracefulShutdowns"` // Graceful host shutdowns
}

func (me *ConnectionLostDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect host or monitoring connection lost problems",
			Optional:    true,
		},
		"on_graceful_shutdowns": {
			Type:        hcl.TypeString,
			Description: "Graceful host shutdowns",
			Required:    true,
		},
	}
}

func (me *ConnectionLostDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":               me.Enabled,
		"on_graceful_shutdowns": me.OnGracefulShutdowns,
	})
}

func (me *ConnectionLostDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":               &me.Enabled,
		"on_graceful_shutdowns": &me.OnGracefulShutdowns,
	})
}
