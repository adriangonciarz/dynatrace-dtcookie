package parameters

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CustomErrorRules []*CustomErrorRule

func (me *CustomErrorRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"customErrorRule": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(CustomErrorRule).Schema()},
		},
	}
}

func (me CustomErrorRules) MarshalHCL() (map[string]interface{}, error) {
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
		result["customErrorRule"] = entries
	}
	return result, nil
}

func (me *CustomErrorRules) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("customErrorRule"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(CustomErrorRule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "customErrorRule", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type CustomErrorRule struct {
	Condition        *CompareOperation `json:"condition"`        // Request attribute condition
	RequestAttribute string            `json:"requestAttribute"` // Request attribute
}

func (me *CustomErrorRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeList,
			Description: "Request attribute condition",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CompareOperation).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"request_attribute": {
			Type:        hcl.TypeString,
			Description: "Request attribute",
			Required:    true,
		},
	}
}

func (me *CustomErrorRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"condition":         me.Condition,
		"request_attribute": me.RequestAttribute,
	})
}

func (me *CustomErrorRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"condition":         &me.Condition,
		"request_attribute": &me.RequestAttribute,
	})
}
