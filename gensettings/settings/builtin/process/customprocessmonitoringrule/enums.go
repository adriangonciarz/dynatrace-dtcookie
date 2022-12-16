package customprocessmonitoringrule

type MonitoringMode string

var MonitoringModes = struct {
	MonitoringOn  MonitoringMode
	MonitoringOff MonitoringMode
}{
	MonitoringMode("MONITORING_ON"),
	MonitoringMode("MONITORING_OFF"),
}

type ConditionOperator string

var ConditionOperators = struct {
	NotEquals   ConditionOperator
	NotExists   ConditionOperator
	NotStarts   ConditionOperator
	Ends        ConditionOperator
	Starts      ConditionOperator
	Exists      ConditionOperator
	NotEnds     ConditionOperator
	Contains    ConditionOperator
	NotContains ConditionOperator
	Equals      ConditionOperator
}{
	ConditionOperator("NOT_EQUALS"),
	ConditionOperator("NOT_EXISTS"),
	ConditionOperator("NOT_STARTS"),
	ConditionOperator("ENDS"),
	ConditionOperator("STARTS"),
	ConditionOperator("EXISTS"),
	ConditionOperator("NOT_ENDS"),
	ConditionOperator("CONTAINS"),
	ConditionOperator("NOT_CONTAINS"),
	ConditionOperator("EQUALS"),
}

type AgentItemName string

