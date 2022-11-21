package vault

import (
	"encoding/json"
	"sort"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

type Credentials struct {
	ID                     *string                `json:"id,omitempty"`
	Name                   string                 `json:"name"`                             // The name of the credentials set.
	Description            *string                `json:"description,omitempty"`            // A short description of the credentials set..
	OwnerAccessOnly        bool                   `json:"ownerAccessOnly"`                  // The credentials set is available to every user (`false`) or to owner only (`true`).
	Scope                  Scope                  `json:"scope"`                            // The scope of the credentials set
	Type                   CredentialsType        `json:"type"`                             // Defines the actual set of fields depending on the value. See one of the following objects: \n\n* `CERTIFICATE` -> CertificateCredentials \n* `PUBLIC_CERTIFICATE` -> PublicCertificateCredentials \n* `USERNAME_PASSWORD` -> UserPasswordCredentials \n* `TOKEN` -> TokenCredentials \n
	Token                  *string                `json:"token,omitempty"`                  // Token in the string format.
	Password               *string                `json:"password,omitempty"`               // The password of the credential (Base64 encoded).
	Username               *string                `json:"user,omitempty"`                   // The username of the credentials set.
	Certificate            *string                `json:"certificate,omitempty"`            // The certificate in the string (Base64) format.
	CertificateFormat      *CertificateFormat     `json:"certificateFormat,omitempty"`      // The certificate format.
	ExternalVault          *ExternalVaultConfig   `json:"externalVault,omitempty"`          // Configuration for external vault synchronization
	CredentialUsageSummary CredentialUsageSummary `json:"credentialUsageSummary,omitempty"` //The list contains summary data related to the use of credentials
}

func (me *Credentials) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the credentials set",
			Required:    true,
		},
		"description": {
			Type:        hcl.TypeString,
			Description: "A short description of the credentials set",
			Optional:    true,
		},
		"owner_access_only": {
			Type:        hcl.TypeBool,
			Description: "The credentials set is available to every user (`false`) or to owner only (`true`)",
			Optional:    true,
		},
		"public": {
			Type:          hcl.TypeBool,
			Description:   "For certificate authentication specifies whether it's public certificate auth (`true`) or not (`false`).",
			ConflictsWith: []string{"username", "token"},
			Optional:      true,
		},
		"scope": {
			Type:        hcl.TypeString,
			Description: "The scope of the credentials set. Possible values are `ALL`, `EXTENSION` and `SYNTHETIC`",
			Required:    true,
		},
		"token": {
			Type:          hcl.TypeString,
			Description:   "Token in the string format. Specifying a token implies `Token Authentication`.",
			ConflictsWith: []string{"username", "password", "certificate", "format", "public"},
			Sensitive:     true,
			Optional:      true,
		},
		"username": {
			Type:          hcl.TypeString,
			Description:   "The username of the credentials set.",
			ConflictsWith: []string{"token", "public", "certificate"},
			RequiredWith:  []string{"password"},
			Sensitive:     true,
			Optional:      true,
		},
		"password": {
			Type:          hcl.TypeString,
			Description:   "The password of the credential.",
			ConflictsWith: []string{"token"},
			Sensitive:     true,
			Optional:      true,
		},
		"certificate": {
			Type:          hcl.TypeString,
			Description:   "The certificate in the string format.",
			ConflictsWith: []string{"token", "username"},
			RequiredWith:  []string{"format", "password"},
			Optional:      true,
		},
		"format": {
			Type:          hcl.TypeString,
			Description:   "The certificate format. Possible values are `PEM`, `PKCS12` and `UNKNOWN`.",
			ConflictsWith: []string{"token", "username"},
			RequiredWith:  []string{"certificate"},
			Optional:      true,
		},
		"external": {
			Type:        hcl.TypeList,
			Description: "External Vault Configuration",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(ExternalVaultConfig).Schema()},
		},
		"credential_usage_summary": {
			Type:        hcl.TypeList,
			Description: "The list contains summary data related to the use of credentials",
			Optional:    true,
			MaxItems:    2,
			Elem:        &hcl.Resource{Schema: new(CredentialUsageObj).Schema()},
		},
	}
}

