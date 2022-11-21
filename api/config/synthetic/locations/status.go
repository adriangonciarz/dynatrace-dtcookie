package locations

type Status string

var Statuses = struct {
	Disabled Status
	Enabled  Status
	Hidden   Status
}{
	Status("DISABLED"),
	Status("ENABLED"),
	Status("HIDDEN"),
}
