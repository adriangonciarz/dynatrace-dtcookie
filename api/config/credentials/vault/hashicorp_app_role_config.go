package vault

type HashicorpAppRoleConfig struct {
	ExternalVaultConfig
	PathtoCredentials string `json:"pathToCredentials"`
	RoleID            string `json:"roleId"`
	SecretID          string `json:"secretId"`
	VaultNameSpace    string `json:"vaultNamespace"`
}
