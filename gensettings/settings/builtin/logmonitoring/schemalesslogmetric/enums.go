package schemalesslogmetric

type Measure string

var Measures = struct {
	Attribute  Measure
	Occurrence Measure
}{
	Measure("ATTRIBUTE"),
	Measure("OCCURRENCE"),
}
