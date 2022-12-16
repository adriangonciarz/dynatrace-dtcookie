package local

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled            bool               `json:"enabled"`            // Enable Extension Execution Controller
	IngestActive       bool               `json:"ingestActive"`       // Enable local PIPE/HTTP metric and Log Ingest API
	PerformanceProfile PerformanceProfile `json:"performanceProfile"` // Select performance profile for Extension Execution Controller
	StatsdActive       bool               `json:"statsdActive"`       // Enable Dynatrace StatsD
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enable Extension Execution Controller",
			Optional:    true,
		},
		"ingest_active": {
			Type:        hcl.TypeBool,
			Description: "Enable local PIPE/HTTP metric and Log Ingest API",
			Optional:    true,
		},
		"performance_profile": {
			Type:        hcl.TypeString,
			Description: "Select performance profile for Extension Execution Controller",
			Required:    true,
		},
		"statsd_active": {
			Type:        hcl.TypeBool,
			Description: "Enable Dynatrace StatsD",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":             me.Enabled,
		"ingest_active":       me.IngestActive,
		"performance_profile": me.PerformanceProfile,
		"statsd_active":       me.StatsdActive,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":             &me.Enabled,
		"ingest_active":       &me.IngestActive,
		"performance_profile": &me.PerformanceProfile,
		"statsd_active":       &me.StatsdActive,
	})
}
