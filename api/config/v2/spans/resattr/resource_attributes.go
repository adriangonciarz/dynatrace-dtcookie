package resattr

import (
	"encoding/json"
	"sort"

	"github.com/dtcookie/hcl"
)

// ResourceAttributes has no documentation
type ResourceAttributes struct {
	AttributeKeys AttributeKeys `json:"attributeKeys"`
}

func (me *ResourceAttributes) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"keys": {
			Type:        hcl.TypeList,
			Description: "Attribute key allow-list",
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(AttributeKeys).Schema()},
		},
	}
}

func (me *ResourceAttributes) EnsurePredictableOrder() {
	if len(me.AttributeKeys) == 0 {
		return
	}
	conds := AttributeKeys{}
	condStrings := sort.StringSlice{}
	for _, entry := range me.AttributeKeys {
		condBytes, _ := json.Marshal(entry)
		condStrings = append(condStrings, string(condBytes))
	}
	condStrings.Sort()
	for _, condString := range condStrings {
		cond := RuleItem{}
		json.Unmarshal([]byte(condString), &cond)
		conds = append(conds, &cond)
	}
	me.AttributeKeys = conds
}

func (me *ResourceAttributes) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me.AttributeKeys) > 0 {
		me.EnsurePredictableOrder()
		marshalled, err := me.AttributeKeys.MarshalHCL()
		if err != nil {
			return nil, err
		}
		result["keys"] = []interface{}{marshalled}
	}
	return result, nil
}

func (me *ResourceAttributes) UnmarshalHCL(decoder hcl.Decoder) error {
	if _, ok := decoder.GetOk("keys.#"); ok {
		me.AttributeKeys = AttributeKeys{}
		if err := me.AttributeKeys.UnmarshalHCL(hcl.NewDecoder(decoder, "keys", 0)); err != nil {
			return err
		}
	}
	return nil
}
