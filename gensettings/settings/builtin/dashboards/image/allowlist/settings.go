package allowlist

import "github.com/dtcookie/hcl"

type Settings struct {
	Allowlist URLPatterns `json:"allowlist"` // List of URL pattern matchers
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"allowlist": {
			Type:        hcl.TypeList,
			Description: "List of URL pattern matchers",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(URLPatterns).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"allowlist": me.Allowlist,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"allowlist": &me.Allowlist,
	})
}
