package browserexclusion

type Browser string

var Browsers = struct {
	Firefox          Browser
	InternetExplorer Browser
	Opera            Browser
	Safari           Browser
	Edge             Browser
	BotsAndSpiders   Browser
	AndroidWebkit    Browser
	Chrome           Browser
}{
	Browser("FIREFOX"),
	Browser("INTERNET_EXPLORER"),
	Browser("OPERA"),
	Browser("SAFARI"),
	Browser("EDGE"),
	Browser("BOTS_AND_SPIDERS"),
	Browser("ANDROID_WEBKIT"),
	Browser("CHROME"),
}

type VersionComparator string

var VersionComparators = struct {
	Equals         VersionComparator
	LessOrEqual    VersionComparator
	GreaterOrEqual VersionComparator
}{
	VersionComparator("EQUALS"),
	VersionComparator("LESS_OR_EQUAL"),
	VersionComparator("GREATER_OR_EQUAL"),
}

type Platform string

var Platforms = struct {
	All     Platform
	Mobile  Platform
	Desktop Platform
}{
	Platform("ALL"),
	Platform("MOBILE"),
	Platform("DESKTOP"),
}
