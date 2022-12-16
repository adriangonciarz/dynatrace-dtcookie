package externalwebrequest

import (
	"github.com/dtcookie/hcl"
)

type ReducedTransformations []*ReducedTransformation

func (me *ReducedTransformations) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"transformation": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(ReducedTransformation).Schema()},
		},
	}
}

func (me ReducedTransformations) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("transformation", me)
}

func (me *ReducedTransformations) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("transformation", me)
}

type ReducedTransformation struct {
	IncludeHexNumbers  bool                          `json:"includeHexNumbers"` // include hexadecimal numbers
	MinDigitCount      int                           `json:"minDigitCount"`     // min digit count
	Prefix             string                        `json:"prefix"`
	ReplacementValue   string                        `json:"replacementValue"` // replacement
	Suffix             string                        `json:"suffix"`
	TransformationType ContextRootTransformationType `json:"transformationType"` // Defines what kind of transformation will be applied on the original value.
}

func (me *ReducedTransformation) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"include_hex_numbers": {
			Type:        hcl.TypeBool,
			Description: "include hexadecimal numbers",
			Optional:    true,
		},
		"min_digit_count": {
			Type:        hcl.TypeInt,
			Description: "min digit count",
			Optional:    true,
		},
		"prefix": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Optional:    true,
		},
		"replacement_value": {
			Type:        hcl.TypeString,
			Description: "replacement",
			Optional:    true,
		},
		"suffix": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Optional:    true,
		},
		"transformation_type": {
			Type:        hcl.TypeString,
			Description: "Defines what kind of transformation will be applied on the original value.",
			Required:    true,
		},
	}
}

func (me *ReducedTransformation) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"include_hex_numbers": me.IncludeHexNumbers,
		"min_digit_count":     me.MinDigitCount,
		"prefix":              me.Prefix,
		"replacement_value":   me.ReplacementValue,
		"suffix":              me.Suffix,
		"transformation_type": me.TransformationType,
	})
}

func (me *ReducedTransformation) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"include_hex_numbers": &me.IncludeHexNumbers,
		"min_digit_count":     &me.MinDigitCount,
		"prefix":              &me.Prefix,
		"replacement_value":   &me.ReplacementValue,
		"suffix":              &me.Suffix,
		"transformation_type": &me.TransformationType,
	})
}
