package processavailability

type ProcessItem string

var ProcessItems = struct {
	Executable     ProcessItem
	Executablepath ProcessItem
	Commandline    ProcessItem
}{
	ProcessItem("Executable"),
	ProcessItem("ExecutablePath"),
	ProcessItem("CommandLine"),
}
