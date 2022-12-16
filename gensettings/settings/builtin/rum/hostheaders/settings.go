package hostheaders

import "github.com/dtcookie/hcl"

type Settings struct {
	HeaderName string `json:"headerName"` // HTTP header format
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"header_name": {
			Type:        hcl.TypeString,
			Description: "HTTP header format",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"header_name": me.HeaderName,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"header_name": &me.HeaderName,
	})
}
