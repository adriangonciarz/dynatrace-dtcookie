package resourcetypes

type PrimaryResourceType string

var PrimaryResourceTypes = struct {
	Script PrimaryResourceType
	Other  PrimaryResourceType
	Css    PrimaryResourceType
	Image  PrimaryResourceType
}{
	PrimaryResourceType("SCRIPT"),
	PrimaryResourceType("OTHER"),
	PrimaryResourceType("CSS"),
	PrimaryResourceType("IMAGE"),
}
