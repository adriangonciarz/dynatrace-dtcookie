package maintenancewindow

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled           bool               `json:"enabled"`           // The status of the maintenance window. If `false`, it is not considered during the maintenance window evaluation.
	Filters           Filters            `json:"filters"`           // ## Filters\nAdd filters to limit the scope of maintenance to only select matching entities. If no filter is defined, the maintenance window is valid for the whole environment. Each filter is evaluated separately (**OR**).
	GeneralProperties *GeneralProperties `json:"generalProperties"` // ## Properties
	Schedule          *Schedule          `json:"schedule"`          // ## Schedule
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "The status of the maintenance window. If `false`, it is not considered during the maintenance window evaluation.",
			Optional:    true,
		},
		"filters": {
			Type:        hcl.TypeList,
			Description: "## Filters\nAdd filters to limit the scope of maintenance to only select matching entities. If no filter is defined, the maintenance window is valid for the whole environment. Each filter is evaluated separately (**OR**).",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Filters).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"general_properties": {
			Type:        hcl.TypeList,
			Description: "## Properties",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(GeneralProperties).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"schedule": {
			Type:        hcl.TypeList,
			Description: "## Schedule",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Schedule).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":            me.Enabled,
		"filters":            me.Filters,
		"general_properties": me.GeneralProperties,
		"schedule":           me.Schedule,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":            &me.Enabled,
		"filters":            &me.Filters,
		"general_properties": &me.GeneralProperties,
		"schedule":           &me.Schedule,
	})
}
