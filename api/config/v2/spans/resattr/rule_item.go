package resattr

import (
	"github.com/dtcookie/hcl"
)

type RuleItem struct {
	Enabled      bool        `json:"enabled"`
	AttributeKey string      `json:"attributeKey"`
	Masking      MaskingType `json:"masking"`
}

func (me *RuleItem) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "If this is true, the value of the specified key is stored.",
			Required:    true,
		},
		"attribute_key": {
			Type:        hcl.TypeString,
			Description: "Attribute key **service.name** is automatically captured by default",
			Required:    true,
		},
		"masking": {
			Type:        hcl.TypeString,
			Description: "Introduce more granular control over the visibility of attribute values.  \nChoose **Do not mask** to allow every user to see the actual value and use it in defining other configurations.  \nChoose **Mask entire value** to hide the whole value of this attribute from everyone who does not have 'View sensitive request data' permission. These attributes can't be used to define other configurations. \nChoose **Mask only confidential data** to apply automatic masking strategies to your data. These strategies include, for example, credit card numbers, IBAN, IPs, email-addresses, etc. It may not be possible to recognize all sensitive data so please always make sure to verify that sensitive data is actually masked. If sensitive data is not recognized, please use **Mask entire value** instead. Users with 'View sensitive request data' permission will be able to see the entire value, others only the non-sensitive parts. These attributes can't be used to define other configurations.",
			Required:    true,
		},
	}
}

func (me *RuleItem) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["enabled"] = me.Enabled
	result["attribute_key"] = me.AttributeKey
	result["masking"] = string(me.Masking)

	return result, nil
}

func (me *RuleItem) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("enabled"); ok {
		me.Enabled = value.(bool)
	}
	if value, ok := decoder.GetOk("attribute_key"); ok {
		me.AttributeKey = value.(string)
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
