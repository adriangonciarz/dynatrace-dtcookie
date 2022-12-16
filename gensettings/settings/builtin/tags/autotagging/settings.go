package autotagging

import "github.com/dtcookie/hcl"

type Settings struct {
	Description *string `json:"description"` // Description
	Name        string  `json:"name"`        // Tag name
	Rules       Rules   `json:"rules"`       // Rules
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"description": {
			Type:        hcl.TypeString,
			Description: "Description",
			Optional:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "Tag name",
			Required:    true,
		},
		"rules": {
			Type:        hcl.TypeList,
			Description: "Rules",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Rules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"description": me.Description,
		"name":        me.Name,
		"rules":       me.Rules,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"description": &me.Description,
		"name":        &me.Name,
		"rules":       &me.Rules,
	})
}
