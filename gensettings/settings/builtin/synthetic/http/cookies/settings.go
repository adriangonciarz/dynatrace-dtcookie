package cookies

import "github.com/dtcookie/hcl"

type Settings struct {
	Cookies CookieEntries `json:"cookies"`
	Enabled bool          `json:"enabled"` // Set cookies
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cookies": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CookieEntries).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Set cookies",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"cookies": me.Cookies,
		"enabled": me.Enabled,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cookies": &me.Cookies,
		"enabled": &me.Enabled,
	})
}
