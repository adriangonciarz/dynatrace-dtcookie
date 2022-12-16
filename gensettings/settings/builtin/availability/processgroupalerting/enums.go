package processgroupalerting

type AlertingMode string

var AlertingModes = struct {
	OnInstanceCountViolation AlertingMode
	OnPgiUnavailability      AlertingMode
}{
	AlertingMode("ON_INSTANCE_COUNT_VIOLATION"),
	AlertingMode("ON_PGI_UNAVAILABILITY"),
}
