package metadata

type ValueType string

var ValueTypes = struct {
	Score   ValueType
	Unknown ValueType
	Error   ValueType
}{
	ValueType("Score"),
	ValueType("Unknown"),
	ValueType("Error"),
}

type UnitDisplayFormat string

var UnitDisplayFormats = struct {
	Binary  UnitDisplayFormat
	Decimal UnitDisplayFormat
}{
	UnitDisplayFormat("Binary"),
	UnitDisplayFormat("Decimal"),
}
