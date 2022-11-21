package attributes

import (
	"github.com/dtcookie/hcl"
)

// SpanAttribute has no documentation
type SpanAttribute struct {
	Key     string      `json:"key"`
	Masking MaskingType `json:"masking"`
}

func (me *SpanAttribute) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "the key of the attribute to capture",
		},
		"masking": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "granular control over the visibility of attribute values",
		},
	}
}

func (me *SpanAttribute) MarshalHCL() (map[string]interface{}, error) {
	return map[string]interface{}{
		"key":     me.Key,
		"masking": string(me.Masking),
	}, nil
}

func (me *SpanAttribute) UnmarshalHCL(decoder hcl.Decoder) error {
	if key, ok := decoder.GetOk("key"); ok {
		me.Key = key.(string)
	}
	if value, ok := decoder.GetOk("masking"); ok {
		me.Masking = MaskingType(value.(string))
	}
	return nil
}

type MaskingType string

var MaskingTypes = struct {
	NotMasked    MaskingType
	Confidential MaskingType
	EntireValue  MaskingType
}{
	MaskingType("NOT_MASKED"),
	MaskingType("MASK_ONLY_CONFIDENTIAL_DATA"),
	MaskingType("MASK_ENTIRE_VALUE"),
}
