package locations

type LocationType string

var LocationTypes = struct {
	Cluster LocationType
	Private LocationType
	Public  LocationType
}{
	LocationType("CLUSTER"),
	LocationType("PRIVATE"),
	LocationType("PUBLIC"),
}
