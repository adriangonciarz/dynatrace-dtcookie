package eulasettings

import "github.com/dtcookie/hcl"

type Settings struct {
	EnableEula bool `json:"enableEula"` // Display end user terms to new users logging in to the environment
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enable_eula": {
			Type:        hcl.TypeBool,
			Description: "Display end user terms to new users logging in to the environment",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enable_eula": me.EnableEula,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enable_eula": &me.EnableEula,
	})
}
