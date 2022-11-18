package vault

type CredentialsType string

var CredentialsTypes = struct {
	Certificate       CredentialsType
	PublicCertificate CredentialsType
	Token             CredentialsType
	UsernamePassword  CredentialsType
	Unknown           CredentialsType
}{
	"CERTIFICATE",
	"PUBLIC_CERTIFICATE",
	"TOKEN",
	"USERNAME_PASSWORD",
	"UNKNOWN",
}
