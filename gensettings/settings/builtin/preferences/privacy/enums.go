package privacy

type DoNotTrackOption string

var DoNotTrackOptions = struct {
	Anonymous  DoNotTrackOption
	DisableRum DoNotTrackOption
}{
	DoNotTrackOption("Anonymous"),
	DoNotTrackOption("Disable_rum"),
}

type IpAddressMaskingOption string

var IpAddressMaskingOptions = struct {
	All    IpAddressMaskingOption
	Public IpAddressMaskingOption
}{
	IpAddressMaskingOption("All"),
	IpAddressMaskingOption("Public"),
}
