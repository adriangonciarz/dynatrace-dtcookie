package assignedapplications

import "github.com/dtcookie/hcl"

type Settings struct {
	Applications Applications `json:"applications"` // Assigned applications
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"applications": {
			Type:        hcl.TypeList,
			Description: "Assigned applications",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Applications).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"applications": me.Applications,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"applications": &me.Applications,
	})
}
