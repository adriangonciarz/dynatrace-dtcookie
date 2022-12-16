package osservicesmonitoring

type LinuxServiceProp string

var LinuxServiceProps = struct {
	Servicename LinuxServiceProp
	Startuptype LinuxServiceProp
}{
	LinuxServiceProp("ServiceName"),
	LinuxServiceProp("StartupType"),
}

type WindowsServiceProps string

var WindowsServicePropss = struct {
	Path         WindowsServiceProps
	Startuptype  WindowsServiceProps
	Manufacturer WindowsServiceProps
	Displayname  WindowsServiceProps
	Servicename  WindowsServiceProps
}{
	WindowsServiceProps("Path"),
	WindowsServiceProps("StartupType"),
	WindowsServiceProps("Manufacturer"),
	WindowsServiceProps("DisplayName"),
	WindowsServiceProps("ServiceName"),
}

type System string

var Systems = struct {
	Windows System
	Linux   System
}{
	System("WINDOWS"),
	System("LINUX"),
}
