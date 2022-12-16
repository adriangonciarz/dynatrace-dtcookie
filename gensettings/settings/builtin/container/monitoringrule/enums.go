package monitoringrule

type ContainerItem string

var ContainerItems = struct {
	KubernetesPoduid        ContainerItem
	ContainerName           ContainerItem
	ImageName               ContainerItem
	KubernetesNamespace     ContainerItem
	KubernetesContainername ContainerItem
	KubernetesBasepodname   ContainerItem
	KubernetesFullpodname   ContainerItem
}{
	ContainerItem("KUBERNETES_PODUID"),
	ContainerItem("CONTAINER_NAME"),
	ContainerItem("IMAGE_NAME"),
	ContainerItem("KUBERNETES_NAMESPACE"),
	ContainerItem("KUBERNETES_CONTAINERNAME"),
	ContainerItem("KUBERNETES_BASEPODNAME"),
	ContainerItem("KUBERNETES_FULLPODNAME"),
}

type MonitoringMode string

var MonitoringModes = struct {
	MonitoringOff MonitoringMode
	MonitoringOn  MonitoringMode
}{
	MonitoringMode("MONITORING_OFF"),
	MonitoringMode("MONITORING_ON"),
}

type ConditionOperator string

var ConditionOperators = struct {
	NotContains ConditionOperator
	Equals      ConditionOperator
	NotExists   ConditionOperator
	Starts      ConditionOperator
	Ends        ConditionOperator
	Contains    ConditionOperator
	NotEquals   ConditionOperator
	Exists      ConditionOperator
	NotStarts   ConditionOperator
	NotEnds     ConditionOperator
}{
	ConditionOperator("NOT_CONTAINS"),
	ConditionOperator("EQUALS"),
	ConditionOperator("NOT_EXISTS"),
	ConditionOperator("STARTS"),
	ConditionOperator("ENDS"),
	ConditionOperator("CONTAINS"),
	ConditionOperator("NOT_EQUALS"),
	ConditionOperator("EXISTS"),
	ConditionOperator("NOT_STARTS"),
	ConditionOperator("NOT_ENDS"),
}
