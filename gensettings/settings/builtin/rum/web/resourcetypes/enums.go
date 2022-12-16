package resourcetypes

type PrimaryResourceType string

var PrimaryResourceTypes = struct {
	Css    PrimaryResourceType
	Image  PrimaryResourceType
	Other  PrimaryResourceType
	Script PrimaryResourceType
}{
	PrimaryResourceType("CSS"),
	PrimaryResourceType("IMAGE"),
	PrimaryResourceType("OTHER"),
	PrimaryResourceType("SCRIPT"),
}
