package tokensettings

import "github.com/dtcookie/hcl"

type Settings struct {
	NewDynatraceTokenFormatEnabled bool `json:"newDynatraceTokenFormatEnabled"` // Check out this [blog post](http://www.dynatrace.com/blog/further-increased-security-of-your-api-tokens-by-automating-token-protection/) to find out more about the new Dynatrace API token format.
	PatEnabled                     bool `json:"patEnabled"`                     // Allow users of this environment to generate personal access tokens based on user permissions. \n Note that existing personal access tokens will become unusable while this setting is disabled.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"new_dynatrace_token_format_enabled": {
			Type:        hcl.TypeBool,
			Description: "Check out this [blog post](http://www.dynatrace.com/blog/further-increased-security-of-your-api-tokens-by-automating-token-protection/) to find out more about the new Dynatrace API token format.",
			Optional:    true,
		},
		"pat_enabled": {
			Type:        hcl.TypeBool,
			Description: "Allow users of this environment to generate personal access tokens based on user permissions. \n Note that existing personal access tokens will become unusable while this setting is disabled.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"new_dynatrace_token_format_enabled": me.NewDynatraceTokenFormatEnabled,
		"pat_enabled":                        me.PatEnabled,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"new_dynatrace_token_format_enabled": &me.NewDynatraceTokenFormatEnabled,
		"pat_enabled":                        &me.PatEnabled,
	})
}
