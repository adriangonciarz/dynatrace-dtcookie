package outagehandling

import "github.com/dtcookie/hcl"

type Settings struct {
	GlobalConsecutiveOutageCountThreshold int  `json:"globalConsecutiveOutageCountThreshold"` // Alert if all locations are unable to access my web application
	GlobalOutages                         bool `json:"globalOutages"`                         // Generate a problem and send an alert when the monitor is unavailable at all configured locations.
	LocalConsecutiveOutageCountThreshold  int  `json:"localConsecutiveOutageCountThreshold"`  // are unable to access my web application
	LocalLocationOutageCountThreshold     int  `json:"localLocationOutageCountThreshold"`     // Alert if at least
	LocalOutages                          bool `json:"localOutages"`                          // Generate a problem and send an alert when the monitor is unavailable for one or more consecutive runs at any location.
	RetryOnError                          bool `json:"retryOnError"`                          // When enabled, which is the default, failing monitor executions are retried immediately one time to avoid false positives and only the second result is used. When disabled, we use the first result right away.\nRequires ActiveGate version 1.207+ for private locations.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"global_consecutive_outage_count_threshold": {
			Type:        hcl.TypeInt,
			Description: "Alert if all locations are unable to access my web application",
			Required:    true,
		},
		"global_outages": {
			Type:        hcl.TypeBool,
			Description: "Generate a problem and send an alert when the monitor is unavailable at all configured locations.",
			Optional:    true,
		},
		"local_consecutive_outage_count_threshold": {
			Type:        hcl.TypeInt,
			Description: "are unable to access my web application",
			Required:    true,
		},
		"local_location_outage_count_threshold": {
			Type:        hcl.TypeInt,
			Description: "Alert if at least",
			Required:    true,
		},
		"local_outages": {
			Type:        hcl.TypeBool,
			Description: "Generate a problem and send an alert when the monitor is unavailable for one or more consecutive runs at any location.",
			Optional:    true,
		},
		"retry_on_error": {
			Type:        hcl.TypeBool,
			Description: "When enabled, which is the default, failing monitor executions are retried immediately one time to avoid false positives and only the second result is used. When disabled, we use the first result right away.\nRequires ActiveGate version 1.207+ for private locations.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"global_consecutive_outage_count_threshold": me.GlobalConsecutiveOutageCountThreshold,
		"global_outages": me.GlobalOutages,
		"local_consecutive_outage_count_threshold": me.LocalConsecutiveOutageCountThreshold,
		"local_location_outage_count_threshold":    me.LocalLocationOutageCountThreshold,
		"local_outages":                            me.LocalOutages,
		"retry_on_error":                           me.RetryOnError,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"global_consecutive_outage_count_threshold": &me.GlobalConsecutiveOutageCountThreshold,
		"global_outages": &me.GlobalOutages,
		"local_consecutive_outage_count_threshold": &me.LocalConsecutiveOutageCountThreshold,
		"local_location_outage_count_threshold":    &me.LocalLocationOutageCountThreshold,
		"local_outages":                            &me.LocalOutages,
		"retry_on_error":                           &me.RetryOnError,
	})
}
