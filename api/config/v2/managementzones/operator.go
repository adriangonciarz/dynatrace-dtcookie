package managementzones

type Operator string

var Operators = struct {
	NotGreaterThan        Operator
	NotRegexMatches       Operator
	IsIpInRange           Operator
	NotIsIpInRange        Operator
	LowerThan             Operator
	TagKeyEquals          Operator
	NotEndsWith           Operator
	NotTagKeyEquals       Operator
	NotGreaterThanOrEqual Operator
	LowerThanOrEqual      Operator
	Exists                Operator
	Contains              Operator
	RegexMatches          Operator
	EndsWith              Operator
	GreaterThanOrEqual    Operator
	NotExists             Operator
	NotContains           Operator
	NotLowerThanOrEqual   Operator
	NotEquals             Operator
	Equals                Operator
	BeginsWith            Operator
	NotBeginsWith         Operator
	NotLowerThan          Operator
	GreaterThan           Operator
}{
	Operator("NOT_GREATER_THAN"),
	Operator("NOT_REGEX_MATCHES"),
	Operator("IS_IP_IN_RANGE"),
	Operator("NOT_IS_IP_IN_RANGE"),
	Operator("LOWER_THAN"),
	Operator("TAG_KEY_EQUALS"),
	Operator("NOT_ENDS_WITH"),
	Operator("NOT_TAG_KEY_EQUALS"),
	Operator("NOT_GREATER_THAN_OR_EQUAL"),
	Operator("LOWER_THAN_OR_EQUAL"),
	Operator("EXISTS"),
	Operator("CONTAINS"),
	Operator("REGEX_MATCHES"),
	Operator("ENDS_WITH"),
	Operator("GREATER_THAN_OR_EQUAL"),
	Operator("NOT_EXISTS"),
	Operator("NOT_CONTAINS"),
	Operator("NOT_LOWER_THAN_OR_EQUAL"),
	Operator("NOT_EQUALS"),
	Operator("EQUALS"),
	Operator("BEGINS_WITH"),
	Operator("NOT_BEGINS_WITH"),
	Operator("NOT_LOWER_THAN"),
	Operator("GREATER_THAN"),
}
