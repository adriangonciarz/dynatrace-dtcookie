package metadata

import (
	"github.com/dtcookie/hcl"
)

type Dimensions []*Dimension

func (me *Dimensions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dimension": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(Dimension).Schema()},
		},
	}
}

func (me Dimensions) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("dimension", me)
}

func (me *Dimensions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("dimension", me)
}

type Dimension struct {
	DisplayName string `json:"displayName"` // Display name
	Key         string `json:"key"`         // Dimension key
}

func (me *Dimension) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"display_name": {
			Type:        hcl.TypeString,
			Description: "Display name",
			Optional:    true,
		},
		"key": {
			Type:        hcl.TypeString,
			Description: "Dimension key",
			Required:    true,
		},
	}
}

func (me *Dimension) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"display_name": me.DisplayName,
		"key":          me.Key,
	})
}

func (me *Dimension) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"display_name": &me.DisplayName,
		"key":          &me.Key,
	})
}
