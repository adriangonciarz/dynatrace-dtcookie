package appdetection

import "github.com/dtcookie/hcl"

type Settings struct {
	ApplicationID string  `json:"applicationId"` // Select an existing application or create a new one.
	Matcher       Matcher `json:"matcher"`       // Matcher
	Pattern       string  `json:"pattern"`       // Pattern
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"application_id": {
			Type:        hcl.TypeString,
			Description: "Select an existing application or create a new one.",
			Required:    true,
		},
		"matcher": {
			Type:        hcl.TypeString,
			Description: "Matcher",
			Required:    true,
		},
		"pattern": {
			Type:        hcl.TypeString,
			Description: "Pattern",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"application_id": me.ApplicationID,
		"matcher":        me.Matcher,
		"pattern":        me.Pattern,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"application_id": &me.ApplicationID,
		"matcher":        &me.Matcher,
		"pattern":        &me.Pattern,
	})
}
