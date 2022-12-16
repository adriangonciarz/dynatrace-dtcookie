package rules

import "github.com/dtcookie/hcl"

type Settings struct {
	Conditions  Conditions `json:"conditions"`  // Conditions
	Description *string    `json:"description"` // Rule description
	Enabled     bool       `json:"enabled"`     // Enabled
	Name        string     `json:"name"`        // Rule name
	ParameterID string     `json:"parameterId"` // Failure detection parameters
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"conditions": {
			Type:        hcl.TypeList,
			Description: "Conditions",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Conditions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"description": {
			Type:        hcl.TypeString,
			Description: "Rule description",
			Optional:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "Rule name",
			Required:    true,
		},
		"parameter_id": {
			Type:        hcl.TypeString,
			Description: "Failure detection parameters",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"conditions":   me.Conditions,
		"description":  me.Description,
		"enabled":      me.Enabled,
		"name":         me.Name,
		"parameter_id": me.ParameterID,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"conditions":   &me.Conditions,
		"description":  &me.Description,
		"enabled":      &me.Enabled,
		"name":         &me.Name,
		"parameter_id": &me.ParameterID,
	})
}
