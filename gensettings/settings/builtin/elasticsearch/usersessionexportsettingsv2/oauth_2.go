package usersessionexportsettingsv2

import "github.com/dtcookie/hcl"

type OAuth2 struct {
	AccessTokenUrl  string          `json:"accessTokenUrl"`  // Access token URL
	ClientID        string          `json:"clientId"`        // Client ID
	ClientSecret    string          `json:"clientSecret"`    // Client secret
	GrantType       GrantType       `json:"grantType"`       // Grant type
	Scope           string          `json:"scope"`           // The scope of access you are requesting
	SendCredentials SendCredentials `json:"sendCredentials"` // Send credentials
}

func (me *OAuth2) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"access_token_url": {
			Type:        hcl.TypeString,
			Description: "Access token URL",
			Required:    true,
		},
		"client_id": {
			Type:        hcl.TypeString,
			Description: "Client ID",
			Required:    true,
		},
		"client_secret": {
			Type:        hcl.TypeString,
			Description: "Client secret",
			Required:    true,
		},
		"grant_type": {
			Type:        hcl.TypeString,
			Description: "Grant type",
			Required:    true,
		},
		"scope": {
			Type:        hcl.TypeString,
			Description: "The scope of access you are requesting",
			Optional:    true,
		},
		"send_credentials": {
			Type:        hcl.TypeString,
			Description: "Send credentials",
			Required:    true,
		},
	}
}

func (me *OAuth2) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"access_token_url": me.AccessTokenUrl,
		"client_id":        me.ClientID,
		"client_secret":    me.ClientSecret,
		"grant_type":       me.GrantType,
		"scope":            me.Scope,
		"send_credentials": me.SendCredentials,
	})
}

func (me *OAuth2) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"access_token_url": &me.AccessTokenUrl,
		"client_id":        &me.ClientID,
		"client_secret":    &me.ClientSecret,
		"grant_type":       &me.GrantType,
		"scope":            &me.Scope,
		"send_credentials": &me.SendCredentials,
	})
}
