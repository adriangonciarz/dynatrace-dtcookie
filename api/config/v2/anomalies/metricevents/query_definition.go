package metricevents

import "github.com/dtcookie/hcl"

type QueryDefinition struct {
	Type            Type             `json:"type"`
	MetricSelector  *string          `json:"metricSelector,omitempty"` /// To learn more, visit [Metric Selector](https://dt-url.net/metselad)
	MetricKey       string           `json:"metricKey"`
	Aggregation     *Aggregation     `json:"aggregation,omitempty"`
	QueryOffset     *int             `json:"queryOffset,omitempty"`  // Minute offset of sliding evaluation window for metrics with latency
	EntityFilter    *EntityFilter    `json:"entityFilter,omitempty"` // Use rule-based filters to define the scope this event monitors.
	DimensionFilter DimensionFilters `json:"dimensionFilter,omitempty"`
}

func (me *QueryDefinition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "",
			Required:    true,
		},
		"metric_selector": {
			Type:        hcl.TypeString,
			Description: "To learn more, visit [Metric Selector](https://dt-url.net/metselad)",
			Optional:    true,
		},
		"metric_key": {
			Type:        hcl.TypeString,
			Description: "",
			Required:    true,
		},
		"aggregation": {
			Type:        hcl.TypeString,
			Description: "",
			Optional:    true,
		},
		"query_offset": {
			Type:        hcl.TypeInt,
			Description: "Minute offset of sliding evaluation window for metrics with latency",
			Optional:    true,
		},
		"entity_filter": {
			Type:        hcl.TypeList,
			Description: "Use rule-based filters to define the scope this event monitors.",
			MaxItems:    1,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(EntityFilter).Schema()},
			Optional:    true,
		},
		"dimension_filter": {
			Type:        hcl.TypeList,
			Description: "",
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(DimensionFilters).Schema()},
			Optional:    true,
		},
	}
}

func (me *QueryDefinition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"type":             me.Type,
		"metric_selector":  me.MetricSelector,
		"metric_key":       me.MetricKey,
		"aggregation":      me.Aggregation,
		"query_offset":     me.QueryOffset,
		"entity_filter":    me.EntityFilter,
		"dimension_filter": me.DimensionFilter,
	})
}

func (me *QueryDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"type":             &me.Type,
		"metric_selector":  &me.MetricSelector,
		"metric_key":       &me.MetricKey,
		"aggregation":      &me.Aggregation,
		"query_offset":     &me.QueryOffset,
		"entity_filter":    &me.EntityFilter,
		"dimension_filter": &me.DimensionFilter,
	})
}
