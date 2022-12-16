package managementzones

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AttributeConditions []*AttributeCondition

func (me *AttributeConditions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(AttributeCondition).Schema()},
		},
	}
}

func (me AttributeConditions) MarshalHCL() (map[string]interface{}, error) {
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
		result["condition"] = entries
	}
	return result, nil
}

func (me *AttributeConditions) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("condition"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(AttributeCondition)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "condition", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type AttributeCondition struct {
	CaseSensitive    bool      `json:"caseSensitive"`    // Case sensitive
	DynamicKey       string    `json:"dynamicKey"`       // Dynamic key
	DynamicKeySource string    `json:"dynamicKeySource"` // Key source
	EntityID         string    `json:"entityId"`         // Value
	EnumValue        string    `json:"enumValue"`        // Value
	IntegerValue     int       `json:"integerValue"`     // Value
	Key              Attribute `json:"key"`              // Property
	Operator         Operator  `json:"operator"`
	StringValue      string    `json:"stringValue"` // Value
	Tag              string    `json:"tag"`         // Format: `[CONTEXT]tagKey:tagValue`
}

func (me *AttributeCondition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"case_sensitive": {
			Type:        hcl.TypeBool,
			Description: "Case sensitive",
			Optional:    true,
		},
		"dynamic_key": {
			Type:        hcl.TypeString,
			Description: "Dynamic key",
			Required:    true,
		},
		"dynamic_key_source": {
			Type:        hcl.TypeString,
			Description: "Key source",
			Required:    true,
		},
		"entity_id": {
			Type:        hcl.TypeString,
			Description: "Value",
			Required:    true,
		},
		"enum_value": {
			Type:        hcl.TypeString,
			Description: "Value",
			Required:    true,
		},
		"integer_value": {
			Type:        hcl.TypeInt,
			Description: "Value",
			Required:    true,
		},
		"key": {
			Type:        hcl.TypeString,
			Description: "Property",
			Required:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"string_value": {
			Type:        hcl.TypeString,
			Description: "Value",
			Required:    true,
		},
		"tag": {
			Type:        hcl.TypeString,
			Description: "Format: `[CONTEXT]tagKey:tagValue`",
			Required:    true,
		},
	}
}

func (me *AttributeCondition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"case_sensitive":     me.CaseSensitive,
		"dynamic_key":        me.DynamicKey,
		"dynamic_key_source": me.DynamicKeySource,
		"entity_id":          me.EntityID,
		"enum_value":         me.EnumValue,
		"integer_value":      me.IntegerValue,
		"key":                me.Key,
		"operator":           me.Operator,
		"string_value":       me.StringValue,
		"tag":                me.Tag,
	})
}

func (me *AttributeCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"case_sensitive":     &me.CaseSensitive,
		"dynamic_key":        &me.DynamicKey,
		"dynamic_key_source": &me.DynamicKeySource,
		"entity_id":          &me.EntityID,
		"enum_value":         &me.EnumValue,
		"integer_value":      &me.IntegerValue,
		"key":                &me.Key,
		"operator":           &me.Operator,
		"string_value":       &me.StringValue,
		"tag":                &me.Tag,
	})
}
