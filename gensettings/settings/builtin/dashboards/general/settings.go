package general

import "github.com/dtcookie/hcl"

type Settings struct {
	DefaultDashboardList UserGroupss `json:"defaultDashboardList"` // Configure home dashboard for selected user group. The selected preset dashboard will be loaded as default landing page for this environment.
	EnablePublicSharing  bool        `json:"enablePublicSharing"`  // Allow users to grant anonymous access to dashboards. No sign-in will be required to view those dashboards read-only.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"default_dashboard_list": {
			Type:        hcl.TypeList,
			Description: "Configure home dashboard for selected user group. The selected preset dashboard will be loaded as default landing page for this environment.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(UserGroupss).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enable_public_sharing": {
			Type:        hcl.TypeBool,
			Description: "Allow users to grant anonymous access to dashboards. No sign-in will be required to view those dashboards read-only.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"default_dashboard_list": me.DefaultDashboardList,
		"enable_public_sharing":  me.EnablePublicSharing,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"default_dashboard_list": &me.DefaultDashboardList,
		"enable_public_sharing":  &me.EnablePublicSharing,
	})
}
