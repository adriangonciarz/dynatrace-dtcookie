package spancapturing

type SpanCaptureAction string

var SpanCaptureActions = struct {
	Capture SpanCaptureAction
	Ignore  SpanCaptureAction
}{
	SpanCaptureAction("CAPTURE"),
	SpanCaptureAction("IGNORE"),
}

type SpanMatcherSource string

var SpanMatcherSources = struct {
	InstrumentationScopeVersion SpanMatcherSource
	SpanName                    SpanMatcherSource
	SpanKind                    SpanMatcherSource
	Attribute                   SpanMatcherSource
	InstrumentationScopeName    SpanMatcherSource
}{
	SpanMatcherSource("INSTRUMENTATION_SCOPE_VERSION"),
	SpanMatcherSource("SPAN_NAME"),
	SpanMatcherSource("SPAN_KIND"),
	SpanMatcherSource("ATTRIBUTE"),
	SpanMatcherSource("INSTRUMENTATION_SCOPE_NAME"),
}

type SpanMatcherType string

var SpanMatcherTypes = struct {
	DoesNotStartWith SpanMatcherType
	DoesNotEndWith   SpanMatcherType
	Equals           SpanMatcherType
	Contains         SpanMatcherType
	StartsWith       SpanMatcherType
	EndsWith         SpanMatcherType
	DoesNotEqual     SpanMatcherType
	DoesNotContain   SpanMatcherType
}{
	SpanMatcherType("DOES_NOT_START_WITH"),
	SpanMatcherType("DOES_NOT_END_WITH"),
	SpanMatcherType("EQUALS"),
	SpanMatcherType("CONTAINS"),
	SpanMatcherType("STARTS_WITH"),
	SpanMatcherType("ENDS_WITH"),
	SpanMatcherType("DOES_NOT_EQUAL"),
	SpanMatcherType("DOES_NOT_CONTAIN"),
}

type SpanKind string

var SpanKinds = struct {
	Internal SpanKind
	Server   SpanKind
	Client   SpanKind
	Producer SpanKind
	Consumer SpanKind
}{
	SpanKind("INTERNAL"),
	SpanKind("SERVER"),
	SpanKind("CLIENT"),
	SpanKind("PRODUCER"),
	SpanKind("CONSUMER"),
}
