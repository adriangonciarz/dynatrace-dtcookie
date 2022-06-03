package web

import (
	"github.com/dtcookie/hcl"
)

// SessionReplaySettings Session replay settings
type SessionReplaySettings struct {
	Enabled                            bool     `json:"enabled"`                            // SessionReplay Enabled
	CostControlPercentage              int32    `json:"costControlPercentage"`              // Session replay sampling rating in percentage
	EnableCSSResourceCapturing         bool     `json:"enableCssResourceCapturing"`         // Capture (`true`) or don't capture (`false`) CSS resources from the session
	CSSResourceCapturingExclusionRules []string `json:"cssResourceCapturingExclusionRules"` // A list of URLs to be excluded from CSS resource capturing
}

func (me *SessionReplaySettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "SessionReplay Enabled/Disabled",
			Optional:    true,
		},
		"cost_control_percentage": {
			Type:        hcl.TypeInt,
			Description: "Session replay sampling rating in percent",
			Required:    true,
		},
		"enable_css_resource_capturing": {
			Type:        hcl.TypeBool,
			Description: "Capture (`true`) or don't capture (`false`) CSS resources from the session",
			Optional:    true,
		},
		"css_resource_capturing_exclusion_rules": {
			Type:        hcl.TypeList,
			Description: "A list of URLs to be excluded from CSS resource capturing",
			Optional:    true,
			// MinItems: 1,
			Elem: &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *SessionReplaySettings) MarshalHCL() (map[string]interface{}, error) {
	res, err := hcl.Properties{}.EncodeAll(map[string]interface{}{
		"enabled":                                me.Enabled,
		"cost_control_percentage":                me.CostControlPercentage,
		"enable_css_resource_capturing":          me.EnableCSSResourceCapturing,
		"css_resource_capturing_exclusion_rules": me.CSSResourceCapturingExclusionRules,
	})
	if err != nil {
		return nil, err
	}
	if len(me.CSSResourceCapturingExclusionRules) == 0 {
		me.CSSResourceCapturingExclusionRules = nil
		res["css_resource_capturing_exclusion_rules"] = nil
		delete(res, "css_resource_capturing_exclusion_rules")
	}
	return res, nil
}

func (me *SessionReplaySettings) UnmarshalHCL(decoder hcl.Decoder) error {
	err := decoder.DecodeAll(map[string]interface{}{
		"enabled":                                &me.Enabled,
		"cost_control_percentage":                &me.CostControlPercentage,
		"enable_css_resource_capturing":          &me.EnableCSSResourceCapturing,
		"css_resource_capturing_exclusion_rules": &me.CSSResourceCapturingExclusionRules,
	})
	if err != nil {
		return err
	}
	if me.CSSResourceCapturingExclusionRules == nil {
		me.CSSResourceCapturingExclusionRules = []string{}
	}
	return nil
}
