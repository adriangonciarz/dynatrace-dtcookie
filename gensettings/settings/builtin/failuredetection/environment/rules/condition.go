package rules

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Conditions []*Condition

func (me *Conditions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(Condition).Schema()},
		},
	}
}

func (me Conditions) MarshalHCL() (map[string]interface{}, error) {
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

func (me *Conditions) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("condition"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(Condition)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "condition", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type Condition struct {
	Attribute Attributes `json:"attribute"` // The attribute to be checked.
	Predicate *Predicate `json:"predicate"` // Condition to check the attribute against
}

func (me *Condition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"attribute": {
			Type:        hcl.TypeString,
			Description: "The attribute to be checked.",
			Required:    true,
		},
		"predicate": {
			Type:        hcl.TypeList,
			Description: "Condition to check the attribute against",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Predicate).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Condition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"attribute": me.Attribute,
		"predicate": me.Predicate,
	})
}

func (me *Condition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute": &me.Attribute,
		"predicate": &me.Predicate,
	})
}
