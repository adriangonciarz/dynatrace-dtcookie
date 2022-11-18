package vault

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

type ExternalVaultConfig struct {
	SourceAuthMethod                          SourceAuthMethod `json:"sourceAuthMethod"` // Defines the actual set of fields depending on the value. See one of the following objects: \n\n* `HASHICORP_VAULT_APPROLE` -> HashicorpApproleConfig \n* `HASHICORP_VAULT_CERTIFICATE` -> HashicorpCertificateConfig \n* `AZURE_KEY_VAULT_CLIENT_SECRET` -> AzureClientSecretConfig \n
	VaultURL                                  *string          `json:"vaultUrl,omitempty"`
	UsernameSecretName                        *string          `json:"usernameSecretName,omitempty"`
	PasswordSecretName                        *string          `json:"passwordSecretName,omitempty"`
	TokenSecretName                           *string          `json:"tokenSecretName,omitempty"`
	CredentialsUsedForExternalSynchronization []string         `json:"credentialsUsedForExternalSynchronization,omitempty"`

	// HashicorpApproleConfig
	PathtoCredentials *string `json:"pathToCredentials,omitempty"`
	RoleID            *string `json:"roleId,omitempty"`
	SecretID          *string `json:"secretId,omitempty"` // The ID of Credentials within the Certificate Vault holding the secret id
	VaultNameSpace    *string `json:"vaultNamespace,omitempty"`
	// HashicorpCertificateConfig
	Certificate *string `json:"certificate,omitempty"` // The ID of Credentials within the Certificate Vault holding the certificate
	// AzureClientSecret
	TenantID     *string `json:"tenantId,omitempty"`     // Tenant (directory) ID of Azure application in Azure Active Directory which has permission to access secrets in Azure Key Vault
	ClientID     *string `json:"clientId,omitempty"`     // Client (application) ID of Azure application in Azure Active Directory which has permission to access secrets in Azure Key Vault
	ClientSecret *string `json:"clientSecret,omitempty"` // Client secret generated for Azure application in Azure Active Directory used for proving identity when requesting a token used later for accessing secrets in Azure Key Vault
}

func (me *ExternalVaultConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"client_secret": {
			Type:        hcl.TypeString,
			Description: "Required for Azure Client Secret. No further documentation available",
			Optional:    true,
		},
		"clientid": {
			Type:        hcl.TypeString,
			Description: "Required for Azure Client Secret. No further documentation available",
			Optional:    true,
		},
		"tenantid": {
			Type:        hcl.TypeString,
			Description: "Required for Azure Client Secret. No further documentation available",
			Optional:    true,
		},
		"certificate": {
			Type:        hcl.TypeString,
			Description: "Required for Hashicorp Certificate. The ID of Credentials within the Certificate Vault holding the certificate",
			Optional:    true,
		},
		"vault_namespace": {
			Type:        hcl.TypeString,
			Description: "Required for Hashicorp App Role. No further documentation available",
			Optional:    true,
		},
		"secretid": {
			Type:        hcl.TypeString,
			Description: "Required for Hashicorp App Role. The ID of Credentials within the Certificate Vault holding the secret id",
			Optional:    true,
		},
		"roleid": {
			Type:        hcl.TypeString,
			Description: "Required for Hashicorp App Role. No further documentation available",
			Optional:    true,
		},
		"path_to_credentials": {
			Type:        hcl.TypeString,
			Description: "Required for Hashicorp App Role or Hashicorp Certificate. No further documentation available",
			Optional:    true,
		},
		"vault_url": {
			Type:        hcl.TypeString,
			Description: "No documentation available",
			Optional:    true,
		},
		"username_secret_name": {
			Type:        hcl.TypeString,
			Description: "No documentation available",
			Optional:    true,
		},
		"password_secret_name": {
			Type:        hcl.TypeString,
			Description: "No documentation available",
			Optional:    true,
		},
		"token_secret_name": {
			Type:        hcl.TypeString,
			Description: "No documentation available",
			Optional:    true,
		},
		"credentials_used_for_external_synchronization": {
			Type:        hcl.TypeSet,
			Description: "No documentation available",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
			Optional:    true,
		},
	}
}

