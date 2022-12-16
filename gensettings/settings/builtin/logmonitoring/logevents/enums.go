package logevents

type EventTypeEnum string

var EventTypeEnums = struct {
	Resource             EventTypeEnum
	CustomAnnotation     EventTypeEnum
	Availability         EventTypeEnum
	CustomConfiguration  EventTypeEnum
	MarkedForTermination EventTypeEnum
	Error                EventTypeEnum
	Slowdown             EventTypeEnum
	CustomDeployment     EventTypeEnum
	Info                 EventTypeEnum
	CustomAlert          EventTypeEnum
}{
	EventTypeEnum("RESOURCE"),
	EventTypeEnum("CUSTOM_ANNOTATION"),
	EventTypeEnum("AVAILABILITY"),
	EventTypeEnum("CUSTOM_CONFIGURATION"),
	EventTypeEnum("MARKED_FOR_TERMINATION"),
	EventTypeEnum("ERROR"),
	EventTypeEnum("SLOWDOWN"),
	EventTypeEnum("CUSTOM_DEPLOYMENT"),
	EventTypeEnum("INFO"),
	EventTypeEnum("CUSTOM_ALERT"),
}
