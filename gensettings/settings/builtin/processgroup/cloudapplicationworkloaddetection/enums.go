package cloudapplicationworkloaddetection

type MatchEnum string

var MatchEnums = struct {
	Starts      MatchEnum
	NotEquals   MatchEnum
	Contains    MatchEnum
	Exists      MatchEnum
	NotContains MatchEnum
	NotEnds     MatchEnum
	Equals      MatchEnum
	Ends        MatchEnum
	NotStarts   MatchEnum
}{
	MatchEnum("STARTS"),
	MatchEnum("NOT_EQUALS"),
	MatchEnum("CONTAINS"),
	MatchEnum("EXISTS"),
	MatchEnum("NOT_CONTAINS"),
	MatchEnum("NOT_ENDS"),
	MatchEnum("EQUALS"),
	MatchEnum("ENDS"),
	MatchEnum("NOT_STARTS"),
}
