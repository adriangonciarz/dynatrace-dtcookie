package managementzones

import "github.com/dtcookie/hcl"

type Settings struct {
	Description *string `json:"description" hcl:"description"` // Description
	Rules       Rules   `json:"rules" hcl:"rules"`             // Rules
	Name        string  `json:"name" hcl:"name"`               // **Be careful when renaming** - if there are policies that are referencing this Management zone, they will need to be adapted to the new name!
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"description": {
			Type:        hcl.TypeString,
			Description: "Description",
			Required:    true,
		},
		"rules": {
			Type:        hcl.TypeList,
			Description: "Rules",
			MinItems:    1,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(Rules).Schema()},
			Required:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "**Be careful when renaming** - if there are policies that are referencing this Management zone, they will need to be adapted to the new name!",
			Required:    true,
		},
	}
}
