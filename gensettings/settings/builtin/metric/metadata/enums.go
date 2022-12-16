package metadata

type UnitDisplayFormat string

var UnitDisplayFormats = struct {
	Binary  UnitDisplayFormat
	Decimal UnitDisplayFormat
}{
	UnitDisplayFormat("Binary"),
	UnitDisplayFormat("Decimal"),
}

type ValueType string

var ValueTypes = struct {
	Error   ValueType
	Score   ValueType
	Unknown ValueType
}{
	ValueType("Error"),
	ValueType("Score"),
	ValueType("Unknown"),
}
