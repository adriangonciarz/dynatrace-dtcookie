package autotagging

type RuleType string

var RuleTypes = struct {
	Me       RuleType
	Selector RuleType
}{
	RuleType("ME"),
	RuleType("SELECTOR"),
}

type AutoTagMeType string

var AutoTagMeTypes = struct {
	EsxiHost                     AutoTagMeType
	Application                  AutoTagMeType
	AwsRelationalDatabaseService AutoTagMeType
	AwsNetworkLoadBalancer       AutoTagMeType
	CustomApplication            AutoTagMeType
	AwsClassicLoadBalancer       AutoTagMeType
	SyntheticTest                AutoTagMeType
	ExternalSyntheticTest        AutoTagMeType
	CustomDevice                 AutoTagMeType
	Service                      AutoTagMeType
	Host                         AutoTagMeType
	MobileApplication            AutoTagMeType
	AwsApplicationLoadBalancer   AutoTagMeType
	Azure                        AutoTagMeType
	ProcessGroup                 AutoTagMeType
	HttpCheck                    AutoTagMeType
	DcrumApplication             AutoTagMeType
}{
	AutoTagMeType("ESXI_HOST"),
	AutoTagMeType("APPLICATION"),
	AutoTagMeType("AWS_RELATIONAL_DATABASE_SERVICE"),
	AutoTagMeType("AWS_NETWORK_LOAD_BALANCER"),
	AutoTagMeType("CUSTOM_APPLICATION"),
	AutoTagMeType("AWS_CLASSIC_LOAD_BALANCER"),
	AutoTagMeType("SYNTHETIC_TEST"),
	AutoTagMeType("EXTERNAL_SYNTHETIC_TEST"),
	AutoTagMeType("CUSTOM_DEVICE"),
	AutoTagMeType("SERVICE"),
	AutoTagMeType("HOST"),
	AutoTagMeType("MOBILE_APPLICATION"),
	AutoTagMeType("AWS_APPLICATION_LOAD_BALANCER"),
	AutoTagMeType("AZURE"),
	AutoTagMeType("PROCESS_GROUP"),
	AutoTagMeType("HTTP_CHECK"),
	AutoTagMeType("DCRUM_APPLICATION"),
}

type Attribute string

