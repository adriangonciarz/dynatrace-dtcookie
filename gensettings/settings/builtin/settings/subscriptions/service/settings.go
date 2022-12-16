package service

import "github.com/dtcookie/hcl"

type Settings struct {
	KeyRequestNames []string `json:"keyRequestNames"` // Key request names
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key_request_names": {
			Type:        hcl.TypeSet,
			Description: "Key request names",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"key_request_names": me.KeyRequestNames,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"key_request_names": &me.KeyRequestNames,
	})
}
