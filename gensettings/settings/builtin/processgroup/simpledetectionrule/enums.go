package simpledetectionrule

type DetectionRuleType string

var DetectionRuleTypes = struct {
	Env  DetectionRuleType
	Prop DetectionRuleType
}{
	DetectionRuleType("Env"),
	DetectionRuleType("Prop"),
}
