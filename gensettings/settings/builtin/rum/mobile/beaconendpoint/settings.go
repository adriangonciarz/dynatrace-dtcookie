package beaconendpoint

import "github.com/dtcookie/hcl"

type Settings struct {
	Type BeaconEndpointType `json:"type"` // Type
	Url  string             `json:"url"`  // This must be a valid beacon endpoint URL.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "Type",
			Required:    true,
		},
		"url": {
			Type:        hcl.TypeString,
			Description: "This must be a valid beacon endpoint URL.",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"type": me.Type,
		"url":  me.Url,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"type": &me.Type,
		"url":  &me.Url,
	})
}
