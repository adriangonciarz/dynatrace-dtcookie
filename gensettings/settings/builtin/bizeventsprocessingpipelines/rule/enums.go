package rule

type TransformationFieldType string

var TransformationFieldTypes = struct {
	Long      TransformationFieldType
	Double    TransformationFieldType
	Duration  TransformationFieldType
	Timestamp TransformationFieldType
	Ipaddr    TransformationFieldType
	String    TransformationFieldType
	Boolean   TransformationFieldType
	Int       TransformationFieldType
}{
	TransformationFieldType("LONG"),
	TransformationFieldType("DOUBLE"),
	TransformationFieldType("DURATION"),
	TransformationFieldType("TIMESTAMP"),
	TransformationFieldType("IPADDR"),
	TransformationFieldType("STRING"),
	TransformationFieldType("BOOLEAN"),
	TransformationFieldType("INT"),
}