var AgentItemNames = struct {
	IisAppPool                         AgentItemName
	GlassfishInstanceName              AgentItemName
	NodejsScriptName                   AgentItemName
	PgIdCalcInputKeyLinkage            AgentItemName
	TibcoBusinessworksHome             AgentItemName
	GaeInstance                        AgentItemName
	Unknown                            AgentItemName
	GaeService                         AgentItemName
	OracleSid                          AgentItemName
	NodejsAppBaseDir                   AgentItemName
	RuxitNodeId                        AgentItemName
	SpringbootStartupClass             AgentItemName
	WebsphereClusterName               AgentItemName
	MssqlInstanceName                  AgentItemName
	GlassfishDomainName                AgentItemName
	IbmImsControl                      AgentItemName
	JavaJarFile                        AgentItemName
	ContainerName                      AgentItemName
	JbossMode                          AgentItemName
	IbmImsConnect                      AgentItemName
	CloudFoundryInstanceIndex          AgentItemName
	SoftwareagInstallRoot              AgentItemName
	CloudFoundrySpaceName              AgentItemName
	WeblogicDomainName                 AgentItemName
	ElasticSearchClusterName           AgentItemName
	CommandLineArgs                    AgentItemName
	TibcoBusinessworksAppSpaceName     AgentItemName
	ServiceName                        AgentItemName
	WeblogicClusterName                AgentItemName
	TibcoBusinessworksAppNodeName      AgentItemName
	ContainerImageName                 AgentItemName
	AwsEcsCluster                      AgentItemName
	WebsphereLibertyServerName         AgentItemName
	IbmCtgName                         AgentItemName
	JavaMainClass                      AgentItemName
	IisRoleName                        AgentItemName
	TibcoBusinessworksCeVersion        AgentItemName
	ExePath                            AgentItemName
	ElasticSearchNodeName              AgentItemName
	HybrisConfigDir                    AgentItemName
	IibExecutionGroupName              AgentItemName
	WebsphereNodeName                  AgentItemName
	WeblogicHome                       AgentItemName
	CloudFoundrySpaceId                AgentItemName
	SpringbootProfileName              AgentItemName
	KubernetesFullpodname              AgentItemName
	HybrisBinDir                       AgentItemName
	KubernetesPoduid                   AgentItemName
	DeclarativeId                      AgentItemName
	AwsRegion                          AgentItemName
	PhpCliWorkingDir                   AgentItemName
	JbossServerName                    AgentItemName
	ApacheConfigPath                   AgentItemName
	DotnetCommand                      AgentItemName
	CloudFoundryApplicationId          AgentItemName
	RuxitClusterId                     AgentItemName
	AwsEcsFamily                       AgentItemName
	IbmCicsRegion                      AgentItemName
	SoftwareagProductpropname          AgentItemName
	PhpCliScriptPath                   AgentItemName
	TibcoBusinessworksDomainName       AgentItemName
	TipcoBusinessworksPropertyFilePath AgentItemName
	IibBrokerName                      AgentItemName
	KubernetesContainername            AgentItemName
	KubernetesNamespace                AgentItemName
	JbossHome                          AgentItemName
	AspNetCoreApplicationPath          AgentItemName
	CatalinaHome                       AgentItemName
	TipcoBusinessworksPropertyFile     AgentItemName
	VarnishInstanceName                AgentItemName
	KubernetesBasepodname              AgentItemName
	NodejsAppName                      AgentItemName
	CatalinaBase                       AgentItemName
	ApacheSparkMasterIpAddress         AgentItemName
	HybrisDataDir                      AgentItemName
	AwsEcsContainername                AgentItemName
	EquinoxConfigPath                  AgentItemName
	SpringbootAppName                  AgentItemName
	ContainerId                        AgentItemName
	AwsEcrRegion                       AgentItemName
	AwsEcsRevision                     AgentItemName
	ExeName                            AgentItemName
	DotnetCommandPath                  AgentItemName
	WebsphereServerName                AgentItemName
	ColdfusionJvmConfigFile            AgentItemName
	ContainerImageVersion              AgentItemName
	JavaJarPath                        AgentItemName
	WeblogicName                       AgentItemName
	WebsphereCellName                  AgentItemName
	AwsLambdaFunctionName              AgentItemName
	IbmImsSoapGwName                   AgentItemName
	TibcoBusinessworksCeAppName        AgentItemName
	AwsEcrAccountId                    AgentItemName
	IbmImsMpr                          AgentItemName
	CloudFoundryAppName                AgentItemName
	GoogleCloudProject                 AgentItemName
}{
	AgentItemName("IIS_APP_POOL"),
	AgentItemName("GLASSFISH_INSTANCE_NAME"),
	AgentItemName("NODEJS_SCRIPT_NAME"),
	AgentItemName("PG_ID_CALC_INPUT_KEY_LINKAGE"),
	AgentItemName("TIBCO_BUSINESSWORKS_HOME"),
	AgentItemName("GAE_INSTANCE"),
	AgentItemName("UNKNOWN"),
	AgentItemName("GAE_SERVICE"),
	AgentItemName("ORACLE_SID"),
	AgentItemName("NODEJS_APP_BASE_DIR"),
	AgentItemName("RUXIT_NODE_ID"),
	AgentItemName("SPRINGBOOT_STARTUP_CLASS"),
	AgentItemName("WEBSPHERE_CLUSTER_NAME"),
	AgentItemName("MSSQL_INSTANCE_NAME"),
	AgentItemName("GLASSFISH_DOMAIN_NAME"),
	AgentItemName("IBM_IMS_CONTROL"),
	AgentItemName("JAVA_JAR_FILE"),
	AgentItemName("CONTAINER_NAME"),
	AgentItemName("JBOSS_MODE"),
	AgentItemName("IBM_IMS_CONNECT"),
	AgentItemName("CLOUD_FOUNDRY_INSTANCE_INDEX"),
	AgentItemName("SOFTWAREAG_INSTALL_ROOT"),
	AgentItemName("CLOUD_FOUNDRY_SPACE_NAME"),
	AgentItemName("WEBLOGIC_DOMAIN_NAME"),
	AgentItemName("ELASTIC_SEARCH_CLUSTER_NAME"),
	AgentItemName("COMMAND_LINE_ARGS"),
	AgentItemName("TIBCO_BUSINESSWORKS_APP_SPACE_NAME"),
	AgentItemName("SERVICE_NAME"),
	AgentItemName("WEBLOGIC_CLUSTER_NAME"),
	AgentItemName("TIBCO_BUSINESSWORKS_APP_NODE_NAME"),
	AgentItemName("CONTAINER_IMAGE_NAME"),
	AgentItemName("AWS_ECS_CLUSTER"),
	AgentItemName("WEBSPHERE_LIBERTY_SERVER_NAME"),
	AgentItemName("IBM_CTG_NAME"),
	AgentItemName("JAVA_MAIN_CLASS"),
	AgentItemName("IIS_ROLE_NAME"),
	AgentItemName("TIBCO_BUSINESSWORKS_CE_VERSION"),
	AgentItemName("EXE_PATH"),
	AgentItemName("ELASTIC_SEARCH_NODE_NAME"),
	AgentItemName("HYBRIS_CONFIG_DIR"),
	AgentItemName("IIB_EXECUTION_GROUP_NAME"),
	AgentItemName("WEBSPHERE_NODE_NAME"),
	AgentItemName("WEBLOGIC_HOME"),
	AgentItemName("CLOUD_FOUNDRY_SPACE_ID"),
	AgentItemName("SPRINGBOOT_PROFILE_NAME"),
	AgentItemName("KUBERNETES_FULLPODNAME"),
	AgentItemName("HYBRIS_BIN_DIR"),
	AgentItemName("KUBERNETES_PODUID"),
	AgentItemName("DECLARATIVE_ID"),
	AgentItemName("AWS_REGION"),
	AgentItemName("PHP_CLI_WORKING_DIR"),
	AgentItemName("JBOSS_SERVER_NAME"),
	AgentItemName("APACHE_CONFIG_PATH"),
	AgentItemName("DOTNET_COMMAND"),
	AgentItemName("CLOUD_FOUNDRY_APPLICATION_ID"),
	AgentItemName("RUXIT_CLUSTER_ID"),
	AgentItemName("AWS_ECS_FAMILY"),
	AgentItemName("IBM_CICS_REGION"),
	AgentItemName("SOFTWAREAG_PRODUCTPROPNAME"),
	AgentItemName("PHP_CLI_SCRIPT_PATH"),
	AgentItemName("TIBCO_BUSINESSWORKS_DOMAIN_NAME"),
	AgentItemName("TIPCO_BUSINESSWORKS_PROPERTY_FILE_PATH"),
	AgentItemName("IIB_BROKER_NAME"),
	AgentItemName("KUBERNETES_CONTAINERNAME"),
	AgentItemName("KUBERNETES_NAMESPACE"),
	AgentItemName("JBOSS_HOME"),
	AgentItemName("ASP_NET_CORE_APPLICATION_PATH"),
	AgentItemName("CATALINA_HOME"),
	AgentItemName("TIPCO_BUSINESSWORKS_PROPERTY_FILE"),
	AgentItemName("VARNISH_INSTANCE_NAME"),
	AgentItemName("KUBERNETES_BASEPODNAME"),
	AgentItemName("NODEJS_APP_NAME"),
	AgentItemName("CATALINA_BASE"),
	AgentItemName("APACHE_SPARK_MASTER_IP_ADDRESS"),
	AgentItemName("HYBRIS_DATA_DIR"),
	AgentItemName("AWS_ECS_CONTAINERNAME"),
	AgentItemName("EQUINOX_CONFIG_PATH"),
	AgentItemName("SPRINGBOOT_APP_NAME"),
	AgentItemName("CONTAINER_ID"),
	AgentItemName("AWS_ECR_REGION"),
	AgentItemName("AWS_ECS_REVISION"),
	AgentItemName("EXE_NAME"),
	AgentItemName("DOTNET_COMMAND_PATH"),
	AgentItemName("WEBSPHERE_SERVER_NAME"),
	AgentItemName("COLDFUSION_JVM_CONFIG_FILE"),
	AgentItemName("CONTAINER_IMAGE_VERSION"),
	AgentItemName("JAVA_JAR_PATH"),
	AgentItemName("WEBLOGIC_NAME"),
	AgentItemName("WEBSPHERE_CELL_NAME"),
	AgentItemName("AWS_LAMBDA_FUNCTION_NAME"),
	AgentItemName("IBM_IMS_SOAP_GW_NAME"),
	AgentItemName("TIBCO_BUSINESSWORKS_CE_APP_NAME"),
	AgentItemName("AWS_ECR_ACCOUNT_ID"),
	AgentItemName("IBM_IMS_MPR"),
	AgentItemName("CLOUD_FOUNDRY_APP_NAME"),
	AgentItemName("GOOGLE_CLOUD_PROJECT"),
}
