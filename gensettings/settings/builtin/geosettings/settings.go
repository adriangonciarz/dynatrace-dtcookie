package geosettings

import "github.com/dtcookie/hcl"

type Settings struct {
	DisplayWorldmap bool `json:"displayWorldmap"` // Display the world map
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"display_worldmap": {
			Type:        hcl.TypeBool,
			Description: "Display the world map",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"display_worldmap": me.DisplayWorldmap,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"display_worldmap": &me.DisplayWorldmap,
	})
}
