package detectionflags

import "github.com/dtcookie/hcl"

type Settings struct {
	AddNodeJsScriptName                    bool `json:"addNodeJsScriptName"`                    // In older versions, Node.js applications were distinguished based on their directory name, omitting the script name. Changing this setting may change the general handling of Node.js process groups. Leave unchanged if in doubt.
	AutoDetectCassandraClusters            bool `json:"autoDetectCassandraClusters"`            // Enabling this flag will detect separate Cassandra process groups based on the configured Cassandra cluster name.
	AutoDetectSpringBoot                   bool `json:"autoDetectSpringBoot"`                   // Enabling this flag will detect Spring Boot process groups based on command line and applications' configuration files.
	AutoDetectTibcoContainerEditionEngines bool `json:"autoDetectTibcoContainerEditionEngines"` // Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.
	AutoDetectTibcoEngines                 bool `json:"autoDetectTibcoEngines"`                 // Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.
	AutoDetectWebMethodsIntegrationServer  bool `json:"autoDetectWebMethodsIntegrationServer"`  // Enabling this flag will detect webMethods Integration Server including specific properties like install root and product name.
	AutoDetectWebSphereLibertyApplication  bool `json:"autoDetectWebSphereLibertyApplication"`  // Enabling this flag will detect separate WebSphere Liberty process groups based on java command line.
	GroupIBMMQbyInstanceName               bool `json:"groupIBMMQbyInstanceName"`               // Enable to group and separately analyze the processes of each IBM MQ Queue manager instance. Each process group receives a unique name based on the queue manager instance name.
	IdentifyJbossServerBySystemProperty    bool `json:"identifyJbossServerBySystemProperty"`    // Enabling this flag will detect the JBoss server name from the system property jboss.server.name=<server-name>, only if -D[Server:<server-name>] is not set.
	IgnoreUniqueIdentifiers                bool `json:"ignoreUniqueIdentifiers"`                // To determine the unique identity of each detected process, and to generate a unique name for each detected process, Dynatrace evaluates the name of the directory that each process binary is contained within. For application containers like Tomcat and JBoss, Dynatrace evaluates important directories like CATALINA_HOME and JBOSS_HOME for this information. In some automated deployment scenarios such directory names are updated automatically with new version numbers, build numbers, dates, or GUIDs. Enable this setting to ensure that automated directory name changes don't result in Dynatrace registering pre-existing processes as new processes.
	ShortLivedProcessesMonitoring          bool `json:"shortLivedProcessesMonitoring"`          // Enable to monitor CPU and memory usage of short lived processes, otherwise being lost by traditional monitoring. Disabling this flag blocks passing data to cluster only, it does not stop data collection and has no effect on performance.
	SplitOracleDatabasePG                  bool `json:"splitOracleDatabasePG"`                  // Enable to group and separately analyze the processes of each Oracle DB. Each process group receives a unique name based on the Oracle DB SID.
	SplitOracleListenerPG                  bool `json:"splitOracleListenerPG"`                  // Enable to group and separately analyze the processes of each Oracle Listener. Each process group receives a unique name based on the Oracle Listener name.
	UseCatalinaBase                        bool `json:"useCatalinaBase"`                        // By default, Tomcat clusters are identified and named based on the CATALINA_HOME directory name. This setting results in the use of the CATALINA_BASE directory name to identify multiple Tomcat nodes within each Tomcat cluster. If this setting is not enabled, each CATALINA_HOME+CATALINA_BASE combination will be considered a separate Tomcat cluster. In other words, Tomcat clusters can't have multiple nodes on a single host.
	UseDockerContainerName                 bool `json:"useDockerContainerName"`                 // By default, Dynatrace uses image names as identifiers for individual process groups, with one process-group instance per host. Normally Docker container names can't serve as stable identifiers of process group instances because they are variable and auto-generated. You can however manually assign proper container names to their Docker instances. Such manually-assigned container names can serve as reliable process-group instance identifiers. This flag instructs Dynatrace to use Docker-provided names to distinguish between multiple instances of the same image. If this flag is not applied and you run multiple containers of the same image on the same host, the resulting processes will be consolidated into a single process view. Use this flag with caution!
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"add_node_js_script_name": {
			Type:        hcl.TypeBool,
			Description: "In older versions, Node.js applications were distinguished based on their directory name, omitting the script name. Changing this setting may change the general handling of Node.js process groups. Leave unchanged if in doubt.",
			Optional:    true,
		},
		"auto_detect_cassandra_clusters": {
			Type:        hcl.TypeBool,
			Description: "Enabling this flag will detect separate Cassandra process groups based on the configured Cassandra cluster name.",
			Optional:    true,
		},
		"auto_detect_spring_boot": {
			Type:        hcl.TypeBool,
			Description: "Enabling this flag will detect Spring Boot process groups based on command line and applications' configuration files.",
			Optional:    true,
		},
		"auto_detect_tibco_container_edition_engines": {
			Type:        hcl.TypeBool,
			Description: "Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.",
			Optional:    true,
		},
		"auto_detect_tibco_engines": {
			Type:        hcl.TypeBool,
			Description: "Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.",
			Optional:    true,
		},
		"auto_detect_web_methods_integration_server": {
			Type:        hcl.TypeBool,
			Description: "Enabling this flag will detect webMethods Integration Server including specific properties like install root and product name.",
			Optional:    true,
		},
		"auto_detect_web_sphere_liberty_application": {
			Type:        hcl.TypeBool,
			Description: "Enabling this flag will detect separate WebSphere Liberty process groups based on java command line.",
			Optional:    true,
		},
		"group_ibmmqby_instance_name": {
			Type:        hcl.TypeBool,
			Description: "Enable to group and separately analyze the processes of each IBM MQ Queue manager instance. Each process group receives a unique name based on the queue manager instance name.",
			Optional:    true,
		},
		"identify_jboss_server_by_system_property": {
			Type:        hcl.TypeBool,
			Description: "Enabling this flag will detect the JBoss server name from the system property jboss.server.name=<server-name>, only if -D[Server:<server-name>] is not set.",
			Optional:    true,
		},
		"ignore_unique_identifiers": {
			Type:        hcl.TypeBool,
			Description: "To determine the unique identity of each detected process, and to generate a unique name for each detected process, Dynatrace evaluates the name of the directory that each process binary is contained within. For application containers like Tomcat and JBoss, Dynatrace evaluates important directories like CATALINA_HOME and JBOSS_HOME for this information. In some automated deployment scenarios such directory names are updated automatically with new version numbers, build numbers, dates, or GUIDs. Enable this setting to ensure that automated directory name changes don't result in Dynatrace registering pre-existing processes as new processes.",
			Optional:    true,
		},
		"short_lived_processes_monitoring": {
			Type:        hcl.TypeBool,
			Description: "Enable to monitor CPU and memory usage of short lived processes, otherwise being lost by traditional monitoring. Disabling this flag blocks passing data to cluster only, it does not stop data collection and has no effect on performance.",
			Optional:    true,
		},
		"split_oracle_database_pg": {
			Type:        hcl.TypeBool,
			Description: "Enable to group and separately analyze the processes of each Oracle DB. Each process group receives a unique name based on the Oracle DB SID.",
			Optional:    true,
		},
		"split_oracle_listener_pg": {
			Type:        hcl.TypeBool,
			Description: "Enable to group and separately analyze the processes of each Oracle Listener. Each process group receives a unique name based on the Oracle Listener name.",
			Optional:    true,
		},
		"use_catalina_base": {
			Type:        hcl.TypeBool,
			Description: "By default, Tomcat clusters are identified and named based on the CATALINA_HOME directory name. This setting results in the use of the CATALINA_BASE directory name to identify multiple Tomcat nodes within each Tomcat cluster. If this setting is not enabled, each CATALINA_HOME+CATALINA_BASE combination will be considered a separate Tomcat cluster. In other words, Tomcat clusters can't have multiple nodes on a single host.",
			Optional:    true,
		},
		"use_docker_container_name": {
			Type:        hcl.TypeBool,
			Description: "By default, Dynatrace uses image names as identifiers for individual process groups, with one process-group instance per host. Normally Docker container names can't serve as stable identifiers of process group instances because they are variable and auto-generated. You can however manually assign proper container names to their Docker instances. Such manually-assigned container names can serve as reliable process-group instance identifiers. This flag instructs Dynatrace to use Docker-provided names to distinguish between multiple instances of the same image. If this flag is not applied and you run multiple containers of the same image on the same host, the resulting processes will be consolidated into a single process view. Use this flag with caution!",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"add_node_js_script_name":                     me.AddNodeJsScriptName,
		"auto_detect_cassandra_clusters":              me.AutoDetectCassandraClusters,
		"auto_detect_spring_boot":                     me.AutoDetectSpringBoot,
		"auto_detect_tibco_container_edition_engines": me.AutoDetectTibcoContainerEditionEngines,
		"auto_detect_tibco_engines":                   me.AutoDetectTibcoEngines,
		"auto_detect_web_methods_integration_server":  me.AutoDetectWebMethodsIntegrationServer,
		"auto_detect_web_sphere_liberty_application":  me.AutoDetectWebSphereLibertyApplication,
		"group_ibmmqby_instance_name":                 me.GroupIBMMQbyInstanceName,
		"identify_jboss_server_by_system_property":    me.IdentifyJbossServerBySystemProperty,
		"ignore_unique_identifiers":                   me.IgnoreUniqueIdentifiers,
		"short_lived_processes_monitoring":            me.ShortLivedProcessesMonitoring,
		"split_oracle_database_pg":                    me.SplitOracleDatabasePG,
		"split_oracle_listener_pg":                    me.SplitOracleListenerPG,
		"use_catalina_base":                           me.UseCatalinaBase,
		"use_docker_container_name":                   me.UseDockerContainerName,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"add_node_js_script_name":                     &me.AddNodeJsScriptName,
		"auto_detect_cassandra_clusters":              &me.AutoDetectCassandraClusters,
		"auto_detect_spring_boot":                     &me.AutoDetectSpringBoot,
		"auto_detect_tibco_container_edition_engines": &me.AutoDetectTibcoContainerEditionEngines,
		"auto_detect_tibco_engines":                   &me.AutoDetectTibcoEngines,
		"auto_detect_web_methods_integration_server":  &me.AutoDetectWebMethodsIntegrationServer,
		"auto_detect_web_sphere_liberty_application":  &me.AutoDetectWebSphereLibertyApplication,
		"group_ibmmqby_instance_name":                 &me.GroupIBMMQbyInstanceName,
		"identify_jboss_server_by_system_property":    &me.IdentifyJbossServerBySystemProperty,
		"ignore_unique_identifiers":                   &me.IgnoreUniqueIdentifiers,
		"short_lived_processes_monitoring":            &me.ShortLivedProcessesMonitoring,
		"split_oracle_database_pg":                    &me.SplitOracleDatabasePG,
		"split_oracle_listener_pg":                    &me.SplitOracleListenerPG,
		"use_catalina_base":                           &me.UseCatalinaBase,
		"use_docker_container_name":                   &me.UseDockerContainerName,
	})
}
