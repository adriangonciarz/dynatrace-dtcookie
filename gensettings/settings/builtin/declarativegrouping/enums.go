package declarativegrouping

type ProcessItem string

var ProcessItems = struct {
	Commandline    ProcessItem
	Executable     ProcessItem
	Executablepath ProcessItem
}{
	ProcessItem("CommandLine"),
	ProcessItem("Executable"),
	ProcessItem("ExecutablePath"),
}
