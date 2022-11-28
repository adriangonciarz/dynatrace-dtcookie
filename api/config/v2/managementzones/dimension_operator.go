package managementzones

type DimensionOperator string

var DimensionOperators = struct {
	Equals     DimensionOperator
	BeginsWith DimensionOperator
}{
	DimensionOperator("EQUALS"),
	DimensionOperator("BEGINS_WITH"),
}
