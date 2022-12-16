package metricevents

import "github.com/dtcookie/hcl"

type ModelProperties struct {
	AlertCondition    AlertCondition `json:"alertCondition"`    // Alert condition
	AlertOnNoData     bool           `json:"alertOnNoData"`     // The ability to set an alert on missing data in a metric. When enabled, missing data samples will contribute as violating samples defined in advanced model properties. We recommend to not alert on missing data for sparse timeseries as this leads to alert spam.
	DealertingSamples int            `json:"dealertingSamples"` // The number of one-minute samples within the evaluation window that must go back to normal to close the event.
	Samples           int            `json:"samples"`           // The number of one-minute samples that form the sliding evaluation window.
	SignalFluctuation float64        `json:"signalFluctuation"` // Controls how many times the signal fluctuation is added to the baseline to produce the actual threshold for alerting
	Threshold         float64        `json:"threshold"`         // Raise an event if this value is violated
	Tolerance         float64        `json:"tolerance"`         // Controls the width of the confidence band and larger values lead to a less sensitive model
	Type              ModelType      `json:"type"`              // Metric-key-based query definitions only support static thresholds.
	ViolatingSamples  int            `json:"violatingSamples"`  // The number of one-minute samples within the evaluation window that must violate to trigger an event.
}

func (me *ModelProperties) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"alert_condition": {
			Type:        hcl.TypeString,
			Description: "Alert condition",
			Required:    true,
		},
		"alert_on_no_data": {
			Type:        hcl.TypeBool,
			Description: "The ability to set an alert on missing data in a metric. When enabled, missing data samples will contribute as violating samples defined in advanced model properties. We recommend to not alert on missing data for sparse timeseries as this leads to alert spam.",
			Optional:    true,
		},
		"dealerting_samples": {
			Type:        hcl.TypeInt,
			Description: "The number of one-minute samples within the evaluation window that must go back to normal to close the event.",
			Required:    true,
		},
		"samples": {
			Type:        hcl.TypeInt,
			Description: "The number of one-minute samples that form the sliding evaluation window.",
			Required:    true,
		},
		"signal_fluctuation": {
			Type:        hcl.TypeFloat,
			Description: "Controls how many times the signal fluctuation is added to the baseline to produce the actual threshold for alerting",
			Required:    true,
		},
		"threshold": {
			Type:        hcl.TypeFloat,
			Description: "Raise an event if this value is violated",
			Required:    true,
		},
		"tolerance": {
			Type:        hcl.TypeFloat,
			Description: "Controls the width of the confidence band and larger values lead to a less sensitive model",
			Required:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "Metric-key-based query definitions only support static thresholds.",
			Required:    true,
		},
		"violating_samples": {
			Type:        hcl.TypeInt,
			Description: "The number of one-minute samples within the evaluation window that must violate to trigger an event.",
			Required:    true,
		},
	}
}

func (me *ModelProperties) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"alert_condition":    me.AlertCondition,
		"alert_on_no_data":   me.AlertOnNoData,
		"dealerting_samples": me.DealertingSamples,
		"samples":            me.Samples,
		"signal_fluctuation": me.SignalFluctuation,
		"threshold":          me.Threshold,
		"tolerance":          me.Tolerance,
		"type":               me.Type,
		"violating_samples":  me.ViolatingSamples,
	})
}

func (me *ModelProperties) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"alert_condition":    &me.AlertCondition,
		"alert_on_no_data":   &me.AlertOnNoData,
		"dealerting_samples": &me.DealertingSamples,
		"samples":            &me.Samples,
		"signal_fluctuation": &me.SignalFluctuation,
		"threshold":          &me.Threshold,
		"tolerance":          &me.Tolerance,
		"type":               &me.Type,
		"violating_samples":  &me.ViolatingSamples,
	})
}
