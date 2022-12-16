package processmonitoring

import "github.com/dtcookie/hcl"

type Settings struct {
	AutoMonitoring bool `json:"autoMonitoring"` // By disabling automatic deep monitoring the Dynatrace OneAgent will only deep monitor processes that are covered by a respective deep monitoring rule or where monitoring is enabled explicitly.\nDisabling only works if all installed Agents have version 1.123 or higher. \n\n With automatic monitoring enabled, you can create rules that define exceptions to automatic process detection and monitoring. With automatic monitoring disabled, you can define rules that identify specific processes that should be monitored. Rules are applied in the order listed in the custom and built-in process monitoring rules settings. This means that you can construct complex operations for fine-grain control over the processes that are monitored in your environment. For example, you might define an inclusion rule that’s followed by an exclusion rule covering the same process.\nOnce created, monitoring rules can be enabled/disabled at any time. The rules will only take effect after restart of the processes in question. Alternatively, you can disable automatic monitoring entirely and instead define \"Include\" rules for those processes you want to monitor.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"auto_monitoring": {
			Type:        hcl.TypeBool,
			Description: "By disabling automatic deep monitoring the Dynatrace OneAgent will only deep monitor processes that are covered by a respective deep monitoring rule or where monitoring is enabled explicitly.\nDisabling only works if all installed Agents have version 1.123 or higher. \n\n With automatic monitoring enabled, you can create rules that define exceptions to automatic process detection and monitoring. With automatic monitoring disabled, you can define rules that identify specific processes that should be monitored. Rules are applied in the order listed in the custom and built-in process monitoring rules settings. This means that you can construct complex operations for fine-grain control over the processes that are monitored in your environment. For example, you might define an inclusion rule that’s followed by an exclusion rule covering the same process.\nOnce created, monitoring rules can be enabled/disabled at any time. The rules will only take effect after restart of the processes in question. Alternatively, you can disable automatic monitoring entirely and instead define \"Include\" rules for those processes you want to monitor.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"auto_monitoring": me.AutoMonitoring,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"auto_monitoring": &me.AutoMonitoring,
	})
}
