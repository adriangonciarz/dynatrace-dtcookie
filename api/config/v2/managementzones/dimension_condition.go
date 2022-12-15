package managementzones

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DimensionConditions []*DimensionCondition

func (me *DimensionConditions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "Dimension conditions",
			Elem:        &hcl.Resource{Schema: new(DimensionCondition).Schema()},
		},
	}
}

func (me DimensionConditions) MarshalHCL() (map[string]interface{}, error) {
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

func (me *DimensionConditions) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("condition"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(DimensionCondition)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "condition", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

// No documentation available
type DimensionCondition struct {
	ConditionType DimensionConditionType `json:"conditionType"` // Type
	Key           *string                `json:"key,omitempty"` // Key
	RuleMatcher   DimensionOperator      `json:"ruleMatcher"`   // Operator
	Value         string                 `json:"value"`         // Value
}

func (me *DimensionCondition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition_type": {
			Type:        hcl.TypeString,
			Description: "Type",
			Required:    true,
		},
		"key": {
			Type:        hcl.TypeString,
			Description: "Key",
			Optional:    true,
		},
		"rule_matcher": {
			Type:        hcl.TypeString,
			Description: "Operator",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "Value",
			Required:    true,
		},
	}
}

func (me *DimensionCondition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"condition_type": me.ConditionType,
		"key":            me.Key,
		"rule_matcher":   me.RuleMatcher,
		"value":          me.Value,
	})
}

func (me *DimensionCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"condition_type": &me.ConditionType,
		"key":            &me.Key,
		"rule_matcher":   &me.RuleMatcher,
		"value":          &me.Value,
	})
}
