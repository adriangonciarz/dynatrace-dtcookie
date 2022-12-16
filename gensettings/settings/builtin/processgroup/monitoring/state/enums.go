package state

type ProcessGroupMonitoringMode string

var ProcessGroupMonitoringModes = struct {
	MonitoringOff ProcessGroupMonitoringMode
	MonitoringOn  ProcessGroupMonitoringMode
	Default       ProcessGroupMonitoringMode
}{
	ProcessGroupMonitoringMode("MONITORING_OFF"),
	ProcessGroupMonitoringMode("MONITORING_ON"),
	ProcessGroupMonitoringMode("DEFAULT"),
}
