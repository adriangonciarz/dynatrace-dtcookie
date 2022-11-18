package vault

type HashicorpCertificateConfig struct {
	ExternalVaultConfig
	PathToCredentials string `json:"pathToCredentials"`
	Certificate       string `json:"certificate"`
}
