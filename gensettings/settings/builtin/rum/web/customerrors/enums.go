package customerrors

type Matcher string

var Matchers = struct {
	All        Matcher
	BeginsWith Matcher
	Contains   Matcher
	EndsWith   Matcher
	Equals     Matcher
}{
	Matcher("ALL"),
	Matcher("BEGINS_WITH"),
	Matcher("CONTAINS"),
	Matcher("ENDS_WITH"),
	Matcher("EQUALS"),
}
