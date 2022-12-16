package name

type ApplicationType string

var ApplicationTypes = struct {
	Ufo        ApplicationType
	Desktop    ApplicationType
	Echo       ApplicationType
	Hololens   ApplicationType
	Iot        ApplicationType
	EmbeddedPc ApplicationType
}{
	ApplicationType("Ufo"),
	ApplicationType("Desktop"),
	ApplicationType("Echo"),
	ApplicationType("Hololens"),
	ApplicationType("Iot"),
	ApplicationType("Embedded_pc"),
}
