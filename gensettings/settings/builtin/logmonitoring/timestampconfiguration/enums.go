package timestampconfiguration

type MatcherType string

var MatcherTypes = struct {
	ContainerName          MatcherType
	DtEntityContainerGroup MatcherType
	DtEntityProcessGroup   MatcherType
	K8SContainerName       MatcherType
	K8SDeploymentName      MatcherType
	K8SNamespaceName       MatcherType
	LogSource              MatcherType
	ProcessTechnology      MatcherType
}{
	MatcherType("Container_name"),
	MatcherType("Dt_entity_container_group"),
	MatcherType("Dt_entity_process_group"),
	MatcherType("K8s_container_name"),
	MatcherType("K8s_deployment_name"),
	MatcherType("K8s_namespace_name"),
	MatcherType("Log_source"),
	MatcherType("Process_technology"),
}

type Operator string

var Operators = struct {
	Matches Operator
}{
	Operator("MATCHES"),
}
