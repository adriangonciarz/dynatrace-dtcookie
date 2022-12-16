package keyperformancemetrics

import "github.com/dtcookie/hcl"

type Settings struct {
	FrustratingIfReportedOrWebRequestError bool        `json:"frustratingIfReportedOrWebRequestError"` // Treat user actions with reported errors or web request errors as erroneous and rate their performance as Frustrating. Turn off this setting if errors should not affect the Apdex rate.
	Thresholds                             *Thresholds `json:"thresholds"`
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"frustrating_if_reported_or_web_request_error": {
			Type:        hcl.TypeBool,
			Description: "Treat user actions with reported errors or web request errors as erroneous and rate their performance as Frustrating. Turn off this setting if errors should not affect the Apdex rate.",
			Optional:    true,
		},
		"thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Thresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"frustrating_if_reported_or_web_request_error": me.FrustratingIfReportedOrWebRequestError,
		"thresholds": me.Thresholds,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"frustrating_if_reported_or_web_request_error": &me.FrustratingIfReportedOrWebRequestError,
		"thresholds": &me.Thresholds,
	})
}
