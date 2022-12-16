package custommetrics

type Operator string

var Operators = struct {
	Equals               Operator
	GreaterThan          Operator
	GreaterThanOrEqualTo Operator
	In                   Operator
	IsNotNull            Operator
	IsNull               Operator
	LessThan             Operator
	LessThanOrEqualTo    Operator
	Like                 Operator
	NotEqual             Operator
	NotLike              Operator
	StartsWith           Operator
}{
	Operator("EQUALS"),
	Operator("GREATER_THAN"),
	Operator("GREATER_THAN_OR_EQUAL_TO"),
	Operator("IN"),
	Operator("IS_NOT_NULL"),
	Operator("IS_NULL"),
	Operator("LESS_THAN"),
	Operator("LESS_THAN_OR_EQUAL_TO"),
	Operator("LIKE"),
	Operator("NOT_EQUAL"),
	Operator("NOT_LIKE"),
	Operator("STARTS_WITH"),
}

type ValueType string

var ValueTypes = struct {
	Counter ValueType
	Field   ValueType
}{
	ValueType("COUNTER"),
	ValueType("FIELD"),
}
