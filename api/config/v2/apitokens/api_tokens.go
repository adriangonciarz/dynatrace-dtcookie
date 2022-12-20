package apitokens

import (
	"github.com/dtcookie/hcl"
)

// ApiToken TODO: documentation
type ApiToken struct {
	ID                  *string  `json:"id,omitempty"`                  // The ID of the api tokem
	Name                string   `json:"name"`                          // The name of the token.
	Enabled             *bool    `json:"enabled,omitempty"`             // The token is enabled (true) or disabled (false), default disabled (false).
	PersonalAccessToken *bool    `json:"personalAccessToken,omitempty"` // The token is a personal access token (true) or an API token (false).
	ExpirationDate      *string  `json:"expirationDate,omitempty"`      // The expiration date of the token.
	Owner               *string  `json:"owner,omitempty"`               // The owner of the token
	CreationDate        *string  `json:"creationDate,omitempty"`        // Token creation date in ISO 8601 format (yyyy-MM-dd'T'HH:mm:ss.SSS'Z')
	ModifiedDate        *string  `json:"modifiedDate,omitempty"`        // Token last modified date in ISO 8601 format (yyyy-MM-dd'T'HH:mm:ss.SSS'Z').
	LastUsedDate        *string  `json:"lastUsedDate,omitempty"`        // Token last used date in ISO 8601 format (yyyy-MM-dd'T'HH:mm:ss.SSS'Z')
	LastUsedIpAddress   *string  `json:"lastUsedIpAddress,omitempty"`   // Token last used IP address.
	Scopes              []string `json:"scopes"`                        // A list of the scopes to be assigned to the token.
	Token               *string  `json:"token,omitempty"`               // The secret of the token.
}

type TokenList struct {
	ApiTokens []*ApiToken `json:"apiTokens"` // An ordered list of api tokens
}

func (me *ApiToken) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the token.",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "The token is enabled (true) or disabled (false), default disabled (false).",
			Optional:    true,
		},
		"personal_access_token": {
			Type:        hcl.TypeBool,
			Description: "The token is a personal access token (true) or an API token (false).",
			Optional:    true,
		},
		"expiration_date": {
			Type:        hcl.TypeString,
			Description: "The expiration date of the token.",
			Optional:    true,
		},
		"owner": {
			Type:        hcl.TypeString,
			Description: "The owner of the token",
			Optional:    true,
			Computed:    true,
		},
		"creation_date": {
			Type:        hcl.TypeString,
			Description: "Token creation date in ISO 8601 format (yyyy-MM-dd'T'HH:mm:ss.SSS'Z')",
			Optional:    true,
			Computed:    true,
		},
		"modified_date": {
			Type:        hcl.TypeString,
			Description: "Token last modified date in ISO 8601 format (yyyy-MM-dd'T'HH:mm:ss.SSS'Z').",
			Optional:    true,
			Computed:    true,
		},
		"last_used_date": {
			Type:        hcl.TypeString,
			Description: "Token last used date in ISO 8601 format (yyyy-MM-dd'T'HH:mm:ss.SSS'Z')",
			Optional:    true,
			Computed:    true,
		},
		"last_used_ip_address": {
			Type:        hcl.TypeString,
			Description: "Token last used IP address.",
			Optional:    true,
			Computed:    true,
		},
		"scopes": {
			Type:        hcl.TypeSet,
			Description: "A list of the scopes to be assigned to the token.",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"token": {
			Type:        hcl.TypeString,
			Description: "The secret of the token.",
			Sensitive:   true,
			Computed:    true,
		},
	}
}

func (me *ApiToken) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	if _, err := properties.EncodeAll(map[string]interface{}{
		"name":                  me.Name,
		"enabled":               me.Enabled,
		"personal_access_token": me.PersonalAccessToken,
		"expiration_date":       me.ExpirationDate,
		"owner":                 me.Owner,
		"creation_date":         me.CreationDate,
		"modified_date":         me.ModifiedDate,
		"last_used_date":        me.LastUsedDate,
		"last_used_ip_address":  me.LastUsedIpAddress,
		"scopes":                me.Scopes,
	}); err != nil {
		return nil, err
	}
	if me.Token != nil {
		if err := properties.Encode("token", me.Token); err != nil {
			return nil, err
		}
	}
	return properties, nil
}

func (me *ApiToken) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":                  &me.Name,
		"enabled":               &me.Enabled,
		"personal_access_token": &me.PersonalAccessToken,
		"expiration_date":       &me.ExpirationDate,
		"scopes":                &me.Scopes,
	})
}