func (me *Credentials) EnsurePredictableOrder() {
	conds := CredentialUsageSummary{}
	condStrings := sort.StringSlice{}
	for _, entry := range me.CredentialUsageSummary {
		condBytes, _ := json.Marshal(entry)
		condStrings = append(condStrings, string(condBytes))
	}
	condStrings.Sort()
	for _, condString := range condStrings {
		cond := CredentialUsageObj{}
		json.Unmarshal([]byte(condString), &cond)
		conds = append(conds, &cond)
	}
	me.CredentialUsageSummary = conds
}

func (me *Credentials) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["name"] = me.Name
	if me.Description != nil && len(*me.Description) > 0 {
		result["description"] = *me.Description
	}
	if me.OwnerAccessOnly {
		result["owner_access_only"] = me.OwnerAccessOnly
	}
	result["scope"] = string(me.Scope)
	if me.Token != nil && len(*me.Token) > 0 {
		result["token"] = *me.Token
	}
	if me.Password != nil && len(*me.Password) > 0 {
		result["password"] = *me.Password
	}
	if me.Username != nil && len(*me.Username) > 0 {
		result["username"] = *me.Username
	}
	if me.Certificate != nil && len(*me.Certificate) > 0 {
		result["certificate"] = *me.Certificate
	}
	if me.CertificateFormat != nil && len(*me.CertificateFormat) > 0 {
		result["format"] = *me.CertificateFormat
	}
	if me.ExternalVault != nil {
		marshalled, err := me.ExternalVault.MarshalHCL()
		if err != nil {
			return nil, err
		}
		result["external"] = []interface{}{marshalled}
	}
	if me.CredentialUsageSummary != nil {
		entries := []interface{}{}
		for _, entry := range me.CredentialUsageSummary {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["credential_usage_summary"] = entries
	}

	switch me.Type {
	case CredentialsTypes.PublicCertificate:
		result["public"] = true
	}
	return result, nil
}

func (me *Credentials) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("external.#"); ok && value.(int) == 1 {
		me.ExternalVault = new(ExternalVaultConfig)
		me.ExternalVault.UnmarshalHCL(hcl.NewDecoder(decoder, "external.0"))
	}
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	if value, ok := decoder.GetOk("description"); ok {
		me.Description = opt.NewString(value.(string))
		if len(*me.Description) == 0 {
			me.Description = nil
		}
	}
	if value, ok := decoder.GetOk("owner_access_only"); ok {
		me.OwnerAccessOnly = value.(bool)
	}
	if value, ok := decoder.GetOk("scope"); ok {
		me.Scope = Scope(value.(string))
	}
	if value, ok := decoder.GetOk("token"); ok {
		me.Token = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("password"); ok {
		me.Password = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("username"); ok {
		me.Username = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("certificate"); ok {
		me.Certificate = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("format"); ok {
		me.CertificateFormat = CertificateFormat(value.(string)).Ref()
	}
	if me.Username != nil {
		me.Type = CredentialsTypes.UsernamePassword
	} else if me.Token != nil {
		me.Type = CredentialsTypes.Token
	} else if me.Certificate != nil || me.CertificateFormat != nil {
		if value, ok := decoder.GetOk("public"); ok && value.(bool) {
			me.Type = CredentialsTypes.PublicCertificate
		} else {
			me.Type = CredentialsTypes.Certificate
		}
	}
	if result, ok := decoder.GetOk("credential_usage_summary.#"); ok {
		me.CredentialUsageSummary = []*CredentialUsageObj{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(CredentialUsageObj)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "credential_usage_summary", idx)); err != nil {
				return err
			}
			me.CredentialUsageSummary = append(me.CredentialUsageSummary, entry)
		}
	}
	return nil
}