func (me *ExternalVaultConfig) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if me.TenantID != nil {
		result["tenantid"] = me.TenantID
	}
	if me.ClientID != nil {
		result["clientid"] = me.ClientID
	}
	if me.ClientSecret != nil {
		result["client_secret"] = me.ClientSecret
	}

	if me.PathtoCredentials != nil {
		result["path_to_credentials"] = me.PathtoCredentials
	}
	if me.RoleID != nil {
		result["roleid"] = me.RoleID
	}
	if me.SecretID != nil {
		result["secretid"] = me.SecretID
	}
	if me.VaultNameSpace != nil {
		result["vault_namespace"] = me.VaultNameSpace
	}
	if me.Certificate != nil {
		result["certificate"] = me.Certificate
	}

	if me.VaultURL != nil {
		result["vault_url"] = me.VaultURL
	}
	if me.UsernameSecretName != nil {
		result["username_secret_name"] = me.UsernameSecretName
	}
	if me.PasswordSecretName != nil {
		result["password_secret_name"] = me.PasswordSecretName
	}
	if me.TokenSecretName != nil {
		result["token_secret_name"] = me.TokenSecretName
	}
	if len(me.CredentialsUsedForExternalSynchronization) > 0 {
		result["credentials_used_for_external_synchronization"] = me.CredentialsUsedForExternalSynchronization
	}

	return result, nil
}

func (me *ExternalVaultConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("client_secret"); ok {
		me.ClientSecret = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("clientid"); ok {
		me.ClientID = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("tenantid"); ok {
		me.TenantID = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("certificate"); ok {
		me.Certificate = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("vault_namespace"); ok {
		me.VaultNameSpace = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("secretid"); ok {
		me.SecretID = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("roleid"); ok {
		me.RoleID = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("path_to_credentials"); ok {
		me.PathtoCredentials = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("vault_url"); ok {
		me.VaultURL = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("username_secret_name"); ok {
		me.UsernameSecretName = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("password_secret_name"); ok {
		me.PasswordSecretName = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("token_secret_name"); ok {
		me.TokenSecretName = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("credentials_used_for_external_synchronization"); ok {
		me.CredentialsUsedForExternalSynchronization = []string{}
		for _, elem := range value.(hcl.Set).List() {
			me.CredentialsUsedForExternalSynchronization = append(me.CredentialsUsedForExternalSynchronization, elem.(string))
		}
	}
	if me.ClientID != nil || me.ClientSecret != nil || me.TenantID != nil {
		me.SourceAuthMethod = SourceAuthMethods.AzureKeyVaultClientSecret
	} else if me.RoleID != nil {
		me.SourceAuthMethod = SourceAuthMethods.HashicorpVaultAppRole
	} else if me.Certificate != nil {
		me.SourceAuthMethod = SourceAuthMethods.HashicoprVaultCertificate
	}
	return nil
}

type ExternalVaultConfigType string

var ExternalVaultConfigTypes = struct {
	AzureCertificateModel    ExternalVaultConfigType
	AzureClientSecretModel   ExternalVaultConfigType
	HashcorpAppRoleModel     ExternalVaultConfigType
	HashcorpCertificateModel ExternalVaultConfigType
}{
	ExternalVaultConfigType("AZURE_CERTIFICATE_MODEL"),
	ExternalVaultConfigType("AZURE_CLIENT_SECRET_MODEL"),
	ExternalVaultConfigType("HASHICORP_APPROLE_MODEL"),
	ExternalVaultConfigType("HASHICORP_CERTIFICATE_MODEL"),
}

type SourceAuthMethod string

var SourceAuthMethods = struct {
	AzureKeyVaultClientSecret SourceAuthMethod
	HashicorpVaultAppRole     SourceAuthMethod
	HashicoprVaultCertificate SourceAuthMethod
}{
	SourceAuthMethod("AZURE_KEY_VAULT_CLIENT_SECRET"),
	SourceAuthMethod("HASHICORP_VAULT_APPROLE"),
	SourceAuthMethod("HASHICORP_VAULT_CERTIFICATE"),
}
