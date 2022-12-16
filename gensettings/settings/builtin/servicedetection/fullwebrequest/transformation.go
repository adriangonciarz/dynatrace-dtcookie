package fullwebrequest

import (
	"github.com/dtcookie/hcl"
)

type Transformations []*Transformation

func (me *Transformations) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"transformation": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(Transformation).Schema()},
		},
	}
}

func (me Transformations) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("transformation", me)
}

func (me *Transformations) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("transformation", me)
}

type Transformation struct {
	IncludeHexNumbers  bool               `json:"includeHexNumbers"` // include hexadecimal numbers
	MinDigitCount      int                `json:"minDigitCount"`     // min digit count
	Prefix             string             `json:"prefix"`
	ReplacementValue   string             `json:"replacementValue"` // replacement
	SegmentCount       int                `json:"segmentCount"`     // How many segments should be taken.
	SelectIndex        int                `json:"selectIndex"`      // select index
	SplitDelimiter     string             `json:"splitDelimiter"`   // split by
	Suffix             string             `json:"suffix"`
	TakeFromEnd        bool               `json:"takeFromEnd"`        // take from end
	TransformationType TransformationType `json:"transformationType"` // Defines what kind of transformation will be applied on the original value.
}

func (me *Transformation) Schema() map[string]*hcl.Schema {
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
		"segment_count": {
			Type:        hcl.TypeInt,
			Description: "How many segments should be taken.",
			Required:    true,
		},
		"select_index": {
			Type:        hcl.TypeInt,
			Description: "select index",
			Optional:    true,
		},
		"split_delimiter": {
			Type:        hcl.TypeString,
			Description: "split by",
			Optional:    true,
		},
		"suffix": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Optional:    true,
		},
		"take_from_end": {
			Type:        hcl.TypeBool,
			Description: "take from end",
			Optional:    true,
		},
		"transformation_type": {
			Type:        hcl.TypeString,
			Description: "Defines what kind of transformation will be applied on the original value.",
			Required:    true,
		},
	}
}

func (me *Transformation) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"include_hex_numbers": me.IncludeHexNumbers,
		"min_digit_count":     me.MinDigitCount,
		"prefix":              me.Prefix,
		"replacement_value":   me.ReplacementValue,
		"segment_count":       me.SegmentCount,
		"select_index":        me.SelectIndex,
		"split_delimiter":     me.SplitDelimiter,
		"suffix":              me.Suffix,
		"take_from_end":       me.TakeFromEnd,
		"transformation_type": me.TransformationType,
	})
}

func (me *Transformation) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"include_hex_numbers": &me.IncludeHexNumbers,
		"min_digit_count":     &me.MinDigitCount,
		"prefix":              &me.Prefix,
		"replacement_value":   &me.ReplacementValue,
		"segment_count":       &me.SegmentCount,
		"select_index":        &me.SelectIndex,
		"split_delimiter":     &me.SplitDelimiter,
		"suffix":              &me.Suffix,
		"take_from_end":       &me.TakeFromEnd,
		"transformation_type": &me.TransformationType,
	})
}
