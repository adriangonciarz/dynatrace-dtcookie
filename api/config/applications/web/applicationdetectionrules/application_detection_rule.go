package applicationdetectionrules

import (
	"github.com/dtcookie/hcl"
)

// ApplicationDetectionRule the configuration of an application detection rule
type ApplicationDetectionRule struct {
	ID                    *string       `json:"id,omitempty"`          // the ID of the application detection rule
	Name                  *string       `json:"name,omitempty"`        // the unique name of the Application detection rule
	Order                 *string       `json:"order,omitempty"`       // the order of the rule in the rules list
	ApplicationIdentifier string        `json:"applicationIdentifier"` // the Dynatrace entity ID of the application, for example APPLICATION-4A3B43
	FilterConfig          *FilterConfig `json:"filterConfig"`          // the condition of an application detection rule
}

func (me *ApplicationDetectionRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The unique name of the Application detection rule",
		},
		"order": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The order of the rule in the rules list",
		},
		"application_identifier": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The Dynatrace entity ID of the application, for example APPLICATION-4A3B43",
		},
		"filter_config": {
			Type:        hcl.TypeList,
			Required:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(FilterConfig).Schema()},
			Description: "The condition of an application detection rule",
		},
	}
}

func (me *ApplicationDetectionRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"id":                     me.ID,
		"name":                   me.Name,
		"order":                  me.Order,
		"application_identifier": me.ApplicationIdentifier,
		"filter_config":          me.FilterConfig,
	})
}

func (me *ApplicationDetectionRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"id":                     &me.ID,
		"name":                   &me.Name,
		"order":                  &me.Order,
		"application_identifier": &me.ApplicationIdentifier,
		"filter_config":          &me.FilterConfig,
	})
}
