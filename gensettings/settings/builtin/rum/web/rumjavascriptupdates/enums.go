package rumjavascriptupdates

type JavascriptVersion string

var JavascriptVersions = struct {
	LatestIe710Supported JavascriptVersion
	Custom               JavascriptVersion
	LatestStable         JavascriptVersion
	PreviousStable       JavascriptVersion
}{
	JavascriptVersion("LATEST_IE7_10_SUPPORTED"),
	JavascriptVersion("CUSTOM"),
	JavascriptVersion("LATEST_STABLE"),
	JavascriptVersion("PREVIOUS_STABLE"),
}
