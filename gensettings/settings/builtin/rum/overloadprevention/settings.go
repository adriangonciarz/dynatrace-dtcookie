package overloadprevention

import "github.com/dtcookie/hcl"

type Settings struct {
	OverloadPreventionLimit int `json:"overloadPreventionLimit"` // Once this limit is reached, Dynatrace [throttles the number of captured user sessions](https://dt-url.net/fm3v0p7g).
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"overload_prevention_limit": {
			Type:        hcl.TypeInt,
			Description: "Once this limit is reached, Dynatrace [throttles the number of captured user sessions](https://dt-url.net/fm3v0p7g).",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"overload_prevention_limit": me.OverloadPreventionLimit,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"overload_prevention_limit": &me.OverloadPreventionLimit,
	})
}
