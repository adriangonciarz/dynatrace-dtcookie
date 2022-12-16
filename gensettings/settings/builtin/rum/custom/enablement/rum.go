package enablement

import "github.com/dtcookie/hcl"

type Rum struct {
	CostAndTrafficControl int  `json:"costAndTrafficControl"` // Percentage of user sessions captured and analyzed. By default, Dynatrace captures all user actions and user sessions for analysis. This approach ensures complete insight into your application’s performance and customer experience. You can optionally reduce the granularity of user-action and user-session analysis by capturing a lower percentage of user sessions. While this approach can reduce monitoring costs, it also results in lower visibility into how your customers are using your applications. For example, a setting of 10% results in Dynatrace analyzing only every tenth user session.
	Enabled               bool `json:"enabled"`               // Enable Real User Monitoring
}

func (me *Rum) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cost_and_traffic_control": {
			Type:        hcl.TypeInt,
			Description: "Percentage of user sessions captured and analyzed. By default, Dynatrace captures all user actions and user sessions for analysis. This approach ensures complete insight into your application’s performance and customer experience. You can optionally reduce the granularity of user-action and user-session analysis by capturing a lower percentage of user sessions. While this approach can reduce monitoring costs, it also results in lower visibility into how your customers are using your applications. For example, a setting of 10% results in Dynatrace analyzing only every tenth user session.",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enable Real User Monitoring",
			Optional:    true,
		},
	}
}

func (me *Rum) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"cost_and_traffic_control": me.CostAndTrafficControl,
		"enabled":                  me.Enabled,
	})
}

func (me *Rum) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cost_and_traffic_control": &me.CostAndTrafficControl,
		"enabled":                  &me.Enabled,
	})
}
