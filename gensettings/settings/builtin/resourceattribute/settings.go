package resourceattribute

import "github.com/dtcookie/hcl"

type Settings struct {
	AttributeKeys RuleItems `json:"attributeKeys"` // Attribute key allow-list
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"attribute_keys": {
			Type:        hcl.TypeList,
			Description: "Attribute key allow-list",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(RuleItems).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"attribute_keys": me.AttributeKeys,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute_keys": &me.AttributeKeys,
	})
}
