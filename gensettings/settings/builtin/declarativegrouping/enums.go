package declarativegrouping

type ProcessItem string

var ProcessItems = struct {
	Executablepath ProcessItem
	Commandline    ProcessItem
	Executable     ProcessItem
}{
	ProcessItem("ExecutablePath"),
	ProcessItem("CommandLine"),
	ProcessItem("Executable"),
}
