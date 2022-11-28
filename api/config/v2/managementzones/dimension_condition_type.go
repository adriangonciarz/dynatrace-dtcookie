package managementzones

type DimensionConditionType string

var DimensionConditionTypes = struct {
	Dimension   DimensionConditionType
	LogFileName DimensionConditionType
	MetricKey   DimensionConditionType
}{
	DimensionConditionType("DIMENSION"),
	DimensionConditionType("LOG_FILE_NAME"),
	DimensionConditionType("METRIC_KEY"),
}
