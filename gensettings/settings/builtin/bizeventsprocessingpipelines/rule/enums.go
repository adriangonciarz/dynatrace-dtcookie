package rule

type TransformationFieldType string

var TransformationFieldTypes = struct {
	Boolean   TransformationFieldType
	Double    TransformationFieldType
	Duration  TransformationFieldType
	Int       TransformationFieldType
	Ipaddr    TransformationFieldType
	Long      TransformationFieldType
	String    TransformationFieldType
	Timestamp TransformationFieldType
}{
	TransformationFieldType("BOOLEAN"),
	TransformationFieldType("DOUBLE"),
	TransformationFieldType("DURATION"),
	TransformationFieldType("INT"),
	TransformationFieldType("IPADDR"),
	TransformationFieldType("LONG"),
	TransformationFieldType("STRING"),
	TransformationFieldType("TIMESTAMP"),
}
