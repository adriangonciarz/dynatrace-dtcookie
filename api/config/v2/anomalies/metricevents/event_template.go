package metricevents

import "github.com/dtcookie/hcl"

type EventTemplate struct {
	Title       string          `json:"title"`              // The title of the event to trigger.
	Description string          `json:"description"`        // The description of the event to trigger.
	EventType   EventTypeEnum   `json:"eventType"`          // The event type to trigger.
	DavisMerge  bool            `json:"davisMerge"`         // Davis® AI will try to merge this event into existing problems, otherwise a new problem will always be created.
	Metadata    []*MetadataItem `json:"metadata,omitempty"` // Set of additional key-value properties to be attached to the triggered event.
}

func (me *EventTemplate) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"title": {
			Type:        hcl.TypeString,
			Description: "The title of the event to trigger.",
			Required:    true,
		},
		"description": {
			Type:        hcl.TypeString,
			Description: "The description of the event to trigger.",
			Required:    true,
		},
		"event_type": {
			Type:        hcl.TypeString,
			Description: "The event type to trigger.",
			Required:    true,
		},
		"davis_merge": {
			Type:        hcl.TypeBool,
			Description: "Davis® AI will try to merge this event into existing problems, otherwise a new problem will always be created.",
			Optional:    true,
		},
		"metadata": {
			Type:        hcl.TypeList,
			Description: "Set of additional key-value properties to be attached to the triggered event.",
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(MetadataItem).Schema()},
			Optional:    true,
		},
	}
}

func (me *EventTemplate) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"title":       me.Title,
		"description": me.Description,
		"event_type":  me.EventType,
		"davis_merge": me.DavisMerge,
		"metadata":    me.Metadata,
	})
}

func (me *EventTemplate) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"title":       &me.Title,
		"description": &me.Description,
		"event_type":  &me.EventType,
		"davis_merge": &me.DavisMerge,
		"metadata":    &me.Metadata,
	})
}
