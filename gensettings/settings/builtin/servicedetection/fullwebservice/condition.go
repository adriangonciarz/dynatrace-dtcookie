package fullwebservice

import (
	"github.com/dtcookie/hcl"
)

type Conditions []*Condition

func (me *Conditions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(Condition).Schema()},
		},
	}
}

func (me Conditions) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("condition", me)
}

func (me *Conditions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("condition", me)
}

type Condition struct {
	Attribute            string          `json:"attribute"`            // Take the value of this attribute
	CompareOperationType string          `json:"compareOperationType"` // Apply this operation
	Framework            []FrameworkType `json:"framework"`            // Technology
	IgnoreCase           bool            `json:"ignoreCase"`           // Ignore case sensitivity for texts.
	IntValue             int             `json:"intValue"`             // Value
	IntValues            []int           `json:"intValues"`            // Values
	IpRangeFrom          string          `json:"ipRangeFrom"`          // From
	IpRangeTo            string          `json:"ipRangeTo"`            // To
	TagValues            []string        `json:"tagValues"`            // If multiple values are specified, at least one of them must match for the condition to match
	TextValues           []string        `json:"textValues"`           // If multiple values are specified, at least one of them must match for the condition to match
}

func (me *Condition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"attribute": {
			Type:        hcl.TypeString,
			Description: "Take the value of this attribute",
			Required:    true,
		},
		"compare_operation_type": {
			Type:        hcl.TypeString,
			Description: "Apply this operation",
			Required:    true,
		},
		"framework": {
			Type:        hcl.TypeSet,
			Description: "Technology",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"ignore_case": {
			Type:        hcl.TypeBool,
			Description: "Ignore case sensitivity for texts.",
			Optional:    true,
		},
		"int_value": {
			Type:        hcl.TypeInt,
			Description: "Value",
			Required:    true,
		},
		"int_values": {
			Type:        hcl.TypeSet,
			Description: "Values",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeInt},
		},
		"ip_range_from": {
			Type:        hcl.TypeString,
			Description: "From",
			Required:    true,
		},
		"ip_range_to": {
			Type:        hcl.TypeString,
			Description: "To",
			Required:    true,
		},
		"tag_values": {
			Type:        hcl.TypeSet,
			Description: "If multiple values are specified, at least one of them must match for the condition to match",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"text_values": {
			Type:        hcl.TypeSet,
			Description: "If multiple values are specified, at least one of them must match for the condition to match",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Condition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"attribute":              me.Attribute,
		"compare_operation_type": me.CompareOperationType,
		"framework":              me.Framework,
		"ignore_case":            me.IgnoreCase,
		"int_value":              me.IntValue,
		"int_values":             me.IntValues,
		"ip_range_from":          me.IpRangeFrom,
		"ip_range_to":            me.IpRangeTo,
		"tag_values":             me.TagValues,
		"text_values":            me.TextValues,
	})
}

func (me *Condition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute":              &me.Attribute,
		"compare_operation_type": &me.CompareOperationType,
		"framework":              &me.Framework,
		"ignore_case":            &me.IgnoreCase,
		"int_value":              &me.IntValue,
		"int_values":             &me.IntValues,
		"ip_range_from":          &me.IpRangeFrom,
		"ip_range_to":            &me.IpRangeTo,
		"tag_values":             &me.TagValues,
		"text_values":            &me.TextValues,
	})
}
