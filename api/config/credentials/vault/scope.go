package vault

type Scope string

var Scopes = struct {
	All       Scope
	Extension Scope
	Synthetic Scope
	Unknown   Scope
}{
	Scope("ALL"),
	Scope("EXTENSION"),
	Scope("SYNTHETIC"),
	Scope("UNKNOWN"),
}
