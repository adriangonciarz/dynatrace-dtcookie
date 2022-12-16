package metricevents

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled                 bool             `json:"enabled"`                 // Enabled
	EventEntityDimensionKey *string          `json:"eventEntityDimensionKey"` // Controls the preferred entity type used for triggered events.
	EventTemplate           *EventTemplate   `json:"eventTemplate"`           // Event template
	LegacyID                *string          `json:"legacyId"`                // Config id
	ModelProperties         *ModelProperties `json:"modelProperties"`         // Monitoring strategy
	QueryDefinition         *QueryDefinition `json:"queryDefinition"`         // Query definition
	Summary                 string           `json:"summary"`                 // The textual summary of the metric event entry
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"event_entity_dimension_key": {
			Type:        hcl.TypeString,
			Description: "Controls the preferred entity type used for triggered events.",
			Optional:    true,
		},
		"event_template": {
			Type:        hcl.TypeList,
			Description: "Event template",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EventTemplate).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"legacy_id": {
			Type:        hcl.TypeString,
			Description: "Config id",
			Optional:    true,
		},
		"model_properties": {
			Type:        hcl.TypeList,
			Description: "Monitoring strategy",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ModelProperties).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"query_definition": {
			Type:        hcl.TypeList,
			Description: "Query definition",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(QueryDefinition).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"summary": {
			Type:        hcl.TypeString,
			Description: "The textual summary of the metric event entry",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":                    me.Enabled,
		"event_entity_dimension_key": me.EventEntityDimensionKey,
		"event_template":             me.EventTemplate,
		"legacy_id":                  me.LegacyID,
		"model_properties":           me.ModelProperties,
		"query_definition":           me.QueryDefinition,
		"summary":                    me.Summary,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":                    &me.Enabled,
		"event_entity_dimension_key": &me.EventEntityDimensionKey,
		"event_template":             &me.EventTemplate,
		"legacy_id":                  &me.LegacyID,
		"model_properties":           &me.ModelProperties,
		"query_definition":           &me.QueryDefinition,
		"summary":                    &me.Summary,
	})
}
