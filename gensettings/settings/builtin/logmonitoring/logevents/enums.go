package logevents

type EventTypeEnum string

var EventTypeEnums = struct {
	Availability         EventTypeEnum
	CustomAlert          EventTypeEnum
	CustomAnnotation     EventTypeEnum
	CustomConfiguration  EventTypeEnum
	CustomDeployment     EventTypeEnum
	Error                EventTypeEnum
	Info                 EventTypeEnum
	MarkedForTermination EventTypeEnum
	Resource             EventTypeEnum
	Slowdown             EventTypeEnum
}{
	EventTypeEnum("AVAILABILITY"),
	EventTypeEnum("CUSTOM_ALERT"),
	EventTypeEnum("CUSTOM_ANNOTATION"),
	EventTypeEnum("CUSTOM_CONFIGURATION"),
	EventTypeEnum("CUSTOM_DEPLOYMENT"),
	EventTypeEnum("ERROR"),
	EventTypeEnum("INFO"),
	EventTypeEnum("MARKED_FOR_TERMINATION"),
	EventTypeEnum("RESOURCE"),
	EventTypeEnum("SLOWDOWN"),
}
