package name

import "github.com/dtcookie/hcl"

type Settings struct {
	ApplicationName string          `json:"applicationName"` // Update application name
	ApplicationType ApplicationType `json:"applicationType"` // Update application type
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"application_name": {
			Type:        hcl.TypeString,
			Description: "Update application name",
			Required:    true,
		},
		"application_type": {
			Type:        hcl.TypeString,
			Description: "Update application type",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"application_name": me.ApplicationName,
		"application_type": me.ApplicationType,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"application_name": &me.ApplicationName,
		"application_type": &me.ApplicationType,
	})
}
