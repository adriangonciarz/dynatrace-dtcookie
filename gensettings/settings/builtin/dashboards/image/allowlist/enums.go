package allowlist

type RuleEnum string

var RuleEnums = struct {
	Startswith RuleEnum
	Equals     RuleEnum
}{
	RuleEnum("StartsWith"),
	RuleEnum("Equals"),
}
