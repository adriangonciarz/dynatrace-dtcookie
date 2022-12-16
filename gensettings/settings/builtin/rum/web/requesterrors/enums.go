package requesterrors

type UrlFilter string

var UrlFilters = struct {
	BeginsWith UrlFilter
	Contains   UrlFilter
	EndsWith   UrlFilter
	Equals     UrlFilter
}{
	UrlFilter("BEGINS_WITH"),
	UrlFilter("CONTAINS"),
	UrlFilter("ENDS_WITH"),
	UrlFilter("EQUALS"),
}
