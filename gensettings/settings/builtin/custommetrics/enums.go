package custommetrics

type ValueType string

var ValueTypes = struct {
	Field   ValueType
	Counter ValueType
}{
	ValueType("FIELD"),
	ValueType("COUNTER"),
}

type Operator string

var Operators = struct {
	Like                 Operator
	LessThan             Operator
	NotEqual             Operator
	StartsWith           Operator
	LessThanOrEqualTo    Operator
	GreaterThan          Operator
	Equals               Operator
	IsNull               Operator
	NotLike              Operator
	GreaterThanOrEqualTo Operator
	IsNotNull            Operator
	In                   Operator
}{
	Operator("LIKE"),
	Operator("LESS_THAN"),
	Operator("NOT_EQUAL"),
	Operator("STARTS_WITH"),
	Operator("LESS_THAN_OR_EQUAL_TO"),
	Operator("GREATER_THAN"),
	Operator("EQUALS"),
	Operator("IS_NULL"),
	Operator("NOT_LIKE"),
	Operator("GREATER_THAN_OR_EQUAL_TO"),
	Operator("IS_NOT_NULL"),
	Operator("IN"),
}
