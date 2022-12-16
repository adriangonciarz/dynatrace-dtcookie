package usersessionexportsettingsv2

import "github.com/dtcookie/hcl"

type BasicAuth struct {
	Password string `json:"password"`
	Username string `json:"username"` // User name
}

func (me *BasicAuth) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"password": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"username": {
			Type:        hcl.TypeString,
			Description: "User name",
			Required:    true,
		},
	}
}

func (me *BasicAuth) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"password": me.Password,
		"username": me.Username,
	})
}

func (me *BasicAuth) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"password": &me.Password,
		"username": &me.Username,
	})
}
