package externalwebrequest

import "github.com/dtcookie/hcl"

type TransformationSet struct {
	ContributionType ContributionTypeWithOverride `json:"contributionType"` // Defines whether the original value should be used or if a transformation set should be used to override a value or transform it.
	Transformations  Transformations              `json:"transformations"`
	ValueOverride    *ValueOverride               `json:"valueOverride"` // The value to be used instead of the detected value.
}

func (me *TransformationSet) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"contribution_type": {
			Type:        hcl.TypeString,
			Description: "Defines whether the original value should be used or if a transformation set should be used to override a value or transform it.",
			Required:    true,
		},
		"transformations": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Transformations).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"value_override": {
			Type:        hcl.TypeList,
			Description: "The value to be used instead of the detected value.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ValueOverride).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *TransformationSet) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"contribution_type": me.ContributionType,
		"transformations":   me.Transformations,
		"value_override":    me.ValueOverride,
	})
}

func (me *TransformationSet) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"contribution_type": &me.ContributionType,
		"transformations":   &me.Transformations,
		"value_override":    &me.ValueOverride,
	})
}
