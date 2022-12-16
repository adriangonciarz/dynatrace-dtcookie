package logevents

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled       bool           `json:"enabled"`       // Enabled
	EventTemplate *EventTemplate `json:"eventTemplate"` // Event template
	Query         string         `json:"query"`         // Log query
	Summary       string         `json:"summary"`       // The textual summary of the log event entry
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
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
		"query": {
			Type:        hcl.TypeString,
			Description: "Log query",
			Required:    true,
		},
		"summary": {
			Type:        hcl.TypeString,
			Description: "The textual summary of the log event entry",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":        me.Enabled,
		"event_template": me.EventTemplate,
		"query":          me.Query,
		"summary":        me.Summary,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":        &me.Enabled,
		"event_template": &me.EventTemplate,
		"query":          &me.Query,
		"summary":        &me.Summary,
	})
}
