package rumjavascriptupdates

type JavascriptVersion string

var JavascriptVersions = struct {
	Custom               JavascriptVersion
	LatestIe710Supported JavascriptVersion
	LatestStable         JavascriptVersion
	PreviousStable       JavascriptVersion
}{
	JavascriptVersion("CUSTOM"),
	JavascriptVersion("LATEST_IE7_10_SUPPORTED"),
	JavascriptVersion("LATEST_STABLE"),
	JavascriptVersion("PREVIOUS_STABLE"),
}
