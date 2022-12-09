package resattr

import (
	"github.com/dtcookie/hcl"
)

type AttributeKeys []*RuleItem

func (me *AttributeKeys) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rule": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "Attribute key allow-list",
			Elem:        &hcl.Resource{Schema: new(RuleItem).Schema()},
		},
	}
}

func (me AttributeKeys) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("rule", me)
}

func (me *AttributeKeys) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("rule", me)
}
