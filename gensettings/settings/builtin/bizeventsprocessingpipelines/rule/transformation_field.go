package rule

import (
	"github.com/dtcookie/hcl"
)

type TransformationFields []*TransformationField

func (me *TransformationFields) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"transformationField": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(TransformationField).Schema()},
		},
	}
}

func (me TransformationFields) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("transformationField", me)
}

func (me *TransformationFields) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("transformationField", me)
}

type TransformationField struct {
	Array    bool                    `json:"array"` // Is Array
	Name     string                  `json:"name"`
	Optional bool                    `json:"optional"`
	Readonly bool                    `json:"readonly"` // Read-only
	Type     TransformationFieldType `json:"type"`
}

func (me *TransformationField) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"array": {
			Type:        hcl.TypeBool,
			Description: "Is Array",
			Optional:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"optional": {
			Type:        hcl.TypeBool,
			Description: "no documentation available",
			Optional:    true,
		},
		"readonly": {
			Type:        hcl.TypeBool,
			Description: "Read-only",
			Optional:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *TransformationField) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"array":    me.Array,
		"name":     me.Name,
		"optional": me.Optional,
		"readonly": me.Readonly,
		"type":     me.Type,
	})
}

func (me *TransformationField) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"array":    &me.Array,
		"name":     &me.Name,
		"optional": &me.Optional,
		"readonly": &me.Readonly,
		"type":     &me.Type,
	})
}
