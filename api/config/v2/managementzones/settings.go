package managementzones

import "github.com/dtcookie/hcl"

type Settings struct {
	Description *string `json:"description,omitempty"` // Description
	Rules       Rules   `json:"rules,omitempty"`       // Rules
	Name        string  `json:"name"`                  // **Be careful when renaming** - if there are policies that are referencing this Management zone, they will need to be adapted to the new name!
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"description": {
			Type:        hcl.TypeString,
			Description: "Description",
			Optional:    true,
		},
		"rules": {
			Type:        hcl.TypeList,
			Description: "Rules",
			MinItems:    1,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(Rules).Schema()},
			Optional:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "**Be careful when renaming** - if there are policies that are referencing this Management zone, they will need to be adapted to the new name!",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"description": me.Description,
		"rules":       me.Rules,
		"name":        me.Name,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"description": &me.Description,
		"rules":       &me.Rules,
		"name":        &me.Name,
	})
}
