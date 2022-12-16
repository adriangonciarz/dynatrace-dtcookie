package monitoringstate

type ProcessGroupMonitoringMode string

var ProcessGroupMonitoringModes = struct {
	Default       ProcessGroupMonitoringMode
	MonitoringOff ProcessGroupMonitoringMode
	MonitoringOn  ProcessGroupMonitoringMode
}{
	ProcessGroupMonitoringMode("DEFAULT"),
	ProcessGroupMonitoringMode("MONITORING_OFF"),
	ProcessGroupMonitoringMode("MONITORING_ON"),
}
