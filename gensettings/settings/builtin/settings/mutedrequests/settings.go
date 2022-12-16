package mutedrequests

import "github.com/dtcookie/hcl"

type Settings struct {
	MutedRequestNames []string `json:"mutedRequestNames"` // Muted request names
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"muted_request_names": {
			Type:        hcl.TypeSet,
			Description: "Muted request names",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"muted_request_names": me.MutedRequestNames,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"muted_request_names": &me.MutedRequestNames,
	})
}
