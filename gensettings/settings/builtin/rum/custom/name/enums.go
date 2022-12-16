package name

type ApplicationType string

var ApplicationTypes = struct {
	Desktop    ApplicationType
	Echo       ApplicationType
	EmbeddedPc ApplicationType
	Hololens   ApplicationType
	Iot        ApplicationType
	Ufo        ApplicationType
}{
	ApplicationType("Desktop"),
	ApplicationType("Echo"),
	ApplicationType("Embedded_pc"),
	ApplicationType("Hololens"),
	ApplicationType("Iot"),
	ApplicationType("Ufo"),
}