var Attributes = struct {
	GeolocationSiteName                       Attribute
	DataCenterServiceTags                     Attribute
	HostLogicalCpuCores                       Attribute
	DataCenterServiceMetadata                 Attribute
	BrowserMonitorTags                        Attribute
	QueueVendor                               Attribute
	Ec2InstancePrivateHostName                Attribute
	ServiceIbmCtgGatewayUrl                   Attribute
	ServiceMessagingListenerClassName         Attribute
	WebApplicationName                        Attribute
	HostCloudType                             Attribute
	HostOsType                                Attribute
	CloudApplicationName                      Attribute
	ProcessGroupName                          Attribute
	EnterpriseApplicationTags                 Attribute
	ServiceRemoteServiceName                  Attribute
	AzureVmName                               Attribute
	DataCenterServiceName                     Attribute
	HostAixSimultaneousThreads                Attribute
	HostTechnology                            Attribute
	ProcessGroupTags                          Attribute
	GoogleComputeInstanceProject              Attribute
	ExternalMonitorEngineName                 Attribute
	Ec2InstancePublicHostName                 Attribute
	GoogleComputeInstanceProjectId            Attribute
	EsxiHostProductVersion                    Attribute
	HostBoshInstanceName                      Attribute
	DockerContainerName                       Attribute
	AwsNetworkLoadBalancerName                Attribute
	HostBitness                               Attribute
	KubernetesServiceName                     Attribute
	AzureRegionName                           Attribute
	ExternalMonitorTags                       Attribute
	Ec2InstanceId                             Attribute
	DataCenterServicePort                     Attribute
	EsxiHostTags                              Attribute
	ServiceDatabaseTopology                   Attribute
	AwsAutoScalingGroupTags                   Attribute
	ServiceDetectedName                       Attribute
	HostBoshInstanceId                        Attribute
	ServiceType                               Attribute
	HostOsVersion                             Attribute
	ServiceTechnologyVersion                  Attribute
	Ec2InstanceAmiId                          Attribute
	OpenstackVmSecurityGroup                  Attribute
	ServicePublicDomainName                   Attribute
	EnterpriseApplicationDecoderType          Attribute
	EnterpriseApplicationPort                 Attribute
	KubernetesNodeName                        Attribute
	ServiceWebServerEndpoint                  Attribute
	CustomDeviceMetadata                      Attribute
	ServiceName                               Attribute
	AwsClassicLoadBalancerName                Attribute
	OpenstackAccountName                      Attribute
	DockerImageVersion                        Attribute
	AwsAccountId                              Attribute
	AzureMgmtGroupName                        Attribute
	CustomDeviceIpAddress                     Attribute
	CloudFoundryOrgName                       Attribute
	HttpMonitorTags                           Attribute
	ServiceDatabaseName                       Attribute
	HostBoshName                              Attribute
	ServiceEsbApplicationName                 Attribute
	OpenstackAccountProjectName               Attribute
	ServiceWebServiceNamespace                Attribute
	AwsApplicationLoadBalancerTags            Attribute
	ProcessGroupListenPort                    Attribute
	GoogleComputeInstanceName                 Attribute
	HostGroupId                               Attribute
	DataCenterServiceDecoderType              Attribute
	CloudFoundryFoundationName                Attribute
	AzureEntityTags                           Attribute
	GoogleComputeInstanceId                   Attribute
	OpenstackVmInstanceType                   Attribute
	GoogleComputeInstancePublicIpAddresses    Attribute
	ServiceWebContextRoot                     Attribute
	OpenstackVmName                           Attribute
	ServiceWebServiceName                     Attribute
	HostAixVirtualCpuCount                    Attribute
	NameOfComputeNode                         Attribute
	CloudApplicationLabels                    Attribute
	WebApplicationNamePattern                 Attribute
	CustomDevicePort                          Attribute
	EnterpriseApplicationName                 Attribute
	ProcessGroupCustomMetadata                Attribute
	ExternalMonitorEngineDescription          Attribute
	EnterpriseApplicationMetadata             Attribute
	ServiceDatabaseVendor                     Attribute
	ProcessGroupAzureHostName                 Attribute
	AwsRelationalDatabaseServiceEngine        Attribute
	DataCenterServiceIpAddress                Attribute
	AwsAccountName                            Attribute
	QueueName                                 Attribute
	ServiceRemoteEndpoint                     Attribute
	DockerFullImageName                       Attribute
	OpenstackAvailabilityZoneName             Attribute
	HostBoshDeploymentId                      Attribute
	HostArchitecture                          Attribute
	HostAzureWebApplicationSiteNames          Attribute
	AwsRelationalDatabaseServiceInstanceClass Attribute
	ProcessGroupTechnologyVersion             Attribute
	Ec2InstanceAwsSecurityGroup               Attribute
	CustomApplicationTags                     Attribute
	ServiceWebServerName                      Attribute
	AwsRelationalDatabaseServiceDbName        Attribute
	HostIpAddress                             Attribute
	AwsClassicLoadBalancerFrontendPorts       Attribute
	HostAwsNameTag                            Attribute
	HttpMonitorName                           Attribute
	BrowserMonitorName                        Attribute
	AwsAutoScalingGroupName                   Attribute
	HostTags                                  Attribute
	ProcessGroupTechnologyEdition             Attribute
	AwsRelationalDatabaseServiceTags          Attribute
	HostBoshStemcellVersion                   Attribute
	CloudApplicationNamespaceName             Attribute
	EsxiHostHardwareModel                     Attribute
	QueueTechnology                           Attribute
	VmwareDatacenterName                      Attribute
	MobileApplicationPlatform                 Attribute
	HostCpuCores                              Attribute
	Ec2InstanceName                           Attribute
	Ec2InstanceTags                           Attribute
	HostAixLogicalCpuCount                    Attribute
	ServiceTechnology                         Attribute
	AppmonServerName                          Attribute
	HostHypervisorType                        Attribute
	CustomDeviceGroupTags                     Attribute
	CustomApplicationName                     Attribute
	ExternalMonitorEngineType                 Attribute
	AwsRelationalDatabaseServicePort          Attribute
	ServiceTags                               Attribute
	CustomApplicationPlatform                 Attribute
	AwsNetworkLoadBalancerTags                Attribute
	CustomDeviceGroupName                     Attribute
	ServicePort                               Attribute
	ProcessGroupId                            Attribute
	AzureTenantUuid                           Attribute
	AzureMgmtGroupUuid                        Attribute
	AzureTenantName                           Attribute
	OpenstackRegionName                       Attribute
	EsxiHostClusterName                       Attribute
	EsxiHostHardwareVendor                    Attribute
	HostAzureSku                              Attribute
	HostDetectedName                          Attribute
	AppmonSystemProfileName                   Attribute
	ServiceTechnologyEdition                  Attribute
	OpenstackProjectName                      Attribute
	ServiceDatabaseHostName                   Attribute
	AwsRelationalDatabaseServiceEndpoint      Attribute
	AwsRelationalDatabaseServiceName          Attribute
	ProcessGroupTechnology                    Attribute
	DockerStrippedImageName                   Attribute
	MobileApplicationName                     Attribute
	AwsApplicationLoadBalancerName            Attribute
	CustomApplicationType                     Attribute
	CloudApplicationNamespaceLabels           Attribute
	ServiceAkkaActorSystem                    Attribute
	GoogleComputeInstanceMachineType          Attribute
	Ec2InstanceAwsInstanceType                Attribute
	KubernetesClusterName                     Attribute
	EnterpriseApplicationIpAddress            Attribute
	AwsClassicLoadBalancerTags                Attribute
	HostAzureWebApplicationHostNames          Attribute
	CustomDeviceDnsAddress                    Attribute
	VmwareVmName                              Attribute
	AzureSubscriptionUuid                     Attribute
	AzureScaleSetName                         Attribute
	HostPaasType                              Attribute
	ProcessGroupDetectedName                  Attribute
	HostPaasMemoryLimit                       Attribute
	CustomDeviceName                          Attribute
	AwsAvailabilityZoneName                   Attribute
	HostAzureComputeMode                      Attribute
	Ec2InstanceBeanstalkEnvName               Attribute
	ExternalMonitorName                       Attribute
	MobileApplicationTags                     Attribute
	AzureEntityName                           Attribute
	ProcessGroupAzureSiteName                 Attribute
	ServiceWebApplicationId                   Attribute
	CustomDeviceTechnology                    Attribute
	ProcessGroupPredefinedMetadata            Attribute
	WebApplicationTags                        Attribute
	ServiceTopology                           Attribute
	HostBoshAvailabilityZone                  Attribute
	HostName                                  Attribute
	WebApplicationType                        Attribute
	GoogleCloudPlatformZoneName               Attribute
	ServiceCtgServiceName                     Attribute
	HostKubernetesLabels                      Attribute
	HostGroupName                             Attribute
	EsxiHostProductName                       Attribute
	CustomDeviceTags                          Attribute
	EsxiHostName                              Attribute
	AzureSubscriptionName                     Attribute
	HostCustomMetadata                        Attribute
	HostOneagentCustomHostName                Attribute
}{
	Attribute("GEOLOCATION_SITE_NAME"),
	Attribute("DATA_CENTER_SERVICE_TAGS"),
	Attribute("HOST_LOGICAL_CPU_CORES"),
	Attribute("DATA_CENTER_SERVICE_METADATA"),
	Attribute("BROWSER_MONITOR_TAGS"),
	Attribute("QUEUE_VENDOR"),
	Attribute("EC2_INSTANCE_PRIVATE_HOST_NAME"),
	Attribute("SERVICE_IBM_CTG_GATEWAY_URL"),
	Attribute("SERVICE_MESSAGING_LISTENER_CLASS_NAME"),
	Attribute("WEB_APPLICATION_NAME"),
	Attribute("HOST_CLOUD_TYPE"),
	Attribute("HOST_OS_TYPE"),
	Attribute("CLOUD_APPLICATION_NAME"),
	Attribute("PROCESS_GROUP_NAME"),
	Attribute("ENTERPRISE_APPLICATION_TAGS"),
	Attribute("SERVICE_REMOTE_SERVICE_NAME"),
	Attribute("AZURE_VM_NAME"),
	Attribute("DATA_CENTER_SERVICE_NAME"),
	Attribute("HOST_AIX_SIMULTANEOUS_THREADS"),
	Attribute("HOST_TECHNOLOGY"),
	Attribute("PROCESS_GROUP_TAGS"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_PROJECT"),
	Attribute("EXTERNAL_MONITOR_ENGINE_NAME"),
	Attribute("EC2_INSTANCE_PUBLIC_HOST_NAME"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_PROJECT_ID"),
	Attribute("ESXI_HOST_PRODUCT_VERSION"),
	Attribute("HOST_BOSH_INSTANCE_NAME"),
	Attribute("DOCKER_CONTAINER_NAME"),
	Attribute("AWS_NETWORK_LOAD_BALANCER_NAME"),
	Attribute("HOST_BITNESS"),
	Attribute("KUBERNETES_SERVICE_NAME"),
	Attribute("AZURE_REGION_NAME"),
	Attribute("EXTERNAL_MONITOR_TAGS"),
	Attribute("EC2_INSTANCE_ID"),
	Attribute("DATA_CENTER_SERVICE_PORT"),
	Attribute("ESXI_HOST_TAGS"),
	Attribute("SERVICE_DATABASE_TOPOLOGY"),
	Attribute("AWS_AUTO_SCALING_GROUP_TAGS"),
	Attribute("SERVICE_DETECTED_NAME"),
	Attribute("HOST_BOSH_INSTANCE_ID"),
	Attribute("SERVICE_TYPE"),
	Attribute("HOST_OS_VERSION"),
	Attribute("SERVICE_TECHNOLOGY_VERSION"),
	Attribute("EC2_INSTANCE_AMI_ID"),
	Attribute("OPENSTACK_VM_SECURITY_GROUP"),
	Attribute("SERVICE_PUBLIC_DOMAIN_NAME"),
	Attribute("ENTERPRISE_APPLICATION_DECODER_TYPE"),
	Attribute("ENTERPRISE_APPLICATION_PORT"),
	Attribute("KUBERNETES_NODE_NAME"),
	Attribute("SERVICE_WEB_SERVER_ENDPOINT"),
	Attribute("CUSTOM_DEVICE_METADATA"),
	Attribute("SERVICE_NAME"),
	Attribute("AWS_CLASSIC_LOAD_BALANCER_NAME"),
	Attribute("OPENSTACK_ACCOUNT_NAME"),
	Attribute("DOCKER_IMAGE_VERSION"),
	Attribute("AWS_ACCOUNT_ID"),
	Attribute("AZURE_MGMT_GROUP_NAME"),
	Attribute("CUSTOM_DEVICE_IP_ADDRESS"),
	Attribute("CLOUD_FOUNDRY_ORG_NAME"),
	Attribute("HTTP_MONITOR_TAGS"),
	Attribute("SERVICE_DATABASE_NAME"),
	Attribute("HOST_BOSH_NAME"),
	Attribute("SERVICE_ESB_APPLICATION_NAME"),
	Attribute("OPENSTACK_ACCOUNT_PROJECT_NAME"),
	Attribute("SERVICE_WEB_SERVICE_NAMESPACE"),
	Attribute("AWS_APPLICATION_LOAD_BALANCER_TAGS"),
	Attribute("PROCESS_GROUP_LISTEN_PORT"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_NAME"),
	Attribute("HOST_GROUP_ID"),
	Attribute("DATA_CENTER_SERVICE_DECODER_TYPE"),
	Attribute("CLOUD_FOUNDRY_FOUNDATION_NAME"),
	Attribute("AZURE_ENTITY_TAGS"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_ID"),
	Attribute("OPENSTACK_VM_INSTANCE_TYPE"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_PUBLIC_IP_ADDRESSES"),
	Attribute("SERVICE_WEB_CONTEXT_ROOT"),
	Attribute("OPENSTACK_VM_NAME"),
	Attribute("SERVICE_WEB_SERVICE_NAME"),
	Attribute("HOST_AIX_VIRTUAL_CPU_COUNT"),
	Attribute("NAME_OF_COMPUTE_NODE"),
	Attribute("CLOUD_APPLICATION_LABELS"),
	Attribute("WEB_APPLICATION_NAME_PATTERN"),
	Attribute("CUSTOM_DEVICE_PORT"),
	Attribute("ENTERPRISE_APPLICATION_NAME"),
	Attribute("PROCESS_GROUP_CUSTOM_METADATA"),
	Attribute("EXTERNAL_MONITOR_ENGINE_DESCRIPTION"),
	Attribute("ENTERPRISE_APPLICATION_METADATA"),
	Attribute("SERVICE_DATABASE_VENDOR"),
	Attribute("PROCESS_GROUP_AZURE_HOST_NAME"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_ENGINE"),
	Attribute("DATA_CENTER_SERVICE_IP_ADDRESS"),
	Attribute("AWS_ACCOUNT_NAME"),
	Attribute("QUEUE_NAME"),
	Attribute("SERVICE_REMOTE_ENDPOINT"),
	Attribute("DOCKER_FULL_IMAGE_NAME"),
	Attribute("OPENSTACK_AVAILABILITY_ZONE_NAME"),
	Attribute("HOST_BOSH_DEPLOYMENT_ID"),
	Attribute("HOST_ARCHITECTURE"),
	Attribute("HOST_AZURE_WEB_APPLICATION_SITE_NAMES"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS"),
	Attribute("PROCESS_GROUP_TECHNOLOGY_VERSION"),
	Attribute("EC2_INSTANCE_AWS_SECURITY_GROUP"),
	Attribute("CUSTOM_APPLICATION_TAGS"),
	Attribute("SERVICE_WEB_SERVER_NAME"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME"),
	Attribute("HOST_IP_ADDRESS"),
	Attribute("AWS_CLASSIC_LOAD_BALANCER_FRONTEND_PORTS"),
	Attribute("HOST_AWS_NAME_TAG"),
	Attribute("HTTP_MONITOR_NAME"),
	Attribute("BROWSER_MONITOR_NAME"),
	Attribute("AWS_AUTO_SCALING_GROUP_NAME"),
	Attribute("HOST_TAGS"),
	Attribute("PROCESS_GROUP_TECHNOLOGY_EDITION"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_TAGS"),
	Attribute("HOST_BOSH_STEMCELL_VERSION"),
	Attribute("CLOUD_APPLICATION_NAMESPACE_NAME"),
	Attribute("ESXI_HOST_HARDWARE_MODEL"),
	Attribute("QUEUE_TECHNOLOGY"),
	Attribute("VMWARE_DATACENTER_NAME"),
	Attribute("MOBILE_APPLICATION_PLATFORM"),
	Attribute("HOST_CPU_CORES"),
	Attribute("EC2_INSTANCE_NAME"),
	Attribute("EC2_INSTANCE_TAGS"),
	Attribute("HOST_AIX_LOGICAL_CPU_COUNT"),
	Attribute("SERVICE_TECHNOLOGY"),
	Attribute("APPMON_SERVER_NAME"),
	Attribute("HOST_HYPERVISOR_TYPE"),
	Attribute("CUSTOM_DEVICE_GROUP_TAGS"),
	Attribute("CUSTOM_APPLICATION_NAME"),
	Attribute("EXTERNAL_MONITOR_ENGINE_TYPE"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_PORT"),
	Attribute("SERVICE_TAGS"),
	Attribute("CUSTOM_APPLICATION_PLATFORM"),
	Attribute("AWS_NETWORK_LOAD_BALANCER_TAGS"),
	Attribute("CUSTOM_DEVICE_GROUP_NAME"),
	Attribute("SERVICE_PORT"),
	Attribute("PROCESS_GROUP_ID"),
	Attribute("AZURE_TENANT_UUID"),
	Attribute("AZURE_MGMT_GROUP_UUID"),
	Attribute("AZURE_TENANT_NAME"),
	Attribute("OPENSTACK_REGION_NAME"),
	Attribute("ESXI_HOST_CLUSTER_NAME"),
	Attribute("ESXI_HOST_HARDWARE_VENDOR"),
	Attribute("HOST_AZURE_SKU"),
	Attribute("HOST_DETECTED_NAME"),
	Attribute("APPMON_SYSTEM_PROFILE_NAME"),
	Attribute("SERVICE_TECHNOLOGY_EDITION"),
	Attribute("OPENSTACK_PROJECT_NAME"),
	Attribute("SERVICE_DATABASE_HOST_NAME"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_NAME"),
	Attribute("PROCESS_GROUP_TECHNOLOGY"),
	Attribute("DOCKER_STRIPPED_IMAGE_NAME"),
	Attribute("MOBILE_APPLICATION_NAME"),
	Attribute("AWS_APPLICATION_LOAD_BALANCER_NAME"),
	Attribute("CUSTOM_APPLICATION_TYPE"),
	Attribute("CLOUD_APPLICATION_NAMESPACE_LABELS"),
	Attribute("SERVICE_AKKA_ACTOR_SYSTEM"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE"),
	Attribute("EC2_INSTANCE_AWS_INSTANCE_TYPE"),
	Attribute("KUBERNETES_CLUSTER_NAME"),
	Attribute("ENTERPRISE_APPLICATION_IP_ADDRESS"),
	Attribute("AWS_CLASSIC_LOAD_BALANCER_TAGS"),
	Attribute("HOST_AZURE_WEB_APPLICATION_HOST_NAMES"),
	Attribute("CUSTOM_DEVICE_DNS_ADDRESS"),
	Attribute("VMWARE_VM_NAME"),
	Attribute("AZURE_SUBSCRIPTION_UUID"),
	Attribute("AZURE_SCALE_SET_NAME"),
	Attribute("HOST_PAAS_TYPE"),
	Attribute("PROCESS_GROUP_DETECTED_NAME"),
	Attribute("HOST_PAAS_MEMORY_LIMIT"),
	Attribute("CUSTOM_DEVICE_NAME"),
	Attribute("AWS_AVAILABILITY_ZONE_NAME"),
	Attribute("HOST_AZURE_COMPUTE_MODE"),
	Attribute("EC2_INSTANCE_BEANSTALK_ENV_NAME"),
	Attribute("EXTERNAL_MONITOR_NAME"),
	Attribute("MOBILE_APPLICATION_TAGS"),
	Attribute("AZURE_ENTITY_NAME"),
	Attribute("PROCESS_GROUP_AZURE_SITE_NAME"),
	Attribute("SERVICE_WEB_APPLICATION_ID"),
	Attribute("CUSTOM_DEVICE_TECHNOLOGY"),
	Attribute("PROCESS_GROUP_PREDEFINED_METADATA"),
	Attribute("WEB_APPLICATION_TAGS"),
	Attribute("SERVICE_TOPOLOGY"),
	Attribute("HOST_BOSH_AVAILABILITY_ZONE"),
	Attribute("HOST_NAME"),
	Attribute("WEB_APPLICATION_TYPE"),
	Attribute("GOOGLE_CLOUD_PLATFORM_ZONE_NAME"),
	Attribute("SERVICE_CTG_SERVICE_NAME"),
	Attribute("HOST_KUBERNETES_LABELS"),
	Attribute("HOST_GROUP_NAME"),
	Attribute("ESXI_HOST_PRODUCT_NAME"),
	Attribute("CUSTOM_DEVICE_TAGS"),
	Attribute("ESXI_HOST_NAME"),
	Attribute("AZURE_SUBSCRIPTION_NAME"),
	Attribute("HOST_CUSTOM_METADATA"),
	Attribute("HOST_ONEAGENT_CUSTOM_HOST_NAME"),
}

type Operator string

var Operators = struct {
	NotLowerThanOrEqual   Operator
	NotEndsWith           Operator
	NotContains           Operator
	NotRegexMatches       Operator
	NotExists             Operator
	NotGreaterThanOrEqual Operator
	Contains              Operator
	IsIpInRange           Operator
	NotIsIpInRange        Operator
	EndsWith              Operator
	TagKeyEquals          Operator
	NotTagKeyEquals       Operator
	GreaterThan           Operator
	Equals                Operator
	NotEquals             Operator
	LowerThan             Operator
	NotLowerThan          Operator
	NotBeginsWith         Operator
	LowerThanOrEqual      Operator
	NotGreaterThan        Operator
	GreaterThanOrEqual    Operator
	Exists                Operator
	BeginsWith            Operator
	RegexMatches          Operator
}{
	Operator("NOT_LOWER_THAN_OR_EQUAL"),
	Operator("NOT_ENDS_WITH"),
	Operator("NOT_CONTAINS"),
	Operator("NOT_REGEX_MATCHES"),
	Operator("NOT_EXISTS"),
	Operator("NOT_GREATER_THAN_OR_EQUAL"),
	Operator("CONTAINS"),
	Operator("IS_IP_IN_RANGE"),
	Operator("NOT_IS_IP_IN_RANGE"),
	Operator("ENDS_WITH"),
	Operator("TAG_KEY_EQUALS"),
	Operator("NOT_TAG_KEY_EQUALS"),
	Operator("GREATER_THAN"),
	Operator("EQUALS"),
	Operator("NOT_EQUALS"),
	Operator("LOWER_THAN"),
	Operator("NOT_LOWER_THAN"),
	Operator("NOT_BEGINS_WITH"),
	Operator("LOWER_THAN_OR_EQUAL"),
	Operator("NOT_GREATER_THAN"),
	Operator("GREATER_THAN_OR_EQUAL"),
	Operator("EXISTS"),
	Operator("BEGINS_WITH"),
	Operator("REGEX_MATCHES"),
}

type Normalization string

var Normalizations = struct {
	LeavetextasIs Normalization
	Touppercase   Normalization
	Tolowercase   Normalization
}{
	Normalization("Leavetextas_is"),
	Normalization("Touppercase"),
	Normalization("Tolowercase"),
}
