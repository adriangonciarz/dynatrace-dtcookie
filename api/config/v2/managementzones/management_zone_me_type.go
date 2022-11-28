package managementzones

type ManagementZoneMeType string

var ManagementZoneMeTypes = struct {
	Azure                        ManagementZoneMeType
	DataCenterService            ManagementZoneMeType
	Service                      ManagementZoneMeType
	MobileApplication            ManagementZoneMeType
	CloudFoundryFoundation       ManagementZoneMeType
	OpenstackAccount             ManagementZoneMeType
	AppmonSystemProfile          ManagementZoneMeType
	EsxiHost                     ManagementZoneMeType
	CustomDevice                 ManagementZoneMeType
	CloudApplication             ManagementZoneMeType
	AwsRelationalDatabaseService ManagementZoneMeType
	ExternalMonitor              ManagementZoneMeType
	EnterpriseApplication        ManagementZoneMeType
	HttpMonitor                  ManagementZoneMeType
	ProcessGroup                 ManagementZoneMeType
	Host                         ManagementZoneMeType
	Queue                        ManagementZoneMeType
	AwsClassicLoadBalancer       ManagementZoneMeType
	AwsApplicationLoadBalancer   ManagementZoneMeType
	CustomApplication            ManagementZoneMeType
	AwsNetworkLoadBalancer       ManagementZoneMeType
	BrowserMonitor               ManagementZoneMeType
	AppmonServer                 ManagementZoneMeType
	CustomDeviceGroup            ManagementZoneMeType
	KubernetesCluster            ManagementZoneMeType
	HostGroup                    ManagementZoneMeType
	WebApplication               ManagementZoneMeType
	AwsAccount                   ManagementZoneMeType
	KubernetesService            ManagementZoneMeType
	AwsAutoScalingGroup          ManagementZoneMeType
	CloudApplicationNamespace    ManagementZoneMeType
}{
	ManagementZoneMeType("AZURE"),
	ManagementZoneMeType("DATA_CENTER_SERVICE"),
	ManagementZoneMeType("SERVICE"),
	ManagementZoneMeType("MOBILE_APPLICATION"),
	ManagementZoneMeType("CLOUD_FOUNDRY_FOUNDATION"),
	ManagementZoneMeType("OPENSTACK_ACCOUNT"),
	ManagementZoneMeType("APPMON_SYSTEM_PROFILE"),
	ManagementZoneMeType("ESXI_HOST"),
	ManagementZoneMeType("CUSTOM_DEVICE"),
	ManagementZoneMeType("CLOUD_APPLICATION"),
	ManagementZoneMeType("AWS_RELATIONAL_DATABASE_SERVICE"),
	ManagementZoneMeType("EXTERNAL_MONITOR"),
	ManagementZoneMeType("ENTERPRISE_APPLICATION"),
	ManagementZoneMeType("HTTP_MONITOR"),
	ManagementZoneMeType("PROCESS_GROUP"),
	ManagementZoneMeType("HOST"),
	ManagementZoneMeType("QUEUE"),
	ManagementZoneMeType("AWS_CLASSIC_LOAD_BALANCER"),
	ManagementZoneMeType("AWS_APPLICATION_LOAD_BALANCER"),
	ManagementZoneMeType("CUSTOM_APPLICATION"),
	ManagementZoneMeType("AWS_NETWORK_LOAD_BALANCER"),
	ManagementZoneMeType("BROWSER_MONITOR"),
	ManagementZoneMeType("APPMON_SERVER"),
	ManagementZoneMeType("CUSTOM_DEVICE_GROUP"),
	ManagementZoneMeType("KUBERNETES_CLUSTER"),
	ManagementZoneMeType("HOST_GROUP"),
	ManagementZoneMeType("WEB_APPLICATION"),
	ManagementZoneMeType("AWS_ACCOUNT"),
	ManagementZoneMeType("KUBERNETES_SERVICE"),
	ManagementZoneMeType("AWS_AUTO_SCALING_GROUP"),
	ManagementZoneMeType("CLOUD_APPLICATION_NAMESPACE"),
}
