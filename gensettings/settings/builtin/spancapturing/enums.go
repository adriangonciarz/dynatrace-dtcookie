package spancapturing

type SpanCaptureAction string

var SpanCaptureActions = struct {
	Capture SpanCaptureAction
	Ignore  SpanCaptureAction
}{
	SpanCaptureAction("CAPTURE"),
	SpanCaptureAction("IGNORE"),
}

type SpanKind string

var SpanKinds = struct {
	Client   SpanKind
	Consumer SpanKind
	Internal SpanKind
	Producer SpanKind
	Server   SpanKind
}{
	SpanKind("CLIENT"),
	SpanKind("CONSUMER"),
	SpanKind("INTERNAL"),
	SpanKind("PRODUCER"),
	SpanKind("SERVER"),
}

type SpanMatcherSource string

var SpanMatcherSources = struct {
	Attribute                   SpanMatcherSource
	InstrumentationScopeName    SpanMatcherSource
	InstrumentationScopeVersion SpanMatcherSource
	SpanKind                    SpanMatcherSource
	SpanName                    SpanMatcherSource
}{
	SpanMatcherSource("ATTRIBUTE"),
	SpanMatcherSource("INSTRUMENTATION_SCOPE_NAME"),
	SpanMatcherSource("INSTRUMENTATION_SCOPE_VERSION"),
	SpanMatcherSource("SPAN_KIND"),
	SpanMatcherSource("SPAN_NAME"),
}

type SpanMatcherType string

var SpanMatcherTypes = struct {
	Contains         SpanMatcherType
	DoesNotContain   SpanMatcherType
	DoesNotEndWith   SpanMatcherType
	DoesNotEqual     SpanMatcherType
	DoesNotStartWith SpanMatcherType
	EndsWith         SpanMatcherType
	Equals           SpanMatcherType
	StartsWith       SpanMatcherType
}{
	SpanMatcherType("CONTAINS"),
	SpanMatcherType("DOES_NOT_CONTAIN"),
	SpanMatcherType("DOES_NOT_END_WITH"),
	SpanMatcherType("DOES_NOT_EQUAL"),
	SpanMatcherType("DOES_NOT_START_WITH"),
	SpanMatcherType("ENDS_WITH"),
	SpanMatcherType("EQUALS"),
	SpanMatcherType("STARTS_WITH"),
}
