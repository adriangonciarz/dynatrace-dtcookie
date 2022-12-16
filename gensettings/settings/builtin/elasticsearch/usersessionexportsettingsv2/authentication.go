package usersessionexportsettingsv2

import "github.com/dtcookie/hcl"

type Authentication struct {
	Active    bool       `json:"active"`    // Activate
	AuthType  AuthType   `json:"authType"`  // Authentication type
	BasicAuth *BasicAuth `json:"basicAuth"` // Basic authentication
	OAuth2    *OAuth2    `json:"oAuth2"`    // OAuth 2.0 (Early Adopter)
}

func (me *Authentication) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"active": {
			Type:        hcl.TypeBool,
			Description: "Activate",
			Optional:    true,
		},
		"auth_type": {
			Type:        hcl.TypeString,
			Description: "Authentication type",
			Required:    true,
		},
		"basic_auth": {
			Type:        hcl.TypeList,
			Description: "Basic authentication",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(BasicAuth).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"o_auth_2": {
			Type:        hcl.TypeList,
			Description: "OAuth 2.0 (Early Adopter)",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(OAuth2).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Authentication) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"active":     me.Active,
		"auth_type":  me.AuthType,
		"basic_auth": me.BasicAuth,
		"o_auth_2":   me.OAuth2,
	})
}

func (me *Authentication) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"active":     &me.Active,
		"auth_type":  &me.AuthType,
		"basic_auth": &me.BasicAuth,
		"o_auth_2":   &me.OAuth2,
	})
}
