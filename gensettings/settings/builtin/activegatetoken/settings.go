package activegatetoken

import "github.com/dtcookie/hcl"

type Settings struct {
	AuthTokenEnforcementManuallyEnabled bool `json:"authTokenEnforcementManuallyEnabled"` // Manually enforce ActiveGate token authentication
	ExpiringTokenNotificationsEnabled   bool `json:"expiringTokenNotificationsEnabled"`   // Note: ActiveGate tokens notifications are sent only when you deployed ActiveGate tokens with expiration dates in your environment and notifications are defined ([see notification settings](/ui/settings/builtin:problem.notifications))
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"auth_token_enforcement_manually_enabled": {
			Type:        hcl.TypeBool,
			Description: "Manually enforce ActiveGate token authentication",
			Optional:    true,
		},
		"expiring_token_notifications_enabled": {
			Type:        hcl.TypeBool,
			Description: "Note: ActiveGate tokens notifications are sent only when you deployed ActiveGate tokens with expiration dates in your environment and notifications are defined ([see notification settings](/ui/settings/builtin:problem.notifications))",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"auth_token_enforcement_manually_enabled": me.AuthTokenEnforcementManuallyEnabled,
		"expiring_token_notifications_enabled":    me.ExpiringTokenNotificationsEnabled,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"auth_token_enforcement_manually_enabled": &me.AuthTokenEnforcementManuallyEnabled,
		"expiring_token_notifications_enabled":    &me.ExpiringTokenNotificationsEnabled,
	})
}
