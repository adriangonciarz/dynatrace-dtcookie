package resourcetypes

import "github.com/dtcookie/hcl"

type Settings struct {
	PrimaryResourceType   PrimaryResourceType `json:"primaryResourceType"`   // The primary type of the resource.
	RegularExpression     string              `json:"regularExpression"`     // The regular expression to detect the resource.
	SecondaryResourceType *string             `json:"secondaryResourceType"` // The secondary type of the resource.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"primary_resource_type": {
			Type:        hcl.TypeString,
			Description: "The primary type of the resource.",
			Required:    true,
		},
		"regular_expression": {
			Type:        hcl.TypeString,
			Description: "The regular expression to detect the resource.",
			Required:    true,
		},
		"secondary_resource_type": {
			Type:        hcl.TypeString,
			Description: "The secondary type of the resource.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"primary_resource_type":   me.PrimaryResourceType,
		"regular_expression":      me.RegularExpression,
		"secondary_resource_type": me.SecondaryResourceType,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"primary_resource_type":   &me.PrimaryResourceType,
		"regular_expression":      &me.RegularExpression,
		"secondary_resource_type": &me.SecondaryResourceType,
	})
}
