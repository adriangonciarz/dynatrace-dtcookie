package profile

type SeverityLevel string

var SeverityLevels = struct {
	Availability          SeverityLevel
	CustomAlert           SeverityLevel
	Errors                SeverityLevel
	MonitoringUnavailable SeverityLevel
	Performance           SeverityLevel
	ResourceContention    SeverityLevel
}{
	SeverityLevel("AVAILABILITY"),
	SeverityLevel("CUSTOM_ALERT"),
	SeverityLevel("ERRORS"),
	SeverityLevel("MONITORING_UNAVAILABLE"),
	SeverityLevel("PERFORMANCE"),
	SeverityLevel("RESOURCE_CONTENTION"),
}

type TagFilterIncludeMode string

var TagFilterIncludeModes = struct {
	None       TagFilterIncludeMode
	IncludeAny TagFilterIncludeMode
	IncludeAll TagFilterIncludeMode
}{
	TagFilterIncludeMode("NONE"),
	TagFilterIncludeMode("INCLUDE_ANY"),
	TagFilterIncludeMode("INCLUDE_ALL"),
}

type ComparisonOperator string

var ComparisonOperators = struct {
	EndsWith     ComparisonOperator
	Contains     ComparisonOperator
	RegexMatches ComparisonOperator
	StringEquals ComparisonOperator
	BeginsWith   ComparisonOperator
}{
	ComparisonOperator("ENDS_WITH"),
	ComparisonOperator("CONTAINS"),
	ComparisonOperator("REGEX_MATCHES"),
	ComparisonOperator("STRING_EQUALS"),
	ComparisonOperator("BEGINS_WITH"),
}

type AlertingProfileEventFilterType string

var AlertingProfileEventFilterTypes = struct {
	Predefined AlertingProfileEventFilterType
	Custom     AlertingProfileEventFilterType
}{
	AlertingProfileEventFilterType("PREDEFINED"),
	AlertingProfileEventFilterType("CUSTOM"),
}

type EventType string

