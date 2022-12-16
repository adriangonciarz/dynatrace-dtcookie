package providerbreakdown

import "github.com/dtcookie/hcl"

type Settings struct {
	DomainNamePatternList   DomainNamePatternListObjects `json:"domainNamePatternList"`   // Domain name pattern
	IconUrl                 *string                      `json:"iconUrl"`                 // Specify an URL for the provider's brand icon
	ReportPublicImprovement bool                         `json:"reportPublicImprovement"` // Send the patterns of this provider to Dynatrace to help us improve 3rd-party detection.
	ResourceName            string                       `json:"resourceName"`            // Resource name
	ResourceType            ResourceType                 `json:"resourceType"`            // Resource type
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"domain_name_pattern_list": {
			Type:        hcl.TypeList,
			Description: "Domain name pattern",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DomainNamePatternListObjects).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"icon_url": {
			Type:        hcl.TypeString,
			Description: "Specify an URL for the provider's brand icon",
			Optional:    true,
		},
		"report_public_improvement": {
			Type:        hcl.TypeBool,
			Description: "Send the patterns of this provider to Dynatrace to help us improve 3rd-party detection.",
			Optional:    true,
		},
		"resource_name": {
			Type:        hcl.TypeString,
			Description: "Resource name",
			Required:    true,
		},
		"resource_type": {
			Type:        hcl.TypeString,
			Description: "Resource type",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"domain_name_pattern_list":  me.DomainNamePatternList,
		"icon_url":                  me.IconUrl,
		"report_public_improvement": me.ReportPublicImprovement,
		"resource_name":             me.ResourceName,
		"resource_type":             me.ResourceType,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"domain_name_pattern_list":  &me.DomainNamePatternList,
		"icon_url":                  &me.IconUrl,
		"report_public_improvement": &me.ReportPublicImprovement,
		"resource_name":             &me.ResourceName,
		"resource_type":             &me.ResourceType,
	})
}
