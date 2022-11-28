package managementzones

type DimensionType string

var DimensionTypes = struct {
	Metric DimensionType
	Any    DimensionType
	Log    DimensionType
}{
	DimensionType("METRIC"),
	DimensionType("ANY"),
	DimensionType("LOG"),
}
