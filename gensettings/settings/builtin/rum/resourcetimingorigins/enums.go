package resourcetimingorigins

type OriginMatcherType string

var OriginMatcherTypes = struct {
	Contains   OriginMatcherType
	EndsWith   OriginMatcherType
	Equals     OriginMatcherType
	StartsWith OriginMatcherType
}{
	OriginMatcherType("CONTAINS"),
	OriginMatcherType("ENDS_WITH"),
	OriginMatcherType("EQUALS"),
	OriginMatcherType("STARTS_WITH"),
}
