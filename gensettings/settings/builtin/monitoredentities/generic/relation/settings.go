package relation

import "github.com/dtcookie/hcl"

type Settings struct {
	CreatedBy      string        `json:"createdBy"`      // The user or extension that created this relationship.
	Enabled        bool          `json:"enabled"`        // Enables or disables the relationship
	FromRole       *string       `json:"fromRole"`       // Specify a role for the source entity. If both source and destination type are the same, referring different roles will allow identification of a relationships direction. If role is left blank, any role of the source type is considered for the relationship.
	FromType       string        `json:"fromType"`       // Define an entity type as the source of the relationship.
	Sources        SourceFilters `json:"sources"`        // Specify all sources which should be evaluated for this relationship rule. The relationship is only created when any of the filters match.
	ToRole         *string       `json:"toRole"`         // Specify a role for the destination entity. If both source and destination type are the same, referring different roles will allow identification of a relationships direction. If role is left blank, any role of the destination type is considered for the relationship.
	ToType         string        `json:"toType"`         // Define an entity type as the destination of the relationship. You can choose the same type as the source type. In this case you also may assign different roles for source and destination for having directed relationships.
	TypeOfRelation RelationType  `json:"typeOfRelation"` // Type of the relationship between the Source Type and the Destination Type
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"created_by": {
			Type:        hcl.TypeString,
			Description: "The user or extension that created this relationship.",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enables or disables the relationship",
			Optional:    true,
		},
		"from_role": {
			Type:        hcl.TypeString,
			Description: "Specify a role for the source entity. If both source and destination type are the same, referring different roles will allow identification of a relationships direction. If role is left blank, any role of the source type is considered for the relationship.",
			Optional:    true,
		},
		"from_type": {
			Type:        hcl.TypeString,
			Description: "Define an entity type as the source of the relationship.",
			Required:    true,
		},
		"sources": {
			Type:        hcl.TypeList,
			Description: "Specify all sources which should be evaluated for this relationship rule. The relationship is only created when any of the filters match.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SourceFilters).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"to_role": {
			Type:        hcl.TypeString,
			Description: "Specify a role for the destination entity. If both source and destination type are the same, referring different roles will allow identification of a relationships direction. If role is left blank, any role of the destination type is considered for the relationship.",
			Optional:    true,
		},
		"to_type": {
			Type:        hcl.TypeString,
			Description: "Define an entity type as the destination of the relationship. You can choose the same type as the source type. In this case you also may assign different roles for source and destination for having directed relationships.",
			Required:    true,
		},
		"type_of_relation": {
			Type:        hcl.TypeString,
			Description: "Type of the relationship between the Source Type and the Destination Type",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"created_by":       me.CreatedBy,
		"enabled":          me.Enabled,
		"from_role":        me.FromRole,
		"from_type":        me.FromType,
		"sources":          me.Sources,
		"to_role":          me.ToRole,
		"to_type":          me.ToType,
		"type_of_relation": me.TypeOfRelation,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"created_by":       &me.CreatedBy,
		"enabled":          &me.Enabled,
		"from_role":        &me.FromRole,
		"from_type":        &me.FromType,
		"sources":          &me.Sources,
		"to_role":          &me.ToRole,
		"to_type":          &me.ToType,
		"type_of_relation": &me.TypeOfRelation,
	})
}
