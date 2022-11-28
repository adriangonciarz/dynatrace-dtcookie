package managementzones

type RuleType string

var RuleTypes = struct {
	Me        RuleType
	Dimension RuleType
	Selector  RuleType
}{
	RuleType("ME"),
	RuleType("DIMENSION"),
	RuleType("SELECTOR"),
}
