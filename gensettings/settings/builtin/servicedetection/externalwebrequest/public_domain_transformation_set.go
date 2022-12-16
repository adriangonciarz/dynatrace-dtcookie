package externalwebrequest

import "github.com/dtcookie/hcl"

type PublicDomainTransformationSet struct {
	ContributionType ContributionTypeWithOverride `json:"contributionType"` // Defines whether the original value should be used or if a transformation set should be used to override a value or transform it.
	CopyFromHostName bool                         `json:"copyFromHostName"` // Use the detected host name instead of the request's domain name.
	Transformations  Transformations              `json:"transformations"`
	ValueOverride    *ValueOverride               `json:"valueOverride"` // The value to be used instead of the detected value.
}

func (me *PublicDomainTransformationSet) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"contribution_type": {
			Type:        hcl.TypeString,
			Description: "Defines whether the original value should be used or if a transformation set should be used to override a value or transform it.",
			Required:    true,
		},
		"copy_from_host_name": {
			Type:        hcl.TypeBool,
			Description: "Use the detected host name instead of the request's domain name.",
			Optional:    true,
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

func (me *PublicDomainTransformationSet) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"contribution_type":   me.ContributionType,
		"copy_from_host_name": me.CopyFromHostName,
		"transformations":     me.Transformations,
		"value_override":      me.ValueOverride,
	})
}

func (me *PublicDomainTransformationSet) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"contribution_type":   &me.ContributionType,
		"copy_from_host_name": &me.CopyFromHostName,
		"transformations":     &me.Transformations,
		"value_override":      &me.ValueOverride,
	})
}
