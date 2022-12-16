package spancontextpropagation

type SpanContextPropagationAction string

var SpanContextPropagationActions = struct {
	Propagate     SpanContextPropagationAction
	DontPropagate SpanContextPropagationAction
}{
	SpanContextPropagationAction("PROPAGATE"),
	SpanContextPropagationAction("DONT_PROPAGATE"),
}

type SpanMatcherSource string

var SpanMatcherSources = struct {
	SpanKind                    SpanMatcherSource
	Attribute                   SpanMatcherSource
	InstrumentationScopeName    SpanMatcherSource
	InstrumentationScopeVersion SpanMatcherSource
	SpanName                    SpanMatcherSource
}{
	SpanMatcherSource("SPAN_KIND"),
	SpanMatcherSource("ATTRIBUTE"),
	SpanMatcherSource("INSTRUMENTATION_SCOPE_NAME"),
	SpanMatcherSource("INSTRUMENTATION_SCOPE_VERSION"),
	SpanMatcherSource("SPAN_NAME"),
}

type SpanMatcherType string

var SpanMatcherTypes = struct {
	DoesNotContain   SpanMatcherType
	DoesNotStartWith SpanMatcherType
	DoesNotEndWith   SpanMatcherType
	Equals           SpanMatcherType
	Contains         SpanMatcherType
	StartsWith       SpanMatcherType
	EndsWith         SpanMatcherType
	DoesNotEqual     SpanMatcherType
}{
	SpanMatcherType("DOES_NOT_CONTAIN"),
	SpanMatcherType("DOES_NOT_START_WITH"),
	SpanMatcherType("DOES_NOT_END_WITH"),
	SpanMatcherType("EQUALS"),
	SpanMatcherType("CONTAINS"),
	SpanMatcherType("STARTS_WITH"),
	SpanMatcherType("ENDS_WITH"),
	SpanMatcherType("DOES_NOT_EQUAL"),
}

type SpanKind string

var SpanKinds = struct {
	Producer SpanKind
	Consumer SpanKind
	Internal SpanKind
	Server   SpanKind
	Client   SpanKind
}{
	SpanKind("PRODUCER"),
	SpanKind("CONSUMER"),
	SpanKind("INTERNAL"),
	SpanKind("SERVER"),
	SpanKind("CLIENT"),
}
