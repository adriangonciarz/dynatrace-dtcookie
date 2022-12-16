package customerrors

type Matcher string

var Matchers = struct {
	EndsWith   Matcher
	Contains   Matcher
	Equals     Matcher
	All        Matcher
	BeginsWith Matcher
}{
	Matcher("ENDS_WITH"),
	Matcher("CONTAINS"),
	Matcher("EQUALS"),
	Matcher("ALL"),
	Matcher("BEGINS_WITH"),
}
