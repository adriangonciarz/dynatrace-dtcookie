package consumption

import (
	"github.com/dtcookie/hcl"
)

// DDUPool TODO: documentation
type DDUPool struct {
	MetricsPool       *DDUPoolConfig `json:"metrics,omitempty"`
	LogMonitoringPool *DDUPoolConfig `json:"logMonitoring,omitempty"`
	ServerlessPool    *DDUPoolConfig `json:"serverless,omitempty"`
	EventsPool        *DDUPoolConfig `json:"events,omitempty"`
	TracesPool        *DDUPoolConfig `json:"traces,omitempty"`
}

func (me *DDUPool) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"metrics": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "DDU pool settings for Metrics",
			Elem: &hcl.Resource{
				Schema: new(DDUPoolConfig).Schema(),
			},
		},
		"log_monitoring": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "DDU pool settings for Log Monitoring",
			Elem: &hcl.Resource{
				Schema: new(DDUPoolConfig).Schema(),
			},
		},
		"serverless": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "DDU pool settings for Serverless",
			Elem: &hcl.Resource{
				Schema: new(DDUPoolConfig).Schema(),
			},
		},
		"events": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "DDU pool settings for Events",
			Elem: &hcl.Resource{
				Schema: new(DDUPoolConfig).Schema(),
			},
		},
		"traces": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "DDU pool settings for Traces",
			Elem: &hcl.Resource{
				Schema: new(DDUPoolConfig).Schema(),
			},
		},
	}
}

func (me *DDUPool) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"metrics":        me.MetricsPool,
		"log_monitoring": me.LogMonitoringPool,
		"serverless":     me.ServerlessPool,
		"events":         me.EventsPool,
		"traces":         me.TracesPool,
	})
}

func (me *DDUPool) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"metrics":        &me.MetricsPool,
		"log_monitoring": &me.LogMonitoringPool,
		"serverless":     &me.ServerlessPool,
		"events":         &me.EventsPool,
		"traces":         &me.TracesPool,
	})
}
