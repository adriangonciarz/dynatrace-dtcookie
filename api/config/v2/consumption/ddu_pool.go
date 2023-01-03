package consumption

import (
	"github.com/dtcookie/hcl"
)

// DDUPool TODO: documentation
type DDUPool struct {
	MetricsPool       DDUPoolConfig `json:"metrics"`
	LogMonitoringPool DDUPoolConfig `json:"logMonitoring"`
	ServerlessPool    DDUPoolConfig `json:"serverless"`
	EventsPool        DDUPoolConfig `json:"events"`
	TracesPool        DDUPoolConfig `json:"traces"`
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

	if me.MetricsPool.LimitEnabled {
		if err := properties.Encode("metrics", &me.MetricsPool); err != nil {
			return nil, err
		}
	}
	if me.LogMonitoringPool.LimitEnabled {
		if err := properties.Encode("log_monitoring", &me.LogMonitoringPool); err != nil {
			return nil, err
		}
	}
	if me.ServerlessPool.LimitEnabled {
		if err := properties.Encode("serverless", &me.ServerlessPool); err != nil {
			return nil, err
		}
	}
	if me.EventsPool.LimitEnabled {
		if err := properties.Encode("events", &me.EventsPool); err != nil {
			return nil, err
		}
	}
	if me.TracesPool.LimitEnabled {
		if err := properties.Encode("traces", &me.TracesPool); err != nil {
			return nil, err
		}
	}

	return properties, nil
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
