package metricevents

import "github.com/dtcookie/hcl"

type QueryDefinition struct {
	Aggregation     Aggregation      `json:"aggregation"`
	DimensionFilter DimensionFilters `json:"dimensionFilter"` // Dimension filter
	EntityFilter    *EntityFilter    `json:"entityFilter"`    // Use rule-based filters to define the scope this event monitors.
	MetricKey       string           `json:"metricKey"`       // Metric key
	MetricSelector  string           `json:"metricSelector"`  // To learn more, visit [Metric Selector](https://dt-url.net/metselad)
	QueryOffset     int              `json:"queryOffset"`     // Minute offset of sliding evaluation window for metrics with latency
	Type            Type             `json:"type"`
}

func (me *QueryDefinition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"aggregation": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"dimension_filter": {
			Type:        hcl.TypeList,
			Description: "Dimension filter",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DimensionFilters).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"entity_filter": {
			Type:        hcl.TypeList,
			Description: "Use rule-based filters to define the scope this event monitors.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EntityFilter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"metric_key": {
			Type:        hcl.TypeString,
			Description: "Metric key",
			Required:    true,
		},
		"metric_selector": {
			Type:        hcl.TypeString,
			Description: "To learn more, visit [Metric Selector](https://dt-url.net/metselad)",
			Required:    true,
		},
		"query_offset": {
			Type:        hcl.TypeInt,
			Description: "Minute offset of sliding evaluation window for metrics with latency",
			Optional:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *QueryDefinition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"aggregation":      me.Aggregation,
		"dimension_filter": me.DimensionFilter,
		"entity_filter":    me.EntityFilter,
		"metric_key":       me.MetricKey,
		"metric_selector":  me.MetricSelector,
		"query_offset":     me.QueryOffset,
		"type":             me.Type,
	})
}

func (me *QueryDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"aggregation":      &me.Aggregation,
		"dimension_filter": &me.DimensionFilter,
		"entity_filter":    &me.EntityFilter,
		"metric_key":       &me.MetricKey,
		"metric_selector":  &me.MetricSelector,
		"query_offset":     &me.QueryOffset,
		"type":             &me.Type,
	})
}
