package providerbreakdown

type ResourceType string

var ResourceTypes = struct {
	Firstparty ResourceType
	Thirdparty ResourceType
	Cdn        ResourceType
}{
	ResourceType("FirstParty"),
	ResourceType("ThirdParty"),
	ResourceType("Cdn"),
}
