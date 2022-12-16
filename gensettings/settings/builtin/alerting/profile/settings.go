package profile

import "github.com/dtcookie/hcl"

type Settings struct {
	EventFilters   AlertingProfileEventFilters  `json:"eventFilters"`   // Define event filters for profile. A maximum of 100 event filters is allowed.
	ManagementZone *string                      `json:"managementZone"` // Define management zone filter for profile
	Name           string                       `json:"name"`           // Name
	SeverityRules  AlertingProfileSeverityRules `json:"severityRules"`  // Define severity rules for profile. A maximum of 100 severity rules is allowed.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"event_filters": {
			Type:        hcl.TypeList,
			Description: "Define event filters for profile. A maximum of 100 event filters is allowed.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(AlertingProfileEventFilters).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"management_zone": {
			Type:        hcl.TypeString,
			Description: "Define management zone filter for profile",
			Optional:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "Name",
			Required:    true,
		},
		"severity_rules": {
			Type:        hcl.TypeList,
			Description: "Define severity rules for profile. A maximum of 100 severity rules is allowed.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(AlertingProfileSeverityRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"event_filters":   me.EventFilters,
		"management_zone": me.ManagementZone,
		"name":            me.Name,
		"severity_rules":  me.SeverityRules,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"event_filters":   &me.EventFilters,
		"management_zone": &me.ManagementZone,
		"name":            &me.Name,
		"severity_rules":  &me.SeverityRules,
	})
}
