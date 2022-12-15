package metricevents

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type EntityFilterConditions []*EntityFilterCondition // Entity filter conditions

func (me *EntityFilterConditions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "Entity filter conditions",
			Elem:        &hcl.Resource{Schema: new(EntityFilterCondition).Schema()},
		},
	}
}

func (me EntityFilterConditions) MarshalHCL() (map[string]interface{}, error) {
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

func (me *EntityFilterConditions) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("condition"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(EntityFilterCondition)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "condition", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type EntityFilterCondition struct {
	Type     EntityFilterType     `json:"type"`
	Operator EntityFilterOperator `json:"operator"`
	Value    string               `json:"value"`
}

func (me *EntityFilterCondition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "",
			Required:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "",
			Required:    true,
		},
	}
}

func (me *EntityFilterCondition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"type":     me.Type,
		"operator": me.Operator,
		"value":    me.Value,
	})
}

func (me *EntityFilterCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"type":     &me.Type,
		"operator": &me.Operator,
		"value":    &me.Value,
	})
}
