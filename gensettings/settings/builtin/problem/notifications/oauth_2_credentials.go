package notifications

import "github.com/dtcookie/hcl"

type OAuth2Credentials struct {
}

func (me *OAuth2Credentials) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{}
}

func (me *OAuth2Credentials) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{})
}

func (me *OAuth2Credentials) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{})
}
