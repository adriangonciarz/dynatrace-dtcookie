package beacondomainorigins

type OriginMatcherType string

var OriginMatcherTypes = struct {
	Equals     OriginMatcherType
	StartsWith OriginMatcherType
	EndsWith   OriginMatcherType
	Contains   OriginMatcherType
}{
	OriginMatcherType("EQUALS"),
	OriginMatcherType("STARTS_WITH"),
	OriginMatcherType("ENDS_WITH"),
	OriginMatcherType("CONTAINS"),
}
