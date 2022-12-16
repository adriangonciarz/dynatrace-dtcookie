package metricevents

type EventTypeEnum string

var EventTypeEnums = struct {
	CustomAlert          EventTypeEnum
	CustomDeployment     EventTypeEnum
	Error                EventTypeEnum
	Slowdown             EventTypeEnum
	Info                 EventTypeEnum
	MarkedForTermination EventTypeEnum
	Availability         EventTypeEnum
	Resource             EventTypeEnum
	CustomAnnotation     EventTypeEnum
	CustomConfiguration  EventTypeEnum
}{
	EventTypeEnum("CUSTOM_ALERT"),
	EventTypeEnum("CUSTOM_DEPLOYMENT"),
	EventTypeEnum("ERROR"),
	EventTypeEnum("SLOWDOWN"),
	EventTypeEnum("INFO"),
	EventTypeEnum("MARKED_FOR_TERMINATION"),
	EventTypeEnum("AVAILABILITY"),
	EventTypeEnum("RESOURCE"),
	EventTypeEnum("CUSTOM_ANNOTATION"),
	EventTypeEnum("CUSTOM_CONFIGURATION"),
}

type Aggregation string

var Aggregations = struct {
	Sum          Aggregation
	Count        Aggregation
	Avg          Aggregation
	Value        Aggregation
	Median       Aggregation
	Percentile90 Aggregation
	Min          Aggregation
	Max          Aggregation
}{
	Aggregation("SUM"),
	Aggregation("COUNT"),
	Aggregation("AVG"),
	Aggregation("VALUE"),
	Aggregation("MEDIAN"),
	Aggregation("PERCENTILE90"),
	Aggregation("MIN"),
	Aggregation("MAX"),
}

type AlertCondition string

var AlertConditions = struct {
	Above   AlertCondition
	Below   AlertCondition
	Outside AlertCondition
}{
	AlertCondition("ABOVE"),
	AlertCondition("BELOW"),
	AlertCondition("OUTSIDE"),
}

type EntityFilterOperator string

var EntityFilterOperators = struct {
	Equals                  EntityFilterOperator
	ContainsCaseInsensitive EntityFilterOperator
	ContainsCaseSensitive   EntityFilterOperator
}{
	EntityFilterOperator("EQUALS"),
	EntityFilterOperator("CONTAINS_CASE_INSENSITIVE"),
	EntityFilterOperator("CONTAINS_CASE_SENSITIVE"),
}

type ModelType string

var ModelTypes = struct {
	StaticThreshold       ModelType
	AutoAdaptiveThreshold ModelType
	SeasonalBaseline      ModelType
}{
	ModelType("STATIC_THRESHOLD"),
	ModelType("AUTO_ADAPTIVE_THRESHOLD"),
	ModelType("SEASONAL_BASELINE"),
}

type Type string

var Types = struct {
	MetricKey      Type
	MetricSelector Type
}{
	Type("METRIC_KEY"),
	Type("METRIC_SELECTOR"),
}

type EntityFilterType string

var EntityFilterTypes = struct {
	EntityId              EntityFilterType
	ProcessGroupId        EntityFilterType
	HostName              EntityFilterType
	ProcessGroupName      EntityFilterType
	CustomDeviceGroupName EntityFilterType
	Name                  EntityFilterType
	ManagementZone        EntityFilterType
	HostGroupName         EntityFilterType
	Tag                   EntityFilterType
}{
	EntityFilterType("ENTITY_ID"),
	EntityFilterType("PROCESS_GROUP_ID"),
	EntityFilterType("HOST_NAME"),
	EntityFilterType("PROCESS_GROUP_NAME"),
	EntityFilterType("CUSTOM_DEVICE_GROUP_NAME"),
	EntityFilterType("NAME"),
	EntityFilterType("MANAGEMENT_ZONE"),
	EntityFilterType("HOST_GROUP_NAME"),
	EntityFilterType("TAG"),
}
