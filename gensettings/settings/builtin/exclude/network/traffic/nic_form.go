package traffic

import (
	"github.com/dtcookie/hcl"
)

type NicForms []*NicForm

func (me *NicForms) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"nic_form": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(NicForm).Schema()},
		},
	}
}

func (me NicForms) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("nic_form", me)
}

func (me *NicForms) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("nic_form", me)
}

type NicForm struct {
	Interface string     `json:"interface"` // Network interface
	Os        OsTypeEnum `json:"os"`        // Operating system
}

func (me *NicForm) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"interface": {
			Type:        hcl.TypeString,
			Description: "Network interface",
			Required:    true,
		},
		"os": {
			Type:        hcl.TypeString,
			Description: "Operating system",
			Required:    true,
		},
	}
}

func (me *NicForm) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"interface": me.Interface,
		"os":        me.Os,
	})
}

func (me *NicForm) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"interface": &me.Interface,
		"os":        &me.Os,
	})
}
