package networkzones

import (
	"github.com/dtcookie/hcl"
)

type NetworkZones struct {
	Enabled bool `json:"enabled"` // Network Zones are enabled (`true`) or disabled (`false`)
}

func (me *NetworkZones) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Network Zones are enabled (`true`) or disabled (`false`)",
		},
	}
}

func (me *NetworkZones) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"enabled": me.Enabled,
	})
}

func (me *NetworkZones) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"enabled": &me.Enabled,
	})
}
