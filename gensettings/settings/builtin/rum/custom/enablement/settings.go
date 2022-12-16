package enablement

import "github.com/dtcookie/hcl"

type Settings struct {
	Rum *Rum `json:"rum"` // Capture and analyze all user actions within your application. Enable [Real User Monitoring (RUM)](https://dt-url.net/1n2b0prq) to monitor and improve your application's performance, identify errors, and gain insight into your user's behavior and experience.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rum": {
			Type:        hcl.TypeList,
			Description: "Capture and analyze all user actions within your application. Enable [Real User Monitoring (RUM)](https://dt-url.net/1n2b0prq) to monitor and improve your application's performance, identify errors, and gain insight into your user's behavior and experience.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Rum).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"rum": me.Rum,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"rum": &me.Rum,
	})
}
