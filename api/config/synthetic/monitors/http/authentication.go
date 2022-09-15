package http

import "github.com/dtcookie/hcl"

// AuthenticationType is a type alias, nailing down currently supported AuthenticationTypes to `BASIC_AUTHENTICATION`, `NTLM` and `KERBEROS`. Additional values ARE however possible.
type AuthenticationType string

// AuthenticationTypes hints the currently supported AuthenticationTypes to `BASIC_AUTHENTICATION`, `NTLM` and `KERBEROS`. Additional values ARE however possible.
var AuthenticationTypes = struct {
	Basic    AuthenticationType
	NTML     AuthenticationType
	Kerberos AuthenticationType
}{
	AuthenticationType("BASIC_AUTHENTICATION"),
	AuthenticationType("NTML"),
	AuthenticationType("KERBEROS"),
}

// Authentication represents authentication options for a HTTP Request
type Authentication struct {
	Type        AuthenticationType `json:"type"`                // The type of authentication. Possible values are `BASIC_AUTHENTICATION`, `NTLM` and `KERBEROS`
	Credentials string             `json:"credentials"`         // The ID of the credentials within the Dynatrace Credentials Vault.
	RealmName   *string            `json:"realmName,omitempty"` // The Realm Name. Valid and required only if the type of authentication is `KERBEROS`
	KdcIP       *string            `json:"kdcIp,omitempty"`     // The KDC IP. Valid and required only if the type of authentication is `KERBEROS`
}

// Schema provides the schema map for the terraform provider
func (me *Authentication) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "The type of authentication. Possible values are `BASIC_AUTHENTICATION`, `NTLM` and `KERBEROS`.",
			Required:    true,
		},
		"credentials": {
			Type:        hcl.TypeString,
			Description: "The ID of the credentials within the Dynatrace Credentials Vault.",
			Required:    true,
		},
		"realm_name": {
			Type:        hcl.TypeString,
			Description: "The Realm Name. Valid and required only if the type of authentication is `KERBEROS`.",
			Optional:    true,
		},
		"kdc_ip": {
			Type:        hcl.TypeString,
			Description: "The KDC IP. Valid and required only if the type of authentication is `KERBEROS`.",
			Optional:    true,
		},
	}
}

// MarshalHCL serializes the fields of an Authentication struct into a map, using the keys specified within the Schema function
func (me *Authentication) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["type"] = string(me.Type)
	result["credentials"] = me.Credentials
	if me.RealmName != nil {
		result["realm_name"] = *me.RealmName
	}
	if me.KdcIP != nil {
		result["kdc_ip"] = *me.KdcIP
	}
	return result, nil
}

// UnmarshalHCL deserializes data available via terraform provider into the Authentication struct. The keys to be used are defined by the Schema function
func (me *Authentication) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("type", &me.Type); err != nil {
		return err
	}
	if err := decoder.Decode("credentials", &me.Credentials); err != nil {
		return err
	}
	if err := decoder.Decode("realm_name", &me.RealmName); err != nil {
		return err
	}
	if err := decoder.Decode("kdc_ip", &me.KdcIP); err != nil {
		return err
	}
	return nil
}
