package simpledetectionrule

type DetectionRuleType string

var DetectionRuleTypes = struct {
	Prop DetectionRuleType
	Env  DetectionRuleType
}{
	DetectionRuleType("Prop"),
	DetectionRuleType("Env"),
}
