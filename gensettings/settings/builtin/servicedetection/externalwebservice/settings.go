package externalwebservice

import "github.com/dtcookie/hcl"

type Settings struct {
	Conditions      Conditions          `json:"conditions"`      // A list of conditions for the rule.\nIf multiple conditions are specified, they *all* must match for the rule to apply.
	Description     *string             `json:"description"`     // Description
	Enabled         bool                `json:"enabled"`         // Enabled
	IdContributors  *IdContributorsType `json:"idContributors"`  // All contributors of the service identifier calculation.
	ManagementZones []string            `json:"managementZones"` // Define a management zone filter for this service detection rule.
	Name            string              `json:"name"`            // Rule name
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"conditions": {
			Type:        hcl.TypeList,
			Description: "A list of conditions for the rule.\nIf multiple conditions are specified, they *all* must match for the rule to apply.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Conditions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"description": {
			Type:        hcl.TypeString,
			Description: "Description",
			Optional:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"id_contributors": {
			Type:        hcl.TypeList,
			Description: "All contributors of the service identifier calculation.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(IdContributorsType).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"management_zones": {
			Type:        hcl.TypeSet,
			Description: "Define a management zone filter for this service detection rule.",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "Rule name",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"conditions":       me.Conditions,
		"description":      me.Description,
		"enabled":          me.Enabled,
		"id_contributors":  me.IdContributors,
		"management_zones": me.ManagementZones,
		"name":             me.Name,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"conditions":       &me.Conditions,
		"description":      &me.Description,
		"enabled":          &me.Enabled,
		"id_contributors":  &me.IdContributors,
		"management_zones": &me.ManagementZones,
		"name":             &me.Name,
	})
}
