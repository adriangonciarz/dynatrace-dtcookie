package name

import "github.com/dtcookie/hcl"

type Settings struct {
	ApplicationName string `json:"applicationName"` // Update application name
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"application_name": {
			Type:        hcl.TypeString,
			Description: "Update application name",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"application_name": me.ApplicationName,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"application_name": &me.ApplicationName,
	})
}
