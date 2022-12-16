package processmonitoring

type MonitoringMode string

var MonitoringModes = struct {
	MonitoringOff MonitoringMode
	MonitoringOn  MonitoringMode
}{
	MonitoringMode("MONITORING_OFF"),
	MonitoringMode("MONITORING_ON"),
}
