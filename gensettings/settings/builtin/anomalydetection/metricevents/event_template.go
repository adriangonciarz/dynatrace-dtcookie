package metricevents

import "github.com/dtcookie/hcl"

type EventTemplate struct {
	DavisMerge  bool          `json:"davisMerge"`  // Davis® AI will try to merge this event into existing problems, otherwise a new problem will always be created.
	Description string        `json:"description"` // The description of the event to trigger.
	EventType   EventTypeEnum `json:"eventType"`   // The event type to trigger.
	Metadata    MetadataItems `json:"metadata"`    // Set of additional key-value properties to be attached to the triggered event.
	Title       string        `json:"title"`       // The title of the event to trigger.
}

func (me *EventTemplate) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"davis_merge": {
			Type:        hcl.TypeBool,
			Description: "Davis® AI will try to merge this event into existing problems, otherwise a new problem will always be created.",
			Optional:    true,
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
		"metadata": {
			Type:        hcl.TypeList,
			Description: "Set of additional key-value properties to be attached to the triggered event.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MetadataItems).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"title": {
			Type:        hcl.TypeString,
			Description: "The title of the event to trigger.",
			Required:    true,
		},
	}
}

func (me *EventTemplate) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"davis_merge": me.DavisMerge,
		"description": me.Description,
		"event_type":  me.EventType,
		"metadata":    me.Metadata,
		"title":       me.Title,
	})
}

func (me *EventTemplate) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"davis_merge": &me.DavisMerge,
		"description": &me.Description,
		"event_type":  &me.EventType,
		"metadata":    &me.Metadata,
		"title":       &me.Title,
	})
}
