package managementzones

type DimensionType string

var DimensionTypes = struct {
	Metric DimensionType
	Any    DimensionType
	Log    DimensionType
}{
	DimensionType("METRIC"),
	DimensionType("ANY"),
	DimensionType("LOG"),
}

type Attribute string

var Attributes = struct {
	Ec2InstanceId                             Attribute
	OpenstackVmName                           Attribute
	OpenstackVmSecurityGroup                  Attribute
	ExternalMonitorName                       Attribute
	WebApplicationTags                        Attribute
	VmwareDatacenterName                      Attribute
	OpenstackVmInstanceType                   Attribute
	ProcessGroupCustomMetadata                Attribute
	CloudApplicationName                      Attribute
	EsxiHostClusterName                       Attribute
	ExternalMonitorTags                       Attribute
	HostCpuCores                              Attribute
	DataCenterServiceMetadata                 Attribute
	AwsAvailabilityZoneName                   Attribute
	DockerFullImageName                       Attribute
	CustomDeviceIpAddress                     Attribute
	CustomApplicationTags                     Attribute
	CloudFoundryFoundationName                Attribute
	ServiceWebServiceName                     Attribute
	ProcessGroupTechnologyEdition             Attribute
	QueueName                                 Attribute
	HostAixVirtualCpuCount                    Attribute
	HostAixLogicalCpuCount                    Attribute
	AwsRelationalDatabaseServiceEndpoint      Attribute
	OpenstackProjectName                      Attribute
	EsxiHostProductName                       Attribute
	CustomDevicePort                          Attribute
	DockerContainerName                       Attribute
	HttpMonitorTags                           Attribute
	DataCenterServiceTags                     Attribute
	ServicePort                               Attribute
	ServiceDatabaseHostName                   Attribute
	CloudApplicationNamespaceLabels           Attribute
	AwsAccountId                              Attribute
	EsxiHostName                              Attribute
	Ec2InstanceAmiId                          Attribute
	ExternalMonitorEngineType                 Attribute
	HostArchitecture                          Attribute
	WebApplicationNamePattern                 Attribute
	AwsRelationalDatabaseServiceDbName        Attribute
	AzureSubscriptionUuid                     Attribute
	HostOneagentCustomHostName                Attribute
	HostGroupName                             Attribute
	EnterpriseApplicationMetadata             Attribute
	Ec2InstanceTags                           Attribute
	AwsClassicLoadBalancerName                Attribute
	HostKubernetesLabels                      Attribute
	CloudApplicationLabels                    Attribute
	AwsAccountName                            Attribute
	AwsAutoScalingGroupName                   Attribute
	ExternalMonitorEngineName                 Attribute
	HostLogicalCpuCores                       Attribute
	HostTechnology                            Attribute
	OpenstackAccountName                      Attribute
	HostCustomMetadata                        Attribute
	GoogleComputeInstanceProject              Attribute
	OpenstackRegionName                       Attribute
	ProcessGroupAzureHostName                 Attribute
	HttpMonitorName                           Attribute
	HostIpAddress                             Attribute
	EsxiHostHardwareVendor                    Attribute
	HostBoshInstanceId                        Attribute
	AwsRelationalDatabaseServiceInstanceClass Attribute
	GoogleComputeInstanceMachineType          Attribute
	HostBoshDeploymentId                      Attribute
	GoogleComputeInstanceProjectId            Attribute
	OpenstackAccountProjectName               Attribute
	EnterpriseApplicationPort                 Attribute
	Ec2InstanceName                           Attribute
	WebApplicationName                        Attribute
	ServiceTags                               Attribute
	HostTags                                  Attribute
	CloudApplicationNamespaceName             Attribute
	DataCenterServiceIpAddress                Attribute
	AwsApplicationLoadBalancerTags            Attribute
	HostAzureWebApplicationHostNames          Attribute
	ServiceWebServerName                      Attribute
	DataCenterServiceName                     Attribute
	AwsAutoScalingGroupTags                   Attribute
	ProcessGroupAzureSiteName                 Attribute
	ServiceWebServiceNamespace                Attribute
	HostBitness                               Attribute
	HostPaasType                              Attribute
	QueueVendor                               Attribute
	ServiceIbmCtgGatewayUrl                   Attribute
	DataCenterServicePort                     Attribute
	ServiceTechnology                         Attribute
	AwsRelationalDatabaseServiceTags          Attribute
	ServiceCtgServiceName                     Attribute
	HostAzureComputeMode                      Attribute
	AzureMgmtGroupName                        Attribute
	KubernetesClusterName                     Attribute
	ServiceDetectedName                       Attribute
	AwsClassicLoadBalancerTags                Attribute
	AwsApplicationLoadBalancerName            Attribute
	ServiceDatabaseVendor                     Attribute
	AzureTenantName                           Attribute
	AzureScaleSetName                         Attribute
	ServiceType                               Attribute
	Ec2InstanceAwsSecurityGroup               Attribute
	ProcessGroupListenPort                    Attribute
	MobileApplicationTags                     Attribute
	HostOsVersion                             Attribute
	DockerImageVersion                        Attribute
	AwsClassicLoadBalancerFrontendPorts       Attribute
	HostAzureWebApplicationSiteNames          Attribute
	ProcessGroupName                          Attribute
	MobileApplicationName                     Attribute
	AwsRelationalDatabaseServiceName          Attribute
	GoogleComputeInstancePublicIpAddresses    Attribute
	AppmonServerName                          Attribute
	ServiceWebApplicationId                   Attribute
	AzureMgmtGroupUuid                        Attribute
	DockerStrippedImageName                   Attribute
	HostBoshInstanceName                      Attribute
	CustomDeviceGroupName                     Attribute
	CustomApplicationPlatform                 Attribute
	AzureEntityName                           Attribute
	CustomDeviceMetadata                      Attribute
	HostAwsNameTag                            Attribute
	AzureEntityTags                           Attribute
	ServiceRemoteEndpoint                     Attribute
	Ec2InstancePrivateHostName                Attribute
	WebApplicationType                        Attribute
	AwsRelationalDatabaseServicePort          Attribute
	HostCloudType                             Attribute
	AwsNetworkLoadBalancerName                Attribute
	BrowserMonitorTags                        Attribute
	ProcessGroupTechnologyVersion             Attribute
	ExternalMonitorEngineDescription          Attribute
	CustomDeviceTags                          Attribute
	Ec2InstanceBeanstalkEnvName               Attribute
	GoogleCloudPlatformZoneName               Attribute
	ServiceDatabaseTopology                   Attribute
	NameOfComputeNode                         Attribute
	ServiceRemoteServiceName                  Attribute
	ServiceMessagingListenerClassName         Attribute
	VmwareVmName                              Attribute
	AzureRegionName                           Attribute
	EnterpriseApplicationIpAddress            Attribute
	CustomDeviceDnsAddress                    Attribute
	ProcessGroupId                            Attribute
	CustomDeviceTechnology                    Attribute
	DataCenterServiceDecoderType              Attribute
	ServiceTechnologyVersion                  Attribute
	CloudFoundryOrgName                       Attribute
	AzureSubscriptionName                     Attribute
	ProcessGroupDetectedName                  Attribute
	Ec2InstancePublicHostName                 Attribute
	EsxiHostProductVersion                    Attribute
	CustomDeviceGroupTags                     Attribute
	ServiceName                               Attribute
	GoogleComputeInstanceName                 Attribute
	EsxiHostHardwareModel                     Attribute
	CustomDeviceName                          Attribute
	AzureVmName                               Attribute
	ServiceEsbApplicationName                 Attribute
	ServiceWebContextRoot                     Attribute
	EnterpriseApplicationName                 Attribute
	QueueTechnology                           Attribute
	EsxiHostTags                              Attribute
	KubernetesServiceName                     Attribute
	MobileApplicationPlatform                 Attribute
	HostAzureSku                              Attribute
	HostBoshName                              Attribute
	CustomApplicationType                     Attribute
	HostBoshAvailabilityZone                  Attribute
	ServicePublicDomainName                   Attribute
	AzureTenantUuid                           Attribute
	ProcessGroupTags                          Attribute
	HostOsType                                Attribute
	HostAixSimultaneousThreads                Attribute
	ServiceWebServerEndpoint                  Attribute
	AwsRelationalDatabaseServiceEngine        Attribute
	ServiceTechnologyEdition                  Attribute
	AwsNetworkLoadBalancerTags                Attribute
	KubernetesNodeName                        Attribute
	GeolocationSiteName                       Attribute
	ServiceTopology                           Attribute
	AppmonSystemProfileName                   Attribute
	BrowserMonitorName                        Attribute
	HostHypervisorType                        Attribute
	HostName                                  Attribute
	HostGroupId                               Attribute
	ServiceAkkaActorSystem                    Attribute
	EnterpriseApplicationDecoderType          Attribute
	Ec2InstanceAwsInstanceType                Attribute
	HostDetectedName                          Attribute
	GoogleComputeInstanceId                   Attribute
	OpenstackAvailabilityZoneName             Attribute
	EnterpriseApplicationTags                 Attribute
	HostBoshStemcellVersion                   Attribute
	CustomApplicationName                     Attribute
	HostPaasMemoryLimit                       Attribute
	ProcessGroupPredefinedMetadata            Attribute
	ProcessGroupTechnology                    Attribute
	ServiceDatabaseName                       Attribute
}{
	Attribute("EC2_INSTANCE_ID"),
	Attribute("OPENSTACK_VM_NAME"),
	Attribute("OPENSTACK_VM_SECURITY_GROUP"),
	Attribute("EXTERNAL_MONITOR_NAME"),
	Attribute("WEB_APPLICATION_TAGS"),
	Attribute("VMWARE_DATACENTER_NAME"),
	Attribute("OPENSTACK_VM_INSTANCE_TYPE"),
	Attribute("PROCESS_GROUP_CUSTOM_METADATA"),
	Attribute("CLOUD_APPLICATION_NAME"),
	Attribute("ESXI_HOST_CLUSTER_NAME"),
	Attribute("EXTERNAL_MONITOR_TAGS"),
	Attribute("HOST_CPU_CORES"),
	Attribute("DATA_CENTER_SERVICE_METADATA"),
	Attribute("AWS_AVAILABILITY_ZONE_NAME"),
	Attribute("DOCKER_FULL_IMAGE_NAME"),
	Attribute("CUSTOM_DEVICE_IP_ADDRESS"),
	Attribute("CUSTOM_APPLICATION_TAGS"),
	Attribute("CLOUD_FOUNDRY_FOUNDATION_NAME"),
	Attribute("SERVICE_WEB_SERVICE_NAME"),
	Attribute("PROCESS_GROUP_TECHNOLOGY_EDITION"),
	Attribute("QUEUE_NAME"),
	Attribute("HOST_AIX_VIRTUAL_CPU_COUNT"),
	Attribute("HOST_AIX_LOGICAL_CPU_COUNT"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT"),
	Attribute("OPENSTACK_PROJECT_NAME"),
	Attribute("ESXI_HOST_PRODUCT_NAME"),
	Attribute("CUSTOM_DEVICE_PORT"),
	Attribute("DOCKER_CONTAINER_NAME"),
	Attribute("HTTP_MONITOR_TAGS"),
	Attribute("DATA_CENTER_SERVICE_TAGS"),
	Attribute("SERVICE_PORT"),
	Attribute("SERVICE_DATABASE_HOST_NAME"),
	Attribute("CLOUD_APPLICATION_NAMESPACE_LABELS"),
	Attribute("AWS_ACCOUNT_ID"),
	Attribute("ESXI_HOST_NAME"),
	Attribute("EC2_INSTANCE_AMI_ID"),
	Attribute("EXTERNAL_MONITOR_ENGINE_TYPE"),
	Attribute("HOST_ARCHITECTURE"),
	Attribute("WEB_APPLICATION_NAME_PATTERN"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME"),
	Attribute("AZURE_SUBSCRIPTION_UUID"),
	Attribute("HOST_ONEAGENT_CUSTOM_HOST_NAME"),
	Attribute("HOST_GROUP_NAME"),
	Attribute("ENTERPRISE_APPLICATION_METADATA"),
	Attribute("EC2_INSTANCE_TAGS"),
	Attribute("AWS_CLASSIC_LOAD_BALANCER_NAME"),
	Attribute("HOST_KUBERNETES_LABELS"),
	Attribute("CLOUD_APPLICATION_LABELS"),
	Attribute("AWS_ACCOUNT_NAME"),
	Attribute("AWS_AUTO_SCALING_GROUP_NAME"),
	Attribute("EXTERNAL_MONITOR_ENGINE_NAME"),
	Attribute("HOST_LOGICAL_CPU_CORES"),
	Attribute("HOST_TECHNOLOGY"),
	Attribute("OPENSTACK_ACCOUNT_NAME"),
	Attribute("HOST_CUSTOM_METADATA"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_PROJECT"),
	Attribute("OPENSTACK_REGION_NAME"),
	Attribute("PROCESS_GROUP_AZURE_HOST_NAME"),
	Attribute("HTTP_MONITOR_NAME"),
	Attribute("HOST_IP_ADDRESS"),
	Attribute("ESXI_HOST_HARDWARE_VENDOR"),
	Attribute("HOST_BOSH_INSTANCE_ID"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE"),
	Attribute("HOST_BOSH_DEPLOYMENT_ID"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_PROJECT_ID"),
	Attribute("OPENSTACK_ACCOUNT_PROJECT_NAME"),
	Attribute("ENTERPRISE_APPLICATION_PORT"),
	Attribute("EC2_INSTANCE_NAME"),
	Attribute("WEB_APPLICATION_NAME"),
	Attribute("SERVICE_TAGS"),
	Attribute("HOST_TAGS"),
	Attribute("CLOUD_APPLICATION_NAMESPACE_NAME"),
	Attribute("DATA_CENTER_SERVICE_IP_ADDRESS"),
	Attribute("AWS_APPLICATION_LOAD_BALANCER_TAGS"),
	Attribute("HOST_AZURE_WEB_APPLICATION_HOST_NAMES"),
	Attribute("SERVICE_WEB_SERVER_NAME"),
	Attribute("DATA_CENTER_SERVICE_NAME"),
	Attribute("AWS_AUTO_SCALING_GROUP_TAGS"),
	Attribute("PROCESS_GROUP_AZURE_SITE_NAME"),
	Attribute("SERVICE_WEB_SERVICE_NAMESPACE"),
	Attribute("HOST_BITNESS"),
	Attribute("HOST_PAAS_TYPE"),
	Attribute("QUEUE_VENDOR"),
	Attribute("SERVICE_IBM_CTG_GATEWAY_URL"),
	Attribute("DATA_CENTER_SERVICE_PORT"),
	Attribute("SERVICE_TECHNOLOGY"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_TAGS"),
	Attribute("SERVICE_CTG_SERVICE_NAME"),
	Attribute("HOST_AZURE_COMPUTE_MODE"),
	Attribute("AZURE_MGMT_GROUP_NAME"),
	Attribute("KUBERNETES_CLUSTER_NAME"),
	Attribute("SERVICE_DETECTED_NAME"),
	Attribute("AWS_CLASSIC_LOAD_BALANCER_TAGS"),
	Attribute("AWS_APPLICATION_LOAD_BALANCER_NAME"),
	Attribute("SERVICE_DATABASE_VENDOR"),
	Attribute("AZURE_TENANT_NAME"),
	Attribute("AZURE_SCALE_SET_NAME"),
	Attribute("SERVICE_TYPE"),
	Attribute("EC2_INSTANCE_AWS_SECURITY_GROUP"),
	Attribute("PROCESS_GROUP_LISTEN_PORT"),
	Attribute("MOBILE_APPLICATION_TAGS"),
	Attribute("HOST_OS_VERSION"),
	Attribute("DOCKER_IMAGE_VERSION"),
	Attribute("AWS_CLASSIC_LOAD_BALANCER_FRONTEND_PORTS"),
	Attribute("HOST_AZURE_WEB_APPLICATION_SITE_NAMES"),
	Attribute("PROCESS_GROUP_NAME"),
	Attribute("MOBILE_APPLICATION_NAME"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_NAME"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_PUBLIC_IP_ADDRESSES"),
	Attribute("APPMON_SERVER_NAME"),
	Attribute("SERVICE_WEB_APPLICATION_ID"),
	Attribute("AZURE_MGMT_GROUP_UUID"),
	Attribute("DOCKER_STRIPPED_IMAGE_NAME"),
	Attribute("HOST_BOSH_INSTANCE_NAME"),
	Attribute("CUSTOM_DEVICE_GROUP_NAME"),
	Attribute("CUSTOM_APPLICATION_PLATFORM"),
	Attribute("AZURE_ENTITY_NAME"),
	Attribute("CUSTOM_DEVICE_METADATA"),
	Attribute("HOST_AWS_NAME_TAG"),
	Attribute("AZURE_ENTITY_TAGS"),
	Attribute("SERVICE_REMOTE_ENDPOINT"),
	Attribute("EC2_INSTANCE_PRIVATE_HOST_NAME"),
	Attribute("WEB_APPLICATION_TYPE"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_PORT"),
	Attribute("HOST_CLOUD_TYPE"),
	Attribute("AWS_NETWORK_LOAD_BALANCER_NAME"),
	Attribute("BROWSER_MONITOR_TAGS"),
	Attribute("PROCESS_GROUP_TECHNOLOGY_VERSION"),
	Attribute("EXTERNAL_MONITOR_ENGINE_DESCRIPTION"),
	Attribute("CUSTOM_DEVICE_TAGS"),
	Attribute("EC2_INSTANCE_BEANSTALK_ENV_NAME"),
	Attribute("GOOGLE_CLOUD_PLATFORM_ZONE_NAME"),
	Attribute("SERVICE_DATABASE_TOPOLOGY"),
	Attribute("NAME_OF_COMPUTE_NODE"),
	Attribute("SERVICE_REMOTE_SERVICE_NAME"),
	Attribute("SERVICE_MESSAGING_LISTENER_CLASS_NAME"),
	Attribute("VMWARE_VM_NAME"),
	Attribute("AZURE_REGION_NAME"),
	Attribute("ENTERPRISE_APPLICATION_IP_ADDRESS"),
	Attribute("CUSTOM_DEVICE_DNS_ADDRESS"),
	Attribute("PROCESS_GROUP_ID"),
	Attribute("CUSTOM_DEVICE_TECHNOLOGY"),
	Attribute("DATA_CENTER_SERVICE_DECODER_TYPE"),
	Attribute("SERVICE_TECHNOLOGY_VERSION"),
	Attribute("CLOUD_FOUNDRY_ORG_NAME"),
	Attribute("AZURE_SUBSCRIPTION_NAME"),
	Attribute("PROCESS_GROUP_DETECTED_NAME"),
	Attribute("EC2_INSTANCE_PUBLIC_HOST_NAME"),
	Attribute("ESXI_HOST_PRODUCT_VERSION"),
	Attribute("CUSTOM_DEVICE_GROUP_TAGS"),
	Attribute("SERVICE_NAME"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_NAME"),
	Attribute("ESXI_HOST_HARDWARE_MODEL"),
	Attribute("CUSTOM_DEVICE_NAME"),
	Attribute("AZURE_VM_NAME"),
	Attribute("SERVICE_ESB_APPLICATION_NAME"),
	Attribute("SERVICE_WEB_CONTEXT_ROOT"),
	Attribute("ENTERPRISE_APPLICATION_NAME"),
	Attribute("QUEUE_TECHNOLOGY"),
	Attribute("ESXI_HOST_TAGS"),
	Attribute("KUBERNETES_SERVICE_NAME"),
	Attribute("MOBILE_APPLICATION_PLATFORM"),
	Attribute("HOST_AZURE_SKU"),
	Attribute("HOST_BOSH_NAME"),
	Attribute("CUSTOM_APPLICATION_TYPE"),
	Attribute("HOST_BOSH_AVAILABILITY_ZONE"),
	Attribute("SERVICE_PUBLIC_DOMAIN_NAME"),
	Attribute("AZURE_TENANT_UUID"),
	Attribute("PROCESS_GROUP_TAGS"),
	Attribute("HOST_OS_TYPE"),
	Attribute("HOST_AIX_SIMULTANEOUS_THREADS"),
	Attribute("SERVICE_WEB_SERVER_ENDPOINT"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_ENGINE"),
	Attribute("SERVICE_TECHNOLOGY_EDITION"),
	Attribute("AWS_NETWORK_LOAD_BALANCER_TAGS"),
	Attribute("KUBERNETES_NODE_NAME"),
	Attribute("GEOLOCATION_SITE_NAME"),
	Attribute("SERVICE_TOPOLOGY"),
	Attribute("APPMON_SYSTEM_PROFILE_NAME"),
	Attribute("BROWSER_MONITOR_NAME"),
	Attribute("HOST_HYPERVISOR_TYPE"),
	Attribute("HOST_NAME"),
	Attribute("HOST_GROUP_ID"),
	Attribute("SERVICE_AKKA_ACTOR_SYSTEM"),
	Attribute("ENTERPRISE_APPLICATION_DECODER_TYPE"),
	Attribute("EC2_INSTANCE_AWS_INSTANCE_TYPE"),
	Attribute("HOST_DETECTED_NAME"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_ID"),
	Attribute("OPENSTACK_AVAILABILITY_ZONE_NAME"),
	Attribute("ENTERPRISE_APPLICATION_TAGS"),
	Attribute("HOST_BOSH_STEMCELL_VERSION"),
	Attribute("CUSTOM_APPLICATION_NAME"),
	Attribute("HOST_PAAS_MEMORY_LIMIT"),
	Attribute("PROCESS_GROUP_PREDEFINED_METADATA"),
	Attribute("PROCESS_GROUP_TECHNOLOGY"),
	Attribute("SERVICE_DATABASE_NAME"),
}

type RuleType string

var RuleTypes = struct {
	Me        RuleType
	Dimension RuleType
	Selector  RuleType
}{
	RuleType("ME"),
	RuleType("DIMENSION"),
	RuleType("SELECTOR"),
}

type Operator string

var Operators = struct {
	NotBeginsWith         Operator
	GreaterThanOrEqual    Operator
	NotExists             Operator
	TagKeyEquals          Operator
	IsIpInRange           Operator
	Exists                Operator
	NotTagKeyEquals       Operator
	BeginsWith            Operator
	NotIsIpInRange        Operator
	Contains              Operator
	NotContains           Operator
	EndsWith              Operator
	NotEndsWith           Operator
	LowerThan             Operator
	NotGreaterThan        Operator
	LowerThanOrEqual      Operator
	NotRegexMatches       Operator
	GreaterThan           Operator
	NotGreaterThanOrEqual Operator
	NotEquals             Operator
	Equals                Operator
	NotLowerThan          Operator
	NotLowerThanOrEqual   Operator
	RegexMatches          Operator
}{
	Operator("NOT_BEGINS_WITH"),
	Operator("GREATER_THAN_OR_EQUAL"),
	Operator("NOT_EXISTS"),
	Operator("TAG_KEY_EQUALS"),
	Operator("IS_IP_IN_RANGE"),
	Operator("EXISTS"),
	Operator("NOT_TAG_KEY_EQUALS"),
	Operator("BEGINS_WITH"),
	Operator("NOT_IS_IP_IN_RANGE"),
	Operator("CONTAINS"),
	Operator("NOT_CONTAINS"),
	Operator("ENDS_WITH"),
	Operator("NOT_ENDS_WITH"),
	Operator("LOWER_THAN"),
	Operator("NOT_GREATER_THAN"),
	Operator("LOWER_THAN_OR_EQUAL"),
	Operator("NOT_REGEX_MATCHES"),
	Operator("GREATER_THAN"),
	Operator("NOT_GREATER_THAN_OR_EQUAL"),
	Operator("NOT_EQUALS"),
	Operator("EQUALS"),
	Operator("NOT_LOWER_THAN"),
	Operator("NOT_LOWER_THAN_OR_EQUAL"),
	Operator("REGEX_MATCHES"),
}

type DimensionOperator string

var DimensionOperators = struct {
	BeginsWith DimensionOperator
	Equals     DimensionOperator
}{
	DimensionOperator("BEGINS_WITH"),
	DimensionOperator("EQUALS"),
}

type DimensionConditionType string

var DimensionConditionTypes = struct {
	LogFileName DimensionConditionType
	MetricKey   DimensionConditionType
	Dimension   DimensionConditionType
}{
	DimensionConditionType("LOG_FILE_NAME"),
	DimensionConditionType("METRIC_KEY"),
	DimensionConditionType("DIMENSION"),
}

type ManagementZoneMeType string

var ManagementZoneMeTypes = struct {
	CloudFoundryFoundation       ManagementZoneMeType
	WebApplication               ManagementZoneMeType
	CloudApplicationNamespace    ManagementZoneMeType
	Service                      ManagementZoneMeType
	Host                         ManagementZoneMeType
	Queue                        ManagementZoneMeType
	ProcessGroup                 ManagementZoneMeType
	CustomApplication            ManagementZoneMeType
	OpenstackAccount             ManagementZoneMeType
	BrowserMonitor               ManagementZoneMeType
	MobileApplication            ManagementZoneMeType
	CloudApplication             ManagementZoneMeType
	AwsApplicationLoadBalancer   ManagementZoneMeType
	AwsNetworkLoadBalancer       ManagementZoneMeType
	AwsClassicLoadBalancer       ManagementZoneMeType
	AppmonSystemProfile          ManagementZoneMeType
	AwsRelationalDatabaseService ManagementZoneMeType
	CustomDeviceGroup            ManagementZoneMeType
	CustomDevice                 ManagementZoneMeType
	HostGroup                    ManagementZoneMeType
	EsxiHost                     ManagementZoneMeType
	EnterpriseApplication        ManagementZoneMeType
	ExternalMonitor              ManagementZoneMeType
	KubernetesCluster            ManagementZoneMeType
	KubernetesService            ManagementZoneMeType
	Azure                        ManagementZoneMeType
	DataCenterService            ManagementZoneMeType
	AwsAccount                   ManagementZoneMeType
	AppmonServer                 ManagementZoneMeType
	AwsAutoScalingGroup          ManagementZoneMeType
	HttpMonitor                  ManagementZoneMeType
}{
	ManagementZoneMeType("CLOUD_FOUNDRY_FOUNDATION"),
	ManagementZoneMeType("WEB_APPLICATION"),
	ManagementZoneMeType("CLOUD_APPLICATION_NAMESPACE"),
	ManagementZoneMeType("SERVICE"),
	ManagementZoneMeType("HOST"),
	ManagementZoneMeType("QUEUE"),
	ManagementZoneMeType("PROCESS_GROUP"),
	ManagementZoneMeType("CUSTOM_APPLICATION"),
	ManagementZoneMeType("OPENSTACK_ACCOUNT"),
	ManagementZoneMeType("BROWSER_MONITOR"),
	ManagementZoneMeType("MOBILE_APPLICATION"),
	ManagementZoneMeType("CLOUD_APPLICATION"),
	ManagementZoneMeType("AWS_APPLICATION_LOAD_BALANCER"),
	ManagementZoneMeType("AWS_NETWORK_LOAD_BALANCER"),
	ManagementZoneMeType("AWS_CLASSIC_LOAD_BALANCER"),
	ManagementZoneMeType("APPMON_SYSTEM_PROFILE"),
	ManagementZoneMeType("AWS_RELATIONAL_DATABASE_SERVICE"),
	ManagementZoneMeType("CUSTOM_DEVICE_GROUP"),
	ManagementZoneMeType("CUSTOM_DEVICE"),
	ManagementZoneMeType("HOST_GROUP"),
	ManagementZoneMeType("ESXI_HOST"),
	ManagementZoneMeType("ENTERPRISE_APPLICATION"),
	ManagementZoneMeType("EXTERNAL_MONITOR"),
	ManagementZoneMeType("KUBERNETES_CLUSTER"),
	ManagementZoneMeType("KUBERNETES_SERVICE"),
	ManagementZoneMeType("AZURE"),
	ManagementZoneMeType("DATA_CENTER_SERVICE"),
	ManagementZoneMeType("AWS_ACCOUNT"),
	ManagementZoneMeType("APPMON_SERVER"),
	ManagementZoneMeType("AWS_AUTO_SCALING_GROUP"),
	ManagementZoneMeType("HTTP_MONITOR"),
}
