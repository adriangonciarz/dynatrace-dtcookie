package profile

import "github.com/dtcookie/hcl"

type PredefinedEventFilter struct {
	EventType EventType `json:"eventType"` // Filter problems by a Dynatrace event type
	Negate    bool      `json:"negate"`
}

func (me *PredefinedEventFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"event_type": {
			Type:        hcl.TypeString,
			Description: "Filter problems by a Dynatrace event type",
			Required:    true,
		},
		"negate": {
			Type:        hcl.TypeBool,
			Description: "no documentation available",
			Optional:    true,
		},
	}
}

func (me *PredefinedEventFilter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"event_type": me.EventType,
		"negate":     me.Negate,
	})
}

func (me *PredefinedEventFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"event_type": &me.EventType,
		"negate":     &me.Negate,
	})
}
