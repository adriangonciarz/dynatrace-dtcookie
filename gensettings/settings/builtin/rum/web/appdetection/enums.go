package appdetection

type Matcher string

var Matchers = struct {
	DomainContains   Matcher
	DomainEndsWith   Matcher
	DomainEquals     Matcher
	DomainMatches    Matcher
	DomainStartsWith Matcher
	UrlContains      Matcher
	UrlEndsWith      Matcher
	UrlEquals        Matcher
	UrlStartsWith    Matcher
}{
	Matcher("DOMAIN_CONTAINS"),
	Matcher("DOMAIN_ENDS_WITH"),
	Matcher("DOMAIN_EQUALS"),
	Matcher("DOMAIN_MATCHES"),
	Matcher("DOMAIN_STARTS_WITH"),
	Matcher("URL_CONTAINS"),
	Matcher("URL_ENDS_WITH"),
	Matcher("URL_EQUALS"),
	Matcher("URL_STARTS_WITH"),
}
