package requesterrors

import (
	"github.com/dtcookie/hcl"
)

type RequestErrorRules []*RequestErrorRule

func (me *RequestErrorRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"errorRule": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(RequestErrorRule).Schema()},
		},
	}
}

func (me RequestErrorRules) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("errorRule", me)
}

func (me *RequestErrorRules) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("errorRule", me)
}

type RequestErrorRule struct {
	CaptureSettings       *CaptureSettings `json:"captureSettings"`       // Capture settings
	ConsiderCspViolations bool             `json:"considerCspViolations"` // Match by errors that have CSP violations
	ConsiderFailedImages  bool             `json:"considerFailedImages"`  // Match by errors that have failed image requests
	ErrorCodes            string           `json:"errorCodes"`            // Match by error code
	FilterSettings        *FilterSettings  `json:"filterSettings"`        // Filter settings
}

func (me *RequestErrorRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"capture_settings": {
			Type:        hcl.TypeList,
			Description: "Capture settings",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CaptureSettings).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"consider_csp_violations": {
			Type:        hcl.TypeBool,
			Description: "Match by errors that have CSP violations",
			Optional:    true,
		},
		"consider_failed_images": {
			Type:        hcl.TypeBool,
			Description: "Match by errors that have failed image requests",
			Optional:    true,
		},
		"error_codes": {
			Type:        hcl.TypeString,
			Description: "Match by error code",
			Optional:    true,
		},
		"filter_settings": {
			Type:        hcl.TypeList,
			Description: "Filter settings",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(FilterSettings).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *RequestErrorRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"capture_settings":        me.CaptureSettings,
		"consider_csp_violations": me.ConsiderCspViolations,
		"consider_failed_images":  me.ConsiderFailedImages,
		"error_codes":             me.ErrorCodes,
		"filter_settings":         me.FilterSettings,
	})
}

func (me *RequestErrorRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"capture_settings":        &me.CaptureSettings,
		"consider_csp_violations": &me.ConsiderCspViolations,
		"consider_failed_images":  &me.ConsiderFailedImages,
		"error_codes":             &me.ErrorCodes,
		"filter_settings":         &me.FilterSettings,
	})
}
