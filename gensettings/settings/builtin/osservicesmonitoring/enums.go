package osservicesmonitoring

type LinuxServiceProp string

var LinuxServiceProps = struct {
	Servicename LinuxServiceProp
	Startuptype LinuxServiceProp
}{
	LinuxServiceProp("ServiceName"),
	LinuxServiceProp("StartupType"),
}

type System string

var Systems = struct {
	Linux   System
	Windows System
}{
	System("LINUX"),
	System("WINDOWS"),
}

type WindowsServiceProps string

var WindowsServicePropss = struct {
	Displayname  WindowsServiceProps
	Manufacturer WindowsServiceProps
	Path         WindowsServiceProps
	Servicename  WindowsServiceProps
	Startuptype  WindowsServiceProps
}{
	WindowsServiceProps("DisplayName"),
	WindowsServiceProps("Manufacturer"),
	WindowsServiceProps("Path"),
	WindowsServiceProps("ServiceName"),
	WindowsServiceProps("StartupType"),
}
