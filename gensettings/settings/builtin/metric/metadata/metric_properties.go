package metadata

import "github.com/dtcookie/hcl"

type MetricProperties struct {
	ImpactRelevant    bool      `json:"impactRelevant"`    // Whether (true or false) the metric is relevant to a problem's impact.\n\nAn impact-relevant metric is highly dependent on other metrics and changes because an underlying root-cause metric has changed.
	Latency           int       `json:"latency"`           // The latency of the metric, in minutes. \n\n The latency is the expected reporting delay (for example, caused by constraints of cloud vendors or other third-party data sources) between the observation of a metric data point and its availability in Dynatrace. \n\nThe allowed value range is from 1 to 60 minutes.
	MaxValue          float64   `json:"maxValue"`          // The maximum allowed value of the metric.
	MinValue          float64   `json:"minValue"`          // The minimum allowed value of the metric.
	RootCauseRelevant bool      `json:"rootCauseRelevant"` // Whether (true or false) the metric is related to a root cause of a problem.\n\nA root-cause relevant metric represents a strong indicator for a faulty component.
	ValueType         ValueType `json:"valueType"`         // The type of the metric's value. You have these options:\n\nscore: A score metric is a metric where high values indicate a good situation, while low values indicate trouble. An example of such a metric is a success rate.\n\nerror: An error metric is a metric where high values indicate trouble, while low values indicate a good situation. An example of such a metric is an error count.
}

func (me *MetricProperties) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"impact_relevant": {
			Type:        hcl.TypeBool,
			Description: "Whether (true or false) the metric is relevant to a problem's impact.\n\nAn impact-relevant metric is highly dependent on other metrics and changes because an underlying root-cause metric has changed.",
			Optional:    true,
		},
		"latency": {
			Type:        hcl.TypeInt,
			Description: "The latency of the metric, in minutes. \n\n The latency is the expected reporting delay (for example, caused by constraints of cloud vendors or other third-party data sources) between the observation of a metric data point and its availability in Dynatrace. \n\nThe allowed value range is from 1 to 60 minutes.",
			Optional:    true,
		},
		"max_value": {
			Type:        hcl.TypeFloat,
			Description: "The maximum allowed value of the metric.",
			Optional:    true,
		},
		"min_value": {
			Type:        hcl.TypeFloat,
			Description: "The minimum allowed value of the metric.",
			Optional:    true,
		},
		"root_cause_relevant": {
			Type:        hcl.TypeBool,
			Description: "Whether (true or false) the metric is related to a root cause of a problem.\n\nA root-cause relevant metric represents a strong indicator for a faulty component.",
			Optional:    true,
		},
		"value_type": {
			Type:        hcl.TypeString,
			Description: "The type of the metric's value. You have these options:\n\nscore: A score metric is a metric where high values indicate a good situation, while low values indicate trouble. An example of such a metric is a success rate.\n\nerror: An error metric is a metric where high values indicate trouble, while low values indicate a good situation. An example of such a metric is an error count.",
			Required:    true,
		},
	}
}

func (me *MetricProperties) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"impact_relevant":     me.ImpactRelevant,
		"latency":             me.Latency,
		"max_value":           me.MaxValue,
		"min_value":           me.MinValue,
		"root_cause_relevant": me.RootCauseRelevant,
		"value_type":          me.ValueType,
	})
}

func (me *MetricProperties) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"impact_relevant":     &me.ImpactRelevant,
		"latency":             &me.Latency,
		"max_value":           &me.MaxValue,
		"min_value":           &me.MinValue,
		"root_cause_relevant": &me.RootCauseRelevant,
		"value_type":          &me.ValueType,
	})
}
