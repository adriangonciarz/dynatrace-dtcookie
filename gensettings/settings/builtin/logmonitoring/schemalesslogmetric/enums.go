package schemalesslogmetric

type Measure string

var Measures = struct {
	Occurrence Measure
	Attribute  Measure
}{
	Measure("OCCURRENCE"),
	Measure("ATTRIBUTE"),
}
