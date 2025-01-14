package metricevents

import (
	"github.com/dtcookie/hcl"
)

type MetricEvents struct {
	Enabled                 bool             `json:"enabled"`                           // Enabled toggle of metric event entry
	Summary                 string           `json:"summary"`                           // The textual summary of the metric event entry
	QueryDefinition         *QueryDefinition `json:"queryDefinition"`                   // The query definition of the metric event entry
	ModelProperties         *ModelProperties `json:"modelProperties"`                   // The model properties of the metric event entry
	EventTemplate           *EventTemplate   `json:"eventTemplate"`                     // The event template of the metric event entry
	EventEntityDimensionKey *string          `json:"eventEntityDimensionKey,omitempty"` // Controls the preferred entity type used for triggered events.
	LegacyId                *string          `json:"legacyId,omitempty"`                // The legacy id of the metric event entry
}

func (me *MetricEvents) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled toggle of metric event entry",
			Optional:    true,
		},
		"summary": {
			Type:        hcl.TypeString,
			Description: "The textual summary of the metric event entry",
			Required:    true,
		},
		"query_definition": {
			Type:        hcl.TypeList,
			Description: "The query definition of the metric event entry",
			MaxItems:    1,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(QueryDefinition).Schema()},
			Required:    true,
		},
		"model_properties": {
			Type:        hcl.TypeList,
			Description: "The model properties of the metric event entry",
			MaxItems:    1,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(ModelProperties).Schema()},
			Required:    true,
		},
		"event_template": {
			Type:        hcl.TypeList,
			Description: "The event template of the metric event entry",
			MaxItems:    1,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(EventTemplate).Schema()},
			Required:    true,
		},
		"event_entity_dimension_key": {
			Type:        hcl.TypeString,
			Description: "Controls the preferred entity type used for triggered events.",
			Optional:    true,
		},
		"legacy_id": {
			Type:        hcl.TypeString,
			Description: "The legacy id of the metric event entry",
			Optional:    true,
		},
	}
}

func (me *MetricEvents) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"enabled":                    me.Enabled,
		"summary":                    me.Summary,
		"query_definition":           me.QueryDefinition,
		"model_properties":           me.ModelProperties,
		"event_template":             me.EventTemplate,
		"event_entity_dimension_key": me.EventEntityDimensionKey,
		"legacy_id":                  me.LegacyId,
	})
}

func (me *MetricEvents) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"enabled":                    &me.Enabled,
		"summary":                    &me.Summary,
		"query_definition":           &me.QueryDefinition,
		"model_properties":           &me.ModelProperties,
		"event_template":             &me.EventTemplate,
		"event_entity_dimension_key": &me.EventEntityDimensionKey,
		"legacy_id":                  &me.LegacyId,
	})
}
