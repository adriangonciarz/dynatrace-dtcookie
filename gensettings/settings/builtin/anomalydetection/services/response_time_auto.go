package services

import "github.com/dtcookie/hcl"

type ResponseTimeAuto struct {
	OverAlertingProtection *OverAlertingProtection  `json:"overAlertingProtection"` // Avoid over-alerting
	ResponseTimeAll        *ResponseTimeAutoAll     `json:"responseTimeAll"`        // Alert if the average response time of all requests degrades beyond **both** the absolute and relative thresholds:
	ResponseTimeSlowest    *ResponseTimeAutoSlowest `json:"responseTimeSlowest"`    // Alert if the average response time of the slowest 10% of requests degrades beyond **both** the absolute and relative thresholds:
}

func (me *ResponseTimeAuto) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"over_alerting_protection": {
			Type:        hcl.TypeList,
			Description: "Avoid over-alerting",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(OverAlertingProtection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"response_time_all": {
			Type:        hcl.TypeList,
			Description: "Alert if the average response time of all requests degrades beyond **both** the absolute and relative thresholds:",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ResponseTimeAutoAll).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"response_time_slowest": {
			Type:        hcl.TypeList,
			Description: "Alert if the average response time of the slowest 10% of requests degrades beyond **both** the absolute and relative thresholds:",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ResponseTimeAutoSlowest).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *ResponseTimeAuto) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"over_alerting_protection": me.OverAlertingProtection,
		"response_time_all":        me.ResponseTimeAll,
		"response_time_slowest":    me.ResponseTimeSlowest,
	})
}

func (me *ResponseTimeAuto) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"over_alerting_protection": &me.OverAlertingProtection,
		"response_time_all":        &me.ResponseTimeAll,
		"response_time_slowest":    &me.ResponseTimeSlowest,
	})
}
