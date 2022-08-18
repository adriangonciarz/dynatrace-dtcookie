package applicationdetectionrules

import (
	"github.com/dtcookie/hcl"
)

type FilterConfig struct {
	Pattern                string `json:"pattern"`
	ApplicationMatchType   string `json:"applicationMatchType"`
	ApplicationMatchTarget string `json:"applicationMatchTarget"`
}

func (me *FilterConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"pattern": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The value to look for with the application detection rule",
		},
		"application_match_type": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The operator used for matching the application detection rule",
		},
		"application_match_target": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "Where to look for the pattern value - domain or URL",
		},
	}
}

func (me *FilterConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"pattern":                  me.Pattern,
		"application_match_type":   me.ApplicationMatchType,
		"application_match_target": me.ApplicationMatchTarget,
	})
}

func (me *FilterConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"pattern":                  &me.Pattern,
		"application_match_type":   &me.ApplicationMatchType,
		"application_match_target": &me.ApplicationMatchTarget,
	})
}
