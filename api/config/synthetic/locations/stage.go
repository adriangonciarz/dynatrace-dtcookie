package locations

type Stage string

var Stages = struct {
	Beta       Stage
	ComingSoon Stage
	Deleted    Stage
	GA         Stage
}{
	Stage("BETA"),
	Stage("COMING_SOON"),
	Stage("DELETED"),
	Stage("GA"),
}
