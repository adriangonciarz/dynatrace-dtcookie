package usersessionexportsettingsv2

type AuthType string

var AuthTypes = struct {
	Basic  AuthType
	Oauth2 AuthType
}{
	AuthType("Basic"),
	AuthType("Oauth2"),
}

type GrantType string

var GrantTypes = struct {
	Clientcredentials GrantType
}{
	GrantType("ClientCredentials"),
}

type SendCredentials string

var SendCredentialss = struct {
	Header SendCredentials
}{
	SendCredentials("Header"),
}
