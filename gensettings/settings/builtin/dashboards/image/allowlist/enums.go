package allowlist

type RuleEnum string

var RuleEnums = struct {
	Equals     RuleEnum
	Startswith RuleEnum
}{
	RuleEnum("Equals"),
	RuleEnum("StartsWith"),
}
