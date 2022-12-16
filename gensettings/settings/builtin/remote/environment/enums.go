package environment

type Scope string

var Scopes = struct {
	Cluster  Scope
	External Scope
	Internal Scope
}{
	Scope("CLUSTER"),
	Scope("EXTERNAL"),
	Scope("INTERNAL"),
}
