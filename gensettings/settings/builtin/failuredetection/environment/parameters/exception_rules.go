package parameters

import "github.com/dtcookie/hcl"

type ExceptionRules struct {
	CustomErrorRules           CustomErrorRules `json:"customErrorRules"`           // Some custom error situations are only detectable via a return value or other means. To support such cases, [define a request attribute](https://dt-url.net/ys5k0p4y) that captures the required data. Then define a custom error rule that determines if the request has failed based on the value of the request attribute.
	CustomHandledExceptions    Exceptions       `json:"customHandledExceptions"`    // There may be situations where your application code handles exceptions gracefully in a manner that these failures aren't detected by Dynatrace. Use this setting to define specific gracefully-handled exceptions that should be treated as service failures.
	IgnoreAllExceptions        bool             `json:"ignoreAllExceptions"`        // Ignore all exceptions
	IgnoreSpanFailureDetection bool             `json:"ignoreSpanFailureDetection"` // Ignore span failure detection
	IgnoredExceptions          Exceptions       `json:"ignoredExceptions"`          // Some exceptions that are thrown by legacy or 3rd-party code indicate a specific response, not an error. Use this setting to instruct Dynatrace to treat such exceptions as non-failed requests.. If an exception matching any of the defined patterns occurs in a request, it will not be considered as a failure. Other exceptions occurring at the same request might still mark the request as failed.
	SuccessForcingExceptions   Exceptions       `json:"successForcingExceptions"`   // Define exceptions which indicate that a service call should not be considered as failed. E.g. an exception indicating that the client aborted the operation.. If an exception matching any of the defined patterns occurs on the entry node of the service, it will be considered successful. Compared to ignored exceptions, the request will be considered successful even if other exceptions occur in the same request.
}

func (me *ExceptionRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_error_rules": {
			Type:        hcl.TypeList,
			Description: "Some custom error situations are only detectable via a return value or other means. To support such cases, [define a request attribute](https://dt-url.net/ys5k0p4y) that captures the required data. Then define a custom error rule that determines if the request has failed based on the value of the request attribute.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CustomErrorRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"custom_handled_exceptions": {
			Type:        hcl.TypeList,
			Description: "There may be situations where your application code handles exceptions gracefully in a manner that these failures aren't detected by Dynatrace. Use this setting to define specific gracefully-handled exceptions that should be treated as service failures.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Exceptions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"ignore_all_exceptions": {
			Type:        hcl.TypeBool,
			Description: "Ignore all exceptions",
			Optional:    true,
		},
		"ignore_span_failure_detection": {
			Type:        hcl.TypeBool,
			Description: "Ignore span failure detection",
			Optional:    true,
		},
		"ignored_exceptions": {
			Type:        hcl.TypeList,
			Description: "Some exceptions that are thrown by legacy or 3rd-party code indicate a specific response, not an error. Use this setting to instruct Dynatrace to treat such exceptions as non-failed requests.. If an exception matching any of the defined patterns occurs in a request, it will not be considered as a failure. Other exceptions occurring at the same request might still mark the request as failed.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Exceptions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"success_forcing_exceptions": {
			Type:        hcl.TypeList,
			Description: "Define exceptions which indicate that a service call should not be considered as failed. E.g. an exception indicating that the client aborted the operation.. If an exception matching any of the defined patterns occurs on the entry node of the service, it will be considered successful. Compared to ignored exceptions, the request will be considered successful even if other exceptions occur in the same request.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Exceptions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *ExceptionRules) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_error_rules":            me.CustomErrorRules,
		"custom_handled_exceptions":     me.CustomHandledExceptions,
		"ignore_all_exceptions":         me.IgnoreAllExceptions,
		"ignore_span_failure_detection": me.IgnoreSpanFailureDetection,
		"ignored_exceptions":            me.IgnoredExceptions,
		"success_forcing_exceptions":    me.SuccessForcingExceptions,
	})
}

func (me *ExceptionRules) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_error_rules":            &me.CustomErrorRules,
		"custom_handled_exceptions":     &me.CustomHandledExceptions,
		"ignore_all_exceptions":         &me.IgnoreAllExceptions,
		"ignore_span_failure_detection": &me.IgnoreSpanFailureDetection,
		"ignored_exceptions":            &me.IgnoredExceptions,
		"success_forcing_exceptions":    &me.SuccessForcingExceptions,
	})
}