var EventTypes = struct {
	CustomAppCrashRateIncreased         EventType
	MobileApplicationErrorRateIncreased EventType
	ExternalSyntheticTestOutage         EventType
	ProcessThreadsResourceExhausted     EventType
	EsxiHostCpuSaturation               EventType
	OsiNicDroppedPacketsHigh            EventType
	ProcessHighGcActivity               EventType
	RdsHighCpu                          EventType
	ServiceSlowdown                     EventType
	OsiNicErrorsHigh                    EventType
	EsxiGuestActiveSwapWait             EventType
	SyntheticGlobalOutage               EventType
	CustomApplicationUnexpectedLowLoad  EventType
	Ec2HighCpu                          EventType
	RdsOfServiceUnavailable             EventType
	EsxiVmImpactHostCpuSaturation       EventType
	ApplicationUnexpectedHighLoad       EventType
	PgiOfServiceUnavailable             EventType
	EbsVolumeHighLatency                EventType
	ProcessCrashed                      EventType
	ServiceUnexpectedHighLoad           EventType
	SyntheticNodeOutage                 EventType
	EsxiHostDiskQueueSlow               EventType
	MobileApplicationUnexpectedLowLoad  EventType
	DatabaseConnectionFailure           EventType
	SyntheticLocalOutage                EventType
	EsxiGuestCpuLimitReached            EventType
	HttpCheckLocalOutage                EventType
	HttpCheckTestLocationSlowdown       EventType
	PgiUnavailable                      EventType
	EsxiHostDiskSlow                    EventType
	RdsLowMemory                        EventType
	PgLowInstanceCount                  EventType
	ServiceErrorRateIncreased           EventType
	SyntheticPrivateLocationOutage      EventType
	ServiceUnexpectedLowLoad            EventType
	ApplicationUnexpectedLowLoad        EventType
	CustomApplicationUnexpectedHighLoad EventType
	MobileAppCrashRateIncreased         EventType
	EsxiHostNetworkProblems             EventType
	ExternalSyntheticTestSlowdown       EventType
	OsiHighCpu                          EventType
	ApplicationSlowdown                 EventType
	CustomApplicationSlowdown           EventType
	ProcessNaHighLossRate               EventType
	EsxiHostMemorySaturation            EventType
	MonitoringUnavailable               EventType
	CustomApplicationErrorRateIncreased EventType
	HostOfServiceUnavailable            EventType
	OsiUnexpectedlyUnavailable          EventType
	EsxiHostOverloadedStorage           EventType
	RdsHighLatency                      EventType
	ProcessMemoryResourceExhausted      EventType
	OsiLowDiskSpace                     EventType
	RdsLowStorageSpace                  EventType
	OsiGracefullyShutdown               EventType
	EsxiVmImpactHostMemorySaturation    EventType
	OsiSlowDisk                         EventType
	SyntheticTestLocationSlowdown       EventType
	HttpCheckGlobalOutage               EventType
	AwsLambdaHighErrorRate              EventType
	MobileApplicationSlowdown           EventType
	OsiHighMemory                       EventType
	OsiNicUtilizationHigh               EventType
	EsxiHostDatastoreLowDiskSpace       EventType
	ElbHighBackendErrorRate             EventType
	MobileApplicationUnexpectedHighLoad EventType
	ApplicationErrorRateIncreased       EventType
	RdsRestartSequence                  EventType
	ProcessNaHighConnFailRate           EventType
	OsiDiskLowInodes                    EventType
}{
	EventType("CUSTOM_APP_CRASH_RATE_INCREASED"),
	EventType("MOBILE_APPLICATION_ERROR_RATE_INCREASED"),
	EventType("EXTERNAL_SYNTHETIC_TEST_OUTAGE"),
	EventType("PROCESS_THREADS_RESOURCE_EXHAUSTED"),
	EventType("ESXI_HOST_CPU_SATURATION"),
	EventType("OSI_NIC_DROPPED_PACKETS_HIGH"),
	EventType("PROCESS_HIGH_GC_ACTIVITY"),
	EventType("RDS_HIGH_CPU"),
	EventType("SERVICE_SLOWDOWN"),
	EventType("OSI_NIC_ERRORS_HIGH"),
	EventType("ESXI_GUEST_ACTIVE_SWAP_WAIT"),
	EventType("SYNTHETIC_GLOBAL_OUTAGE"),
	EventType("CUSTOM_APPLICATION_UNEXPECTED_LOW_LOAD"),
	EventType("EC2_HIGH_CPU"),
	EventType("RDS_OF_SERVICE_UNAVAILABLE"),
	EventType("ESXI_VM_IMPACT_HOST_CPU_SATURATION"),
	EventType("APPLICATION_UNEXPECTED_HIGH_LOAD"),
	EventType("PGI_OF_SERVICE_UNAVAILABLE"),
	EventType("EBS_VOLUME_HIGH_LATENCY"),
	EventType("PROCESS_CRASHED"),
	EventType("SERVICE_UNEXPECTED_HIGH_LOAD"),
	EventType("SYNTHETIC_NODE_OUTAGE"),
	EventType("ESXI_HOST_DISK_QUEUE_SLOW"),
	EventType("MOBILE_APPLICATION_UNEXPECTED_LOW_LOAD"),
	EventType("DATABASE_CONNECTION_FAILURE"),
	EventType("SYNTHETIC_LOCAL_OUTAGE"),
	EventType("ESXI_GUEST_CPU_LIMIT_REACHED"),
	EventType("HTTP_CHECK_LOCAL_OUTAGE"),
	EventType("HTTP_CHECK_TEST_LOCATION_SLOWDOWN"),
	EventType("PGI_UNAVAILABLE"),
	EventType("ESXI_HOST_DISK_SLOW"),
	EventType("RDS_LOW_MEMORY"),
	EventType("PG_LOW_INSTANCE_COUNT"),
	EventType("SERVICE_ERROR_RATE_INCREASED"),
	EventType("SYNTHETIC_PRIVATE_LOCATION_OUTAGE"),
	EventType("SERVICE_UNEXPECTED_LOW_LOAD"),
	EventType("APPLICATION_UNEXPECTED_LOW_LOAD"),
	EventType("CUSTOM_APPLICATION_UNEXPECTED_HIGH_LOAD"),
	EventType("MOBILE_APP_CRASH_RATE_INCREASED"),
	EventType("ESXI_HOST_NETWORK_PROBLEMS"),
	EventType("EXTERNAL_SYNTHETIC_TEST_SLOWDOWN"),
	EventType("OSI_HIGH_CPU"),
	EventType("APPLICATION_SLOWDOWN"),
	EventType("CUSTOM_APPLICATION_SLOWDOWN"),
	EventType("PROCESS_NA_HIGH_LOSS_RATE"),
	EventType("ESXI_HOST_MEMORY_SATURATION"),
	EventType("MONITORING_UNAVAILABLE"),
	EventType("CUSTOM_APPLICATION_ERROR_RATE_INCREASED"),
	EventType("HOST_OF_SERVICE_UNAVAILABLE"),
	EventType("OSI_UNEXPECTEDLY_UNAVAILABLE"),
	EventType("ESXI_HOST_OVERLOADED_STORAGE"),
	EventType("RDS_HIGH_LATENCY"),
	EventType("PROCESS_MEMORY_RESOURCE_EXHAUSTED"),
	EventType("OSI_LOW_DISK_SPACE"),
	EventType("RDS_LOW_STORAGE_SPACE"),
	EventType("OSI_GRACEFULLY_SHUTDOWN"),
	EventType("ESXI_VM_IMPACT_HOST_MEMORY_SATURATION"),
	EventType("OSI_SLOW_DISK"),
	EventType("SYNTHETIC_TEST_LOCATION_SLOWDOWN"),
	EventType("HTTP_CHECK_GLOBAL_OUTAGE"),
	EventType("AWS_LAMBDA_HIGH_ERROR_RATE"),
	EventType("MOBILE_APPLICATION_SLOWDOWN"),
	EventType("OSI_HIGH_MEMORY"),
	EventType("OSI_NIC_UTILIZATION_HIGH"),
	EventType("ESXI_HOST_DATASTORE_LOW_DISK_SPACE"),
	EventType("ELB_HIGH_BACKEND_ERROR_RATE"),
	EventType("MOBILE_APPLICATION_UNEXPECTED_HIGH_LOAD"),
	EventType("APPLICATION_ERROR_RATE_INCREASED"),
	EventType("RDS_RESTART_SEQUENCE"),
	EventType("PROCESS_NA_HIGH_CONN_FAIL_RATE"),
	EventType("OSI_DISK_LOW_INODES"),
}
