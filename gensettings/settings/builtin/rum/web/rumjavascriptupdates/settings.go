package rumjavascriptupdates

import "github.com/dtcookie/hcl"

type Settings struct {
	JavascriptVersion JavascriptVersion `json:"JavascriptVersion"` // Choose version
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"javascript_version": {
			Type:        hcl.TypeString,
			Description: "Choose version",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"javascript_version": me.JavascriptVersion,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"javascript_version": &me.JavascriptVersion,
	})
}
