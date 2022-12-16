package usersessionexportsettingsv2

import "github.com/dtcookie/hcl"

type ExportBehavior struct {
	CustomConfiguration string `json:"customConfiguration"` // Here you can set additional properties for this export configuration. Changing these values might only be necessary in some rare cases. Please contact support before changing this field.
	DisableNotification bool   `json:"disableNotification"` // Disable notification
	ManagementZone      string `json:"managementZone"`      // Restrict exported sessions to management zone
}

func (me *ExportBehavior) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_configuration": {
			Type:        hcl.TypeString,
			Description: "Here you can set additional properties for this export configuration. Changing these values might only be necessary in some rare cases. Please contact support before changing this field.",
			Optional:    true,
		},
		"disable_notification": {
			Type:        hcl.TypeBool,
			Description: "Disable notification",
			Optional:    true,
		},
		"management_zone": {
			Type:        hcl.TypeString,
			Description: "Restrict exported sessions to management zone",
			Optional:    true,
		},
	}
}

func (me *ExportBehavior) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_configuration": me.CustomConfiguration,
		"disable_notification": me.DisableNotification,
		"management_zone":      me.ManagementZone,
	})
}

func (me *ExportBehavior) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_configuration": &me.CustomConfiguration,
		"disable_notification": &me.DisableNotification,
		"management_zone":      &me.ManagementZone,
	})
}
