package privacy

type IpAddressMaskingOption string

var IpAddressMaskingOptions = struct {
	All    IpAddressMaskingOption
	Public IpAddressMaskingOption
}{
	IpAddressMaskingOption("All"),
	IpAddressMaskingOption("Public"),
}

type DoNotTrackOption string

var DoNotTrackOptions = struct {
	DisableRum DoNotTrackOption
	Anonymous  DoNotTrackOption
}{
	DoNotTrackOption("Disable_rum"),
	DoNotTrackOption("Anonymous"),
}
