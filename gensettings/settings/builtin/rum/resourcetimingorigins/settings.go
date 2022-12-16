package resourcetimingorigins

import "github.com/dtcookie/hcl"

type Settings struct {
	Matcher OriginMatcherType `json:"matcher"` // Matcher
	Pattern string            `json:"pattern"` // Pattern
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
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
		"matcher": me.Matcher,
		"pattern": me.Pattern,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"matcher": &me.Matcher,
		"pattern": &me.Pattern,
	})
}
