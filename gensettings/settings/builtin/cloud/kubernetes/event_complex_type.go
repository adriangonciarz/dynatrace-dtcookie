package kubernetes

import (
	"github.com/dtcookie/hcl"
)

type EventComplexTypes []*EventComplexType

func (me *EventComplexTypes) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"eventPattern": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(EventComplexType).Schema()},
		},
	}
}

func (me EventComplexTypes) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("eventPattern", me)
}

func (me *EventComplexTypes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("eventPattern", me)
}

type EventComplexType struct {
	Active  bool   `json:"active"`  // Activate
	Label   string `json:"label"`   // Field selector name
	Pattern string `json:"pattern"` // Field selector expression
}

func (me *EventComplexType) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"active": {
			Type:        hcl.TypeBool,
			Description: "Activate",
			Optional:    true,
		},
		"label": {
			Type:        hcl.TypeString,
			Description: "Field selector name",
			Required:    true,
		},
		"pattern": {
			Type:        hcl.TypeString,
			Description: "Field selector expression",
			Required:    true,
		},
	}
}

func (me *EventComplexType) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"active":  me.Active,
		"label":   me.Label,
		"pattern": me.Pattern,
	})
}

func (me *EventComplexType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"active":  &me.Active,
		"label":   &me.Label,
		"pattern": &me.Pattern,
	})
}
