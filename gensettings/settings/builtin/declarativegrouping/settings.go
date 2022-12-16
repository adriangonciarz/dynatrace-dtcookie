package declarativegrouping

import "github.com/dtcookie/hcl"

type Settings struct {
	Detection ProcessDefinitions `json:"detection"` // Enter a descriptive process group display name and a unique identifier that Dynatrace can use to recognize this process group.
	Enabled   bool               `json:"enabled"`   // Enabled
	Name      string             `json:"name"`      // Monitored technology name
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"detection": {
			Type:        hcl.TypeList,
			Description: "Enter a descriptive process group display name and a unique identifier that Dynatrace can use to recognize this process group.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ProcessDefinitions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "Monitored technology name",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"detection": me.Detection,
		"enabled":   me.Enabled,
		"name":      me.Name,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"detection": &me.Detection,
		"enabled":   &me.Enabled,
		"name":      &me.Name,
	})
}
