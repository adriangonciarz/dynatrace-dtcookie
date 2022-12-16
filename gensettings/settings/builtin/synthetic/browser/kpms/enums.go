package kpms

type LoadKpm string

var LoadKpms = struct {
	VisuallyComplete       LoadKpm
	ResponseEnd            LoadKpm
	LargestContentfulPaint LoadKpm
	SpeedIndex             LoadKpm
	LoadEventStart         LoadKpm
	ResponseStart          LoadKpm
	CumulativeLayoutShift  LoadKpm
	DomInteractive         LoadKpm
	LoadEventEnd           LoadKpm
	UserActionDuration     LoadKpm
}{
	LoadKpm("VISUALLY_COMPLETE"),
	LoadKpm("RESPONSE_END"),
	LoadKpm("LARGEST_CONTENTFUL_PAINT"),
	LoadKpm("SPEED_INDEX"),
	LoadKpm("LOAD_EVENT_START"),
	LoadKpm("RESPONSE_START"),
	LoadKpm("CUMULATIVE_LAYOUT_SHIFT"),
	LoadKpm("DOM_INTERACTIVE"),
	LoadKpm("LOAD_EVENT_END"),
	LoadKpm("USER_ACTION_DURATION"),
}

type XhrKpm string

var XhrKpms = struct {
	VisuallyComplete   XhrKpm
	ResponseStart      XhrKpm
	ResponseEnd        XhrKpm
	UserActionDuration XhrKpm
}{
	XhrKpm("VISUALLY_COMPLETE"),
	XhrKpm("RESPONSE_START"),
	XhrKpm("RESPONSE_END"),
	XhrKpm("USER_ACTION_DURATION"),
}
