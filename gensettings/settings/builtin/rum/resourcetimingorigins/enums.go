package resourcetimingorigins

type OriginMatcherType string

var OriginMatcherTypes = struct {
	EndsWith   OriginMatcherType
	Contains   OriginMatcherType
	Equals     OriginMatcherType
	StartsWith OriginMatcherType
}{
	OriginMatcherType("ENDS_WITH"),
	OriginMatcherType("CONTAINS"),
	OriginMatcherType("EQUALS"),
	OriginMatcherType("STARTS_WITH"),
}
