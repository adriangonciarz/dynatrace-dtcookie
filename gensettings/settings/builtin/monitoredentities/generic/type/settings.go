package generictype

import "github.com/dtcookie/hcl"

type Settings struct {
	CreatedBy   string          `json:"createdBy"`   // The user or extension that created this type.
	DisplayName string          `json:"displayName"` // The human readable type name for this entity type.
	Enabled     bool            `json:"enabled"`     // Enables or disables the type
	Name        string          `json:"name"`        // The entity type name. This type name must be unique and must not be changed after creation.
	Rules       ExtractionRules `json:"rules"`       // Specify a list of rules which are evaluated in order. When **any** rule matches, the entity defined according to that rule will be extracted. Subsequent rules will not be evaluated.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"created_by": {
			Type:        hcl.TypeString,
			Description: "The user or extension that created this type.",
			Required:    true,
		},
		"display_name": {
			Type:        hcl.TypeString,
			Description: "The human readable type name for this entity type.",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enables or disables the type",
			Optional:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "The entity type name. This type name must be unique and must not be changed after creation.",
			Required:    true,
		},
		"rules": {
			Type:        hcl.TypeList,
			Description: "Specify a list of rules which are evaluated in order. When **any** rule matches, the entity defined according to that rule will be extracted. Subsequent rules will not be evaluated.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ExtractionRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"created_by":   me.CreatedBy,
		"display_name": me.DisplayName,
		"enabled":      me.Enabled,
		"name":         me.Name,
		"rules":        me.Rules,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"created_by":   &me.CreatedBy,
		"display_name": &me.DisplayName,
		"enabled":      &me.Enabled,
		"name":         &me.Name,
		"rules":        &me.Rules,
	})
}
