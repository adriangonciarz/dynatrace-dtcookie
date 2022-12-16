package requesterrors

type UrlFilter string

var UrlFilters = struct {
	BeginsWith UrlFilter
	EndsWith   UrlFilter
	Contains   UrlFilter
	Equals     UrlFilter
}{
	UrlFilter("BEGINS_WITH"),
	UrlFilter("ENDS_WITH"),
	UrlFilter("CONTAINS"),
	UrlFilter("EQUALS"),
}
