package timestampconfiguration

type MatcherType string

var MatcherTypes = struct {
	DtEntityProcessGroup   MatcherType
	LogSource              MatcherType
	K8SContainerName       MatcherType
	K8SNamespaceName       MatcherType
	K8SDeploymentName      MatcherType
	ContainerName          MatcherType
	DtEntityContainerGroup MatcherType
	ProcessTechnology      MatcherType
}{
	MatcherType("Dt_entity_process_group"),
	MatcherType("Log_source"),
	MatcherType("K8s_container_name"),
	MatcherType("K8s_namespace_name"),
	MatcherType("K8s_deployment_name"),
	MatcherType("Container_name"),
	MatcherType("Dt_entity_container_group"),
	MatcherType("Process_technology"),
}

type Operator string

var Operators = struct {
	Matches Operator
}{
	Operator("MATCHES"),
}
