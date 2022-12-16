package privacy

import "github.com/dtcookie/hcl"

type UserTracking struct {
	PersistentCookieEnabled bool `json:"persistentCookieEnabled"` // When enabled, Dynatrace places a [persistent cookie](https://dt-url.net/313o0p4n) on all end-user devices to identify returning users.
}

func (me *UserTracking) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"persistent_cookie_enabled": {
			Type:        hcl.TypeBool,
			Description: "When enabled, Dynatrace places a [persistent cookie](https://dt-url.net/313o0p4n) on all end-user devices to identify returning users.",
			Optional:    true,
		},
	}
}

func (me *UserTracking) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"persistent_cookie_enabled": me.PersistentCookieEnabled,
	})
}

func (me *UserTracking) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"persistent_cookie_enabled": &me.PersistentCookieEnabled,
	})
}
