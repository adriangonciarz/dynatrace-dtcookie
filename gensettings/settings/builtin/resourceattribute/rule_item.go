package resourceattribute

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type RuleItems []*RuleItem

func (me *RuleItems) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"attributeKey": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(RuleItem).Schema()},
		},
	}
}

func (me RuleItems) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me) > 0 {
		entries := []interface{}{}
		for _, entry := range me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["attributeKey"] = entries
	}
	return result, nil
}

func (me *RuleItems) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("attributeKey"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(RuleItem)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "attributeKey", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type RuleItem struct {
	AttributeKey string      `json:"attributeKey"` // Attribute key **service.name** is automatically captured by default
	Enabled      bool        `json:"enabled"`      // If this is true, the value of the specified key is stored.
	Masking      MaskingType `json:"masking"`      // If this attribute contains confidential data, turn on masking to conceal its value from users. Introduce more granular control over the visibility of attribute values.  \nChoose **Do not mask** to allow every user to see the actual value and use it in defining other configurations.  \nChoose **Mask entire value** to hide the whole value of this attribute from everyone who does not have 'View sensitive request data' permission. These attributes can't be used to define other configurations. \nChoose **Mask only confidential data** to apply automatic masking strategies to your data. These strategies include, for example, credit card numbers, IBAN, IPs, email-addresses, etc. It may not be possible to recognize all sensitive data so please always make sure to verify that sensitive data is actually masked. If sensitive data is not recognized, please use **Mask entire value** instead. Users with 'View sensitive request data' permission will be able to see the entire value, others only the non-sensitive parts. These attributes can't be used to define other configurations.
}

func (me *RuleItem) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"attribute_key": {
			Type:        hcl.TypeString,
			Description: "Attribute key **service.name** is automatically captured by default",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "If this is true, the value of the specified key is stored.",
			Optional:    true,
		},
		"masking": {
			Type:        hcl.TypeString,
			Description: "If this attribute contains confidential data, turn on masking to conceal its value from users. Introduce more granular control over the visibility of attribute values.  \nChoose **Do not mask** to allow every user to see the actual value and use it in defining other configurations.  \nChoose **Mask entire value** to hide the whole value of this attribute from everyone who does not have 'View sensitive request data' permission. These attributes can't be used to define other configurations. \nChoose **Mask only confidential data** to apply automatic masking strategies to your data. These strategies include, for example, credit card numbers, IBAN, IPs, email-addresses, etc. It may not be possible to recognize all sensitive data so please always make sure to verify that sensitive data is actually masked. If sensitive data is not recognized, please use **Mask entire value** instead. Users with 'View sensitive request data' permission will be able to see the entire value, others only the non-sensitive parts. These attributes can't be used to define other configurations.",
			Required:    true,
		},
	}
}

func (me *RuleItem) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"attribute_key": me.AttributeKey,
		"enabled":       me.Enabled,
		"masking":       me.Masking,
	})
}

func (me *RuleItem) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute_key": &me.AttributeKey,
		"enabled":       &me.Enabled,
		"masking":       &me.Masking,
	})
}
