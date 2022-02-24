package cloudfoundry

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// CloudFoundryCredentials Configuration for specific Cloud Foundry credentials.
type CloudFoundryCredentials struct {
	LoginURL string                     `json:"loginUrl"`           // The login URL of the Cloud Foundry foundation credentials.  The URL must be valid according to RFC 2396.  Leading or trailing whitespaces are not allowed.
	Password *string                    `json:"password,omitempty"` // The password of the Cloud Foundry foundation credentials.
	Active   bool                       `json:"active"`             // The monitoring is enabled (`true`) or disabled (`false`) for given credentials configuration.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.
	Name     string                     `json:"name"`               // The name of the Cloud Foundry foundation credentials.  Allowed characters are letters, numbers, whitespaces, and the following characters: `.+-_`. Leading or trailing whitespace is not allowed.
	Username string                     `json:"username"`           // The username of the Cloud Foundry foundation credentials.  Leading and trailing whitespaces are not allowed.
	APIURL   string                     `json:"apiUrl"`             // The URL of the Cloud Foundry foundation credentials.  The URL must be valid according to RFC 2396.  Leading or trailing whitespaces are not allowed.
	Unknowns map[string]json.RawMessage `json:"-"`
}

func (me *CloudFoundryCredentials) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"login_url": {
			Type:        hcl.TypeString,
			Description: "The login URL of the Cloud Foundry foundation credentials. The URL must be valid according to RFC 2396.  Leading or trailing whitespaces are not allowed.",
			Required:    true,
		},
		"api_url": {
			Type:        hcl.TypeString,
			Description: "The URL of the Cloud Foundry foundation credentials.  The URL must be valid according to RFC 2396.  Leading or trailing whitespaces are not allowed.",
			Required:    true,
		},
		"username": {
			Type:        hcl.TypeString,
			Description: "The username of the Cloud Foundry foundation credentials.  Leading and trailing whitespaces are not allowed.",
			Required:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the Cloud Foundry foundation credentials.  Allowed characters are letters, numbers, whitespaces, and the following characters: `.+-_`. Leading or trailing whitespace is not allowed.",
			Required:    true,
		},
		"password": {
			Type:        hcl.TypeString,
			Description: "The password of the Cloud Foundry foundation credentials.",
			Optional:    true,
			Sensitive:   true,
		},
		"active": {
			Type:        hcl.TypeBool,
			Description: "The monitoring is enabled (`true`) or disabled (`false`) for given credentials configuration.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (me *CloudFoundryCredentials) MarshalHCL() (map[string]interface{}, error) {
	properties, err := hcl.NewProperties(me, me.Unknowns)
	if err != nil {
		return nil, err
	}
	return properties.EncodeAll(map[string]interface{}{
		"login_url": me.LoginURL,
		"api_url":   me.APIURL,
		"password":  me.Password,
		"active":    me.Active,
		"name":      me.Name,
		"username":  me.Username,
		"unknowns":  me.Unknowns,
	})
}

func (me *CloudFoundryCredentials) UnmarshalHCL(decoder hcl.Decoder) error {
	err := decoder.DecodeAll(map[string]interface{}{
		"login_url": &me.LoginURL,
		"api_url":   &me.APIURL,
		"password":  &me.Password,
		"active":    &me.Active,
		"name":      &me.Name,
		"username":  &me.Username,
		"unknowns":  &me.Unknowns,
	})
	return err
}

func (me *CloudFoundryCredentials) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"loginUrl": me.LoginURL,
		"apiUrl":   me.APIURL,
		"password": me.Password,
		"active":   me.Active,
		"name":     me.Name,
		"username": me.Username,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *CloudFoundryCredentials) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	err := properties.UnmarshalAll(map[string]interface{}{
		"loginUrl": &me.LoginURL,
		"apiUrl":   &me.APIURL,
		"password": &me.Password,
		"active":   &me.Active,
		"name":     &me.Name,
		"username": &me.Username,
	})
	return err
}
