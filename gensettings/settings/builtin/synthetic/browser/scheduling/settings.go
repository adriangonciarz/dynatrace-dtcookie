package scheduling

import "github.com/dtcookie/hcl"

type Settings struct {
	Frequency int       `json:"frequency"` // How often the monitor is executed. Supported values are 5, 10, 15, 30, 60, 120 and 240 minutes
	Locations Locations `json:"locations"` // Locations
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"frequency": {
			Type:        hcl.TypeInt,
			Description: "How often the monitor is executed. Supported values are 5, 10, 15, 30, 60, 120 and 240 minutes",
			Required:    true,
		},
		"locations": {
			Type:        hcl.TypeList,
			Description: "Locations",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Locations).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"frequency": me.Frequency,
		"locations": me.Locations,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"frequency": &me.Frequency,
		"locations": &me.Locations,
	})
}
