package txmonitoring

import "github.com/dtcookie/hcl"

type Settings struct {
	GroupCicsRegions                         bool `json:"groupCicsRegions"`                         // If enabled, CICS regions belonging to the same CICSPlex will be grouped into a single process group. If disabled, a process group will be created for each CICS region.
	GroupImsRegions                          bool `json:"groupImsRegions"`                          // If enabled, IMS regions belonging to the same subsystem will be grouped into a single process group. If disabled, a process group will be created for each IMS region.
	MonitorAllCtgProtocols                   bool `json:"monitorAllCtgProtocols"`                   // If enabled, the CICS Transaction Gateway sensor will trace all EXCI requests including those that are using the TCP/IP or SNA protocol.
	MonitorAllIncomingWebRequests            bool `json:"monitorAllIncomingWebRequests"`            // Dynatrace automatically traces incoming web requests when they are called by already-monitored services. Enable this setting to monitor all incoming web requests. We recommend enabling it only over a short period of time.
	NodeLimit                                int  `json:"nodeLimit"`                                // We recommend the default limit of 500 nodes. The value 0 means unlimited number of nodes.
	ZosCicsServiceDetectionUsesTransactionID bool `json:"zosCicsServiceDetectionUsesTransactionId"` // If enabled, a CICS service will be created for each monitored transaction ID within a process group. If disabled, a CICS service will be created for each monitored CICS region within a process group. We recommend enabling it only when the CICS regions are grouped by their CICSPlex.
	ZosImsServiceDetectionUsesTransactionID  bool `json:"zosImsServiceDetectionUsesTransactionId"`  // If enabled, an IMS service will be created for each monitored transaction ID within a process group. If disabled, an IMS service will be created for each monitored IMS region within a process group. We recommend enabling it only when the IMS regions are grouped by their subsystem.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"group_cics_regions": {
			Type:        hcl.TypeBool,
			Description: "If enabled, CICS regions belonging to the same CICSPlex will be grouped into a single process group. If disabled, a process group will be created for each CICS region.",
			Optional:    true,
		},
		"group_ims_regions": {
			Type:        hcl.TypeBool,
			Description: "If enabled, IMS regions belonging to the same subsystem will be grouped into a single process group. If disabled, a process group will be created for each IMS region.",
			Optional:    true,
		},
		"monitor_all_ctg_protocols": {
			Type:        hcl.TypeBool,
			Description: "If enabled, the CICS Transaction Gateway sensor will trace all EXCI requests including those that are using the TCP/IP or SNA protocol.",
			Optional:    true,
		},
		"monitor_all_incoming_web_requests": {
			Type:        hcl.TypeBool,
			Description: "Dynatrace automatically traces incoming web requests when they are called by already-monitored services. Enable this setting to monitor all incoming web requests. We recommend enabling it only over a short period of time.",
			Optional:    true,
		},
		"node_limit": {
			Type:        hcl.TypeInt,
			Description: "We recommend the default limit of 500 nodes. The value 0 means unlimited number of nodes.",
			Required:    true,
		},
		"zos_cics_service_detection_uses_transaction_id": {
			Type:        hcl.TypeBool,
			Description: "If enabled, a CICS service will be created for each monitored transaction ID within a process group. If disabled, a CICS service will be created for each monitored CICS region within a process group. We recommend enabling it only when the CICS regions are grouped by their CICSPlex.",
			Optional:    true,
		},
		"zos_ims_service_detection_uses_transaction_id": {
			Type:        hcl.TypeBool,
			Description: "If enabled, an IMS service will be created for each monitored transaction ID within a process group. If disabled, an IMS service will be created for each monitored IMS region within a process group. We recommend enabling it only when the IMS regions are grouped by their subsystem.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"group_cics_regions":                             me.GroupCicsRegions,
		"group_ims_regions":                              me.GroupImsRegions,
		"monitor_all_ctg_protocols":                      me.MonitorAllCtgProtocols,
		"monitor_all_incoming_web_requests":              me.MonitorAllIncomingWebRequests,
		"node_limit":                                     me.NodeLimit,
		"zos_cics_service_detection_uses_transaction_id": me.ZosCicsServiceDetectionUsesTransactionID,
		"zos_ims_service_detection_uses_transaction_id":  me.ZosImsServiceDetectionUsesTransactionID,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"group_cics_regions":                             &me.GroupCicsRegions,
		"group_ims_regions":                              &me.GroupImsRegions,
		"monitor_all_ctg_protocols":                      &me.MonitorAllCtgProtocols,
		"monitor_all_incoming_web_requests":              &me.MonitorAllIncomingWebRequests,
		"node_limit":                                     &me.NodeLimit,
		"zos_cics_service_detection_uses_transaction_id": &me.ZosCicsServiceDetectionUsesTransactionID,
		"zos_ims_service_detection_uses_transaction_id":  &me.ZosImsServiceDetectionUsesTransactionID,
	})
}
