package general

import (
	"github.com/dtcookie/hcl"
)

type UserGroupss []*UserGroups

func (me *UserGroupss) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"defaultDashboard": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(UserGroups).Schema()},
		},
	}
}

func (me UserGroupss) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("defaultDashboard", me)
}

func (me *UserGroupss) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("defaultDashboard", me)
}

type UserGroups struct {
	Dashboard string `json:"Dashboard"` // Preset dashboard to show as default landing page
	UserGroup string `json:"UserGroup"` // Show selected dashboard by default for this user group
}

func (me *UserGroups) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dashboard": {
			Type:        hcl.TypeString,
			Description: "Preset dashboard to show as default landing page",
			Required:    true,
		},
		"user_group": {
			Type:        hcl.TypeString,
			Description: "Show selected dashboard by default for this user group",
			Required:    true,
		},
	}
}

func (me *UserGroups) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"dashboard":  me.Dashboard,
		"user_group": me.UserGroup,
	})
}

func (me *UserGroups) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"dashboard":  &me.Dashboard,
		"user_group": &me.UserGroup,
	})
}
