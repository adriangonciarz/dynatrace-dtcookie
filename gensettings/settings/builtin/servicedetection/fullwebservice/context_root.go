package fullwebservice

import "github.com/dtcookie/hcl"

type ContextRoot struct {
	ContributionType ContributionType       `json:"contributionType"` // Defines whether the original value should be used or if a transformation set should be used to override a value or transform it.
	SegmentCount     int                    `json:"segmentCount"`     // The number of segments of the URL to be kept.\nThe URL is divided by slashes (/), the indexing starts with 1 at context root.\nFor example, if you specify 2 for the www.dynatrace.com/support/help/dynatrace-api/ URL, the value of support/help is used.
	Transformations  ReducedTransformations `json:"transformations"`
}

func (me *ContextRoot) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"contribution_type": {
			Type:        hcl.TypeString,
			Description: "Defines whether the original value should be used or if a transformation set should be used to override a value or transform it.",
			Required:    true,
		},
		"segment_count": {
			Type:        hcl.TypeInt,
			Description: "The number of segments of the URL to be kept.\nThe URL is divided by slashes (/), the indexing starts with 1 at context root.\nFor example, if you specify 2 for the www.dynatrace.com/support/help/dynatrace-api/ URL, the value of support/help is used.",
			Required:    true,
		},
		"transformations": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ReducedTransformations).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *ContextRoot) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"contribution_type": me.ContributionType,
		"segment_count":     me.SegmentCount,
		"transformations":   me.Transformations,
	})
}

func (me *ContextRoot) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"contribution_type": &me.ContributionType,
		"segment_count":     &me.SegmentCount,
		"transformations":   &me.Transformations,
	})
}
