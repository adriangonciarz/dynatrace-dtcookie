package rumweb

import "github.com/dtcookie/hcl"

type ResponseTimeFixed struct {
	OverAlertingProtection *OverAlertingProtectionAuto `json:"overAlertingProtection"` // Avoid over-alerting
	ResponseTimeAll        *ResponseTimeFixedAll       `json:"responseTimeAll"`        // Alert if the key performance metric of all requests degrades beyond this threshold:
	ResponseTimeSlowest    *ResponseTimeFixedSlowest   `json:"responseTimeSlowest"`    // Alert if the key performance metric of the slowest 10% of requests degrades beyond this threshold:
	Sensitivity            Sensitivity                 `json:"sensitivity"`
}

func (me *ResponseTimeFixed) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"over_alerting_protection": {
			Type:        hcl.TypeList,
			Description: "Avoid over-alerting",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(OverAlertingProtectionAuto).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"response_time_all": {
			Type:        hcl.TypeList,
			Description: "Alert if the key performance metric of all requests degrades beyond this threshold:",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ResponseTimeFixedAll).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"response_time_slowest": {
			Type:        hcl.TypeList,
			Description: "Alert if the key performance metric of the slowest 10% of requests degrades beyond this threshold:",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ResponseTimeFixedSlowest).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"sensitivity": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *ResponseTimeFixed) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"over_alerting_protection": me.OverAlertingProtection,
		"response_time_all":        me.ResponseTimeAll,
		"response_time_slowest":    me.ResponseTimeSlowest,
		"sensitivity":              me.Sensitivity,
	})
}

func (me *ResponseTimeFixed) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"over_alerting_protection": &me.OverAlertingProtection,
		"response_time_all":        &me.ResponseTimeAll,
		"response_time_slowest":    &me.ResponseTimeSlowest,
		"sensitivity":              &me.Sensitivity,
	})
}
