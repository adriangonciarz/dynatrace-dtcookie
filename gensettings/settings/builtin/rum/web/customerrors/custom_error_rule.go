package customerrors

import (
	"github.com/dtcookie/hcl"
)

type CustomErrorRules []*CustomErrorRule

func (me *CustomErrorRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"errorRule": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(CustomErrorRule).Schema()},
		},
	}
}

func (me CustomErrorRules) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("errorRule", me)
}

func (me *CustomErrorRules) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("errorRule", me)
}

type CustomErrorRule struct {
	CaptureSettings *CaptureSettings `json:"captureSettings"` // Capture settings
	KeyMatcher      Matcher          `json:"keyMatcher"`      // Match key
	KeyPattern      string           `json:"keyPattern"`      // A case-insensitive key pattern
	ValueMatcher    Matcher          `json:"valueMatcher"`    // Match value
	ValuePattern    string           `json:"valuePattern"`    // A case-insensitive value pattern
}

func (me *CustomErrorRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"capture_settings": {
			Type:        hcl.TypeList,
			Description: "Capture settings",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CaptureSettings).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"key_matcher": {
			Type:        hcl.TypeString,
			Description: "Match key",
			Required:    true,
		},
		"key_pattern": {
			Type:        hcl.TypeString,
			Description: "A case-insensitive key pattern",
			Required:    true,
		},
		"value_matcher": {
			Type:        hcl.TypeString,
			Description: "Match value",
			Required:    true,
		},
		"value_pattern": {
			Type:        hcl.TypeString,
			Description: "A case-insensitive value pattern",
			Required:    true,
		},
	}
}

func (me *CustomErrorRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"capture_settings": me.CaptureSettings,
		"key_matcher":      me.KeyMatcher,
		"key_pattern":      me.KeyPattern,
		"value_matcher":    me.ValueMatcher,
		"value_pattern":    me.ValuePattern,
	})
}

func (me *CustomErrorRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"capture_settings": &me.CaptureSettings,
		"key_matcher":      &me.KeyMatcher,
		"key_pattern":      &me.KeyPattern,
		"value_matcher":    &me.ValueMatcher,
		"value_pattern":    &me.ValuePattern,
	})
}
