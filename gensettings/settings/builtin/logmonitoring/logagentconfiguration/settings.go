package logagentconfiguration

import "github.com/dtcookie/hcl"

type Settings struct {
	LAConfigContainersLogsDetectionEnabled bool   `json:"LAConfigContainersLogsDetectionEnabled"` // If set to FALSE - logs from containers won't be detected
	LAConfigDateSearchLimit_Bytes          int    `json:"LAConfigDateSearchLimit_Bytes"`          // Defines the number of characters in every log line (starting from the first chracter in the line) where the timestamp is searched.
	LAConfigDefaultTimezone                string `json:"LAConfigDefaultTimezone"`                // Default timezone for agent if more specific configurations is not defined.
	LAConfigEventLogQueryTimeout_Sec       int    `json:"LAConfigEventLogQueryTimeout_Sec"`       // Defines the maximum timeout value, in seconds, for the query extracting Windows Event Logs
	LAConfigIISDetectionEnabled            bool   `json:"LAConfigIISDetectionEnabled"`            // If set to FALSE - IIS log detection mechanism will be disabled
	LAConfigLogScannerLinuxNfsEnabled      bool   `json:"LAConfigLogScannerLinuxNfsEnabled"`      // If set to FALSE - logs on NFS won't be detected
	LAConfigMaxLgisPerEntityCount          int    `json:"LAConfigMaxLgisPerEntityCount"`          // Defines the maximum number of log group instances per entity after which, the new automatic ones wouldn't be added.
	LAConfigMinBinaryDetectionLimit_Bytes  int    `json:"LAConfigMinBinaryDetectionLimit_Bytes"`  // Defines the minimum number of bytes in log file required for binary detection.
	LAConfigMonitorOwnLogsEnabled          bool   `json:"LAConfigMonitorOwnLogsEnabled"`          // If set to FALSE - logs produced by OneAgent won't be monitored. Enabling this option may affect your DDU consumption.
	LAConfigOpenLogFilesDetectionEnabled   bool   `json:"LAConfigOpenLogFilesDetectionEnabled"`   // Enables auto-detection of log files on hosts. If set to false, logs won't be auto-detected.
	LAConfigSeverityDetectionLimit_Bytes   int    `json:"LAConfigSeverityDetectionLimit_Bytes"`   // Defines the number of characters in every log line (starting from the first character in the line) where severity is searched.
	LAConfigSeverityDetectionLinesLimit    int    `json:"LAConfigSeverityDetectionLinesLimit"`    // Defines the number of the first lines of every log entry where severity is searched.
	LAConfigSystemLogsDetectionEnabled     bool   `json:"LAConfigSystemLogsDetectionEnabled"`     // If set to FALSE - system logs detection mechanism will be disabled (Linux: syslog, message log) (Windows: System, Application, Security event logs)
	LAConfigUTCAsDefaultContainerTimezone  bool   `json:"LAConfigUTCAsDefaultContainerTimezone"`  // Defines if UTC is used as a default timezone in containers. This option is deprecated for agents in version 1.247 or newer.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"laconfig_containers_logs_detection_enabled": {
			Type:        hcl.TypeBool,
			Description: "If set to FALSE - logs from containers won't be detected",
			Optional:    true,
		},
		"laconfig_date_search_limit_bytes": {
			Type:        hcl.TypeInt,
			Description: "Defines the number of characters in every log line (starting from the first chracter in the line) where the timestamp is searched.",
			Required:    true,
		},
		"laconfig_default_timezone": {
			Type:        hcl.TypeString,
			Description: "Default timezone for agent if more specific configurations is not defined.",
			Required:    true,
		},
		"laconfig_event_log_query_timeout_sec": {
			Type:        hcl.TypeInt,
			Description: "Defines the maximum timeout value, in seconds, for the query extracting Windows Event Logs",
			Required:    true,
		},
		"laconfig_iisdetection_enabled": {
			Type:        hcl.TypeBool,
			Description: "If set to FALSE - IIS log detection mechanism will be disabled",
			Optional:    true,
		},
		"laconfig_log_scanner_linux_nfs_enabled": {
			Type:        hcl.TypeBool,
			Description: "If set to FALSE - logs on NFS won't be detected",
			Optional:    true,
		},
		"laconfig_max_lgis_per_entity_count": {
			Type:        hcl.TypeInt,
			Description: "Defines the maximum number of log group instances per entity after which, the new automatic ones wouldn't be added.",
			Required:    true,
		},
		"laconfig_min_binary_detection_limit_bytes": {
			Type:        hcl.TypeInt,
			Description: "Defines the minimum number of bytes in log file required for binary detection.",
			Required:    true,
		},
		"laconfig_monitor_own_logs_enabled": {
			Type:        hcl.TypeBool,
			Description: "If set to FALSE - logs produced by OneAgent won't be monitored. Enabling this option may affect your DDU consumption.",
			Optional:    true,
		},
		"laconfig_open_log_files_detection_enabled": {
			Type:        hcl.TypeBool,
			Description: "Enables auto-detection of log files on hosts. If set to false, logs won't be auto-detected.",
			Optional:    true,
		},
		"laconfig_severity_detection_limit_bytes": {
			Type:        hcl.TypeInt,
			Description: "Defines the number of characters in every log line (starting from the first character in the line) where severity is searched.",
			Required:    true,
		},
		"laconfig_severity_detection_lines_limit": {
			Type:        hcl.TypeInt,
			Description: "Defines the number of the first lines of every log entry where severity is searched.",
			Required:    true,
		},
		"laconfig_system_logs_detection_enabled": {
			Type:        hcl.TypeBool,
			Description: "If set to FALSE - system logs detection mechanism will be disabled (Linux: syslog, message log) (Windows: System, Application, Security event logs)",
			Optional:    true,
		},
		"laconfig_utcas_default_container_timezone": {
			Type:        hcl.TypeBool,
			Description: "Defines if UTC is used as a default timezone in containers. This option is deprecated for agents in version 1.247 or newer.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"laconfig_containers_logs_detection_enabled": me.LAConfigContainersLogsDetectionEnabled,
		"laconfig_date_search_limit_bytes":           me.LAConfigDateSearchLimit_Bytes,
		"laconfig_default_timezone":                  me.LAConfigDefaultTimezone,
		"laconfig_event_log_query_timeout_sec":       me.LAConfigEventLogQueryTimeout_Sec,
		"laconfig_iisdetection_enabled":              me.LAConfigIISDetectionEnabled,
		"laconfig_log_scanner_linux_nfs_enabled":     me.LAConfigLogScannerLinuxNfsEnabled,
		"laconfig_max_lgis_per_entity_count":         me.LAConfigMaxLgisPerEntityCount,
		"laconfig_min_binary_detection_limit_bytes":  me.LAConfigMinBinaryDetectionLimit_Bytes,
		"laconfig_monitor_own_logs_enabled":          me.LAConfigMonitorOwnLogsEnabled,
		"laconfig_open_log_files_detection_enabled":  me.LAConfigOpenLogFilesDetectionEnabled,
		"laconfig_severity_detection_limit_bytes":    me.LAConfigSeverityDetectionLimit_Bytes,
		"laconfig_severity_detection_lines_limit":    me.LAConfigSeverityDetectionLinesLimit,
		"laconfig_system_logs_detection_enabled":     me.LAConfigSystemLogsDetectionEnabled,
		"laconfig_utcas_default_container_timezone":  me.LAConfigUTCAsDefaultContainerTimezone,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"laconfig_containers_logs_detection_enabled": &me.LAConfigContainersLogsDetectionEnabled,
		"laconfig_date_search_limit_bytes":           &me.LAConfigDateSearchLimit_Bytes,
		"laconfig_default_timezone":                  &me.LAConfigDefaultTimezone,
		"laconfig_event_log_query_timeout_sec":       &me.LAConfigEventLogQueryTimeout_Sec,
		"laconfig_iisdetection_enabled":              &me.LAConfigIISDetectionEnabled,
		"laconfig_log_scanner_linux_nfs_enabled":     &me.LAConfigLogScannerLinuxNfsEnabled,
		"laconfig_max_lgis_per_entity_count":         &me.LAConfigMaxLgisPerEntityCount,
		"laconfig_min_binary_detection_limit_bytes":  &me.LAConfigMinBinaryDetectionLimit_Bytes,
		"laconfig_monitor_own_logs_enabled":          &me.LAConfigMonitorOwnLogsEnabled,
		"laconfig_open_log_files_detection_enabled":  &me.LAConfigOpenLogFilesDetectionEnabled,
		"laconfig_severity_detection_limit_bytes":    &me.LAConfigSeverityDetectionLimit_Bytes,
		"laconfig_severity_detection_lines_limit":    &me.LAConfigSeverityDetectionLinesLimit,
		"laconfig_system_logs_detection_enabled":     &me.LAConfigSystemLogsDetectionEnabled,
		"laconfig_utcas_default_container_timezone":  &me.LAConfigUTCAsDefaultContainerTimezone,
	})
}
