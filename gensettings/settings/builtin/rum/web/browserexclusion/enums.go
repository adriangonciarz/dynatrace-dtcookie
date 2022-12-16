package browserexclusion

type Browser string

var Browsers = struct {
	AndroidWebkit    Browser
	BotsAndSpiders   Browser
	Chrome           Browser
	Edge             Browser
	Firefox          Browser
	InternetExplorer Browser
	Opera            Browser
	Safari           Browser
}{
	Browser("ANDROID_WEBKIT"),
	Browser("BOTS_AND_SPIDERS"),
	Browser("CHROME"),
	Browser("EDGE"),
	Browser("FIREFOX"),
	Browser("INTERNET_EXPLORER"),
	Browser("OPERA"),
	Browser("SAFARI"),
}

type Platform string

var Platforms = struct {
	All     Platform
	Desktop Platform
	Mobile  Platform
}{
	Platform("ALL"),
	Platform("DESKTOP"),
	Platform("MOBILE"),
}

type VersionComparator string

var VersionComparators = struct {
	Equals         VersionComparator
	GreaterOrEqual VersionComparator
	LessOrEqual    VersionComparator
}{
	VersionComparator("EQUALS"),
	VersionComparator("GREATER_OR_EQUAL"),
	VersionComparator("LESS_OR_EQUAL"),
}
