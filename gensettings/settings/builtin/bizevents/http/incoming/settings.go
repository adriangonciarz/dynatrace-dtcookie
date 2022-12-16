package incoming

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled  bool             `json:"enabled"`  // Enabled
	Event    *EventComplex    `json:"event"`    // Event meta data
	RuleName string           `json:"ruleName"` // Rule name
	Triggers MatcherComplexes `json:"triggers"` // Define conditions to trigger business events from incoming web requests. Whenever one condition applies the event gets captured.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"event": {
			Type:        hcl.TypeList,
			Description: "Event meta data",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EventComplex).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"rule_name": {
			Type:        hcl.TypeString,
			Description: "Rule name",
			Required:    true,
		},
		"triggers": {
			Type:        hcl.TypeList,
			Description: "Define conditions to trigger business events from incoming web requests. Whenever one condition applies the event gets captured.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MatcherComplexes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":   me.Enabled,
		"event":     me.Event,
		"rule_name": me.RuleName,
		"triggers":  me.Triggers,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":   &me.Enabled,
		"event":     &me.Event,
		"rule_name": &me.RuleName,
		"triggers":  &me.Triggers,
	})
}
