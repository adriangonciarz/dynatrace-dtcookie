package cloudapplicationworkloaddetection

type MatchEnum string

var MatchEnums = struct {
	Contains    MatchEnum
	Ends        MatchEnum
	Equals      MatchEnum
	Exists      MatchEnum
	NotContains MatchEnum
	NotEnds     MatchEnum
	NotEquals   MatchEnum
	NotStarts   MatchEnum
	Starts      MatchEnum
}{
	MatchEnum("CONTAINS"),
	MatchEnum("ENDS"),
	MatchEnum("EQUALS"),
	MatchEnum("EXISTS"),
	MatchEnum("NOT_CONTAINS"),
	MatchEnum("NOT_ENDS"),
	MatchEnum("NOT_EQUALS"),
	MatchEnum("NOT_STARTS"),
	MatchEnum("STARTS"),
}
