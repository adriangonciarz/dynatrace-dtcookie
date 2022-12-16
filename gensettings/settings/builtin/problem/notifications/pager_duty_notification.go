package notifications

import "github.com/dtcookie/hcl"

type PagerDutyNotification struct {
	Account       string `json:"account"`       // The name of the PagerDuty account.
	ServiceApiKey string `json:"serviceApiKey"` // The API key to access PagerDuty.
	ServiceName   string `json:"serviceName"`   // The name of the service.
}

func (me *PagerDutyNotification) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"account": {
			Type:        hcl.TypeString,
			Description: "The name of the PagerDuty account.",
			Required:    true,
		},
		"service_api_key": {
			Type:        hcl.TypeString,
			Description: "The API key to access PagerDuty.",
			Required:    true,
		},
		"service_name": {
			Type:        hcl.TypeString,
			Description: "The name of the service.",
			Required:    true,
		},
	}
}

func (me *PagerDutyNotification) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"account":         me.Account,
		"service_api_key": me.ServiceApiKey,
		"service_name":    me.ServiceName,
	})
}

func (me *PagerDutyNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"account":         &me.Account,
		"service_api_key": &me.ServiceApiKey,
		"service_name":    &me.ServiceName,
	})
}
