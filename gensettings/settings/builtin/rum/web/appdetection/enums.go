package appdetection

type Matcher string

var Matchers = struct {
	UrlContains      Matcher
	UrlEndsWith      Matcher
	DomainEquals     Matcher
	DomainMatches    Matcher
	DomainStartsWith Matcher
	UrlStartsWith    Matcher
	DomainEndsWith   Matcher
	UrlEquals        Matcher
	DomainContains   Matcher
}{
	Matcher("URL_CONTAINS"),
	Matcher("URL_ENDS_WITH"),
	Matcher("DOMAIN_EQUALS"),
	Matcher("DOMAIN_MATCHES"),
	Matcher("DOMAIN_STARTS_WITH"),
	Matcher("URL_STARTS_WITH"),
	Matcher("DOMAIN_ENDS_WITH"),
	Matcher("URL_EQUALS"),
	Matcher("DOMAIN_CONTAINS"),
}
