package customrumjavascriptversion

import "github.com/dtcookie/hcl"

type Settings struct {
	CustomJavaScriptVersion *string `json:"customJavaScriptVersion"` // Choose custom version
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_java_script_version": {
			Type:        hcl.TypeString,
			Description: "Choose custom version",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_java_script_version": me.CustomJavaScriptVersion,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_java_script_version": &me.CustomJavaScriptVersion,
	})
}
