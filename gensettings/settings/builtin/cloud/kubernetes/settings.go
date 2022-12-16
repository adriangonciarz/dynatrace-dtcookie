package kubernetes

import "github.com/dtcookie/hcl"

type Settings struct {
	ActiveGateGroup                 *string           `json:"activeGateGroup"`                 // ActiveGate Group
	AuthToken                       string            `json:"authToken"`                       // Create a bearer token for [Kubernetes](https://dt-url.net/og43szq \"Kubernetes\") or [OpenShift](https://dt-url.net/7l43xtp \"OpenShift\").
	CertificateCheckEnabled         bool              `json:"certificateCheckEnabled"`         // Require valid certificates for communication with API server (recommended)
	CloudApplicationPipelineEnabled bool              `json:"cloudApplicationPipelineEnabled"` // Monitor Kubernetes namespaces, services, workloads, and pods
	ClusterID                       string            `json:"clusterId"`                       // UUID of the kube-system namespace
	ClusterIdEnabled                bool              `json:"clusterIdEnabled"`                // This is required for monitoring persistent volume claims. For more information on local Kubernetes API monitoring, see the [documentation](https://dt-url.net/6q62uep).
	Enabled                         bool              `json:"enabled"`                         // Enabled
	EndpointUrl                     string            `json:"endpointUrl"`                     // Get the API URL for [Kubernetes](https://dt-url.net/kz23snj \"Kubernetes\") or [OpenShift](https://dt-url.net/d623xgw \"OpenShift\").
	EventPatterns                   EventComplexTypes `json:"eventPatterns"`                   // Define Kubernetes event filters to ingest events into your environment. For more details, see the [documentation](https://dt-url.net/2201p0u).
	EventProcessingActive           bool              `json:"eventProcessingActive"`           // All events are monitored by default unless event filters are specified.\n\nKubernetes events are subject to Davis data units (DDU) licensing.\nSee [DDUs for events](https://dt-url.net/5n03vcu) for details.
	FilterEvents                    bool              `json:"filterEvents"`                    // Include only events specified by Events Field Selectors
	HostnameVerificationEnabled     bool              `json:"hostnameVerificationEnabled"`     // Verify hostname in certificate against Kubernetes API URL
	IncludeAllFdiEvents             bool              `json:"includeAllFdiEvents"`             // For a list of included events, see the [documentation](https://dt-url.net/l61d02no).
	Label                           string            `json:"label"`                           // Renaming the cluster breaks configurations that are based on its name (e.g., management zones, and alerting).
	OpenMetricsPipelineEnabled      bool              `json:"openMetricsPipelineEnabled"`      // For annotation guidance, see the [documentation](https://dt-url.net/g42i0ppw).
	PvcMonitoringEnabled            bool              `json:"pvcMonitoringEnabled"`            // To enable dashboards and alerts, add the [Kubernetes persistent volume claims](ui/hub/ext/com.dynatrace.extension.kubernetes-pvc) extension to your environment.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"active_gate_group": {
			Type:        hcl.TypeString,
			Description: "ActiveGate Group",
			Optional:    true,
		},
		"auth_token": {
			Type:        hcl.TypeString,
			Description: "Create a bearer token for [Kubernetes](https://dt-url.net/og43szq \"Kubernetes\") or [OpenShift](https://dt-url.net/7l43xtp \"OpenShift\").",
			Required:    true,
		},
		"certificate_check_enabled": {
			Type:        hcl.TypeBool,
			Description: "Require valid certificates for communication with API server (recommended)",
			Optional:    true,
		},
		"cloud_application_pipeline_enabled": {
			Type:        hcl.TypeBool,
			Description: "Monitor Kubernetes namespaces, services, workloads, and pods",
			Optional:    true,
		},
		"cluster_id": {
			Type:        hcl.TypeString,
			Description: "UUID of the kube-system namespace",
			Required:    true,
		},
		"cluster_id_enabled": {
			Type:        hcl.TypeBool,
			Description: "This is required for monitoring persistent volume claims. For more information on local Kubernetes API monitoring, see the [documentation](https://dt-url.net/6q62uep).",
			Optional:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"endpoint_url": {
			Type:        hcl.TypeString,
			Description: "Get the API URL for [Kubernetes](https://dt-url.net/kz23snj \"Kubernetes\") or [OpenShift](https://dt-url.net/d623xgw \"OpenShift\").",
			Required:    true,
		},
		"event_patterns": {
			Type:        hcl.TypeList,
			Description: "Define Kubernetes event filters to ingest events into your environment. For more details, see the [documentation](https://dt-url.net/2201p0u).",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EventComplexTypes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"event_processing_active": {
			Type:        hcl.TypeBool,
			Description: "All events are monitored by default unless event filters are specified.\n\nKubernetes events are subject to Davis data units (DDU) licensing.\nSee [DDUs for events](https://dt-url.net/5n03vcu) for details.",
			Optional:    true,
		},
		"filter_events": {
			Type:        hcl.TypeBool,
			Description: "Include only events specified by Events Field Selectors",
			Optional:    true,
		},
		"hostname_verification_enabled": {
			Type:        hcl.TypeBool,
			Description: "Verify hostname in certificate against Kubernetes API URL",
			Optional:    true,
		},
		"include_all_fdi_events": {
			Type:        hcl.TypeBool,
			Description: "For a list of included events, see the [documentation](https://dt-url.net/l61d02no).",
			Optional:    true,
		},
		"label": {
			Type:        hcl.TypeString,
			Description: "Renaming the cluster breaks configurations that are based on its name (e.g., management zones, and alerting).",
			Required:    true,
		},
		"open_metrics_pipeline_enabled": {
			Type:        hcl.TypeBool,
			Description: "For annotation guidance, see the [documentation](https://dt-url.net/g42i0ppw).",
			Optional:    true,
		},
		"pvc_monitoring_enabled": {
			Type:        hcl.TypeBool,
			Description: "To enable dashboards and alerts, add the [Kubernetes persistent volume claims](ui/hub/ext/com.dynatrace.extension.kubernetes-pvc) extension to your environment.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"active_gate_group":                  me.ActiveGateGroup,
		"auth_token":                         me.AuthToken,
		"certificate_check_enabled":          me.CertificateCheckEnabled,
		"cloud_application_pipeline_enabled": me.CloudApplicationPipelineEnabled,
		"cluster_id":                         me.ClusterID,
		"cluster_id_enabled":                 me.ClusterIdEnabled,
		"enabled":                            me.Enabled,
		"endpoint_url":                       me.EndpointUrl,
		"event_patterns":                     me.EventPatterns,
		"event_processing_active":            me.EventProcessingActive,
		"filter_events":                      me.FilterEvents,
		"hostname_verification_enabled":      me.HostnameVerificationEnabled,
		"include_all_fdi_events":             me.IncludeAllFdiEvents,
		"label":                              me.Label,
		"open_metrics_pipeline_enabled":      me.OpenMetricsPipelineEnabled,
		"pvc_monitoring_enabled":             me.PvcMonitoringEnabled,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"active_gate_group":                  &me.ActiveGateGroup,
		"auth_token":                         &me.AuthToken,
		"certificate_check_enabled":          &me.CertificateCheckEnabled,
		"cloud_application_pipeline_enabled": &me.CloudApplicationPipelineEnabled,
		"cluster_id":                         &me.ClusterID,
		"cluster_id_enabled":                 &me.ClusterIdEnabled,
		"enabled":                            &me.Enabled,
		"endpoint_url":                       &me.EndpointUrl,
		"event_patterns":                     &me.EventPatterns,
		"event_processing_active":            &me.EventProcessingActive,
		"filter_events":                      &me.FilterEvents,
		"hostname_verification_enabled":      &me.HostnameVerificationEnabled,
		"include_all_fdi_events":             &me.IncludeAllFdiEvents,
		"label":                              &me.Label,
		"open_metrics_pipeline_enabled":      &me.OpenMetricsPipelineEnabled,
		"pvc_monitoring_enabled":             &me.PvcMonitoringEnabled,
	})
}
