package providerbreakdown

type ResourceType string

var ResourceTypes = struct {
	Cdn        ResourceType
	Firstparty ResourceType
	Thirdparty ResourceType
}{
	ResourceType("Cdn"),
	ResourceType("FirstParty"),
	ResourceType("ThirdParty"),
}
