package limit

import "github.com/dtcookie/hcl"

type Settings struct {
	Events        *Limit `json:"events"`        // Events
	LogMonitoring *Limit `json:"logMonitoring"` // Log Monitoring
	Metrics       *Limit `json:"metrics"`       // Metrics
	Serverless    *Limit `json:"serverless"`    // Serverless
	Traces        *Limit `json:"traces"`        // Traces
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"events": {
			Type:        hcl.TypeList,
			Description: "Events",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Limit).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"log_monitoring": {
			Type:        hcl.TypeList,
			Description: "Log Monitoring",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Limit).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"metrics": {
			Type:        hcl.TypeList,
			Description: "Metrics",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Limit).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"serverless": {
			Type:        hcl.TypeList,
			Description: "Serverless",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Limit).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"traces": {
			Type:        hcl.TypeList,
			Description: "Traces",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Limit).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"events":         me.Events,
		"log_monitoring": me.LogMonitoring,
		"metrics":        me.Metrics,
		"serverless":     me.Serverless,
		"traces":         me.Traces,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"events":         &me.Events,
		"log_monitoring": &me.LogMonitoring,
		"metrics":        &me.Metrics,
		"serverless":     &me.Serverless,
		"traces":         &me.Traces,
	})
}
