package privacypreferences

import "github.com/dtcookie/hcl"

type Settings struct {
	EnableOptInMode         bool                 `json:"enableOptInMode"`         // When [Session Replay opt-in mode](https://dt-url.net/sr-opt-in-mode) is turned on, Session Replay is deactivated until explicitly activated via an API call.
	MaskingPresets          *MaskingPresetConfig `json:"maskingPresets"`          // To protect your end users' privacy, select or customize [predefined masking options](https://dt-url.net/sr-masking-preset-options) that suit your content recording and playback requirements.
	UrlExclusionPatternList []string             `json:"urlExclusionPatternList"` // Exclude webpages or views from Session Replay recording by adding [URL exclusion rules](https://dt-url.net/sr-url-exclusion)
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enable_opt_in_mode": {
			Type:        hcl.TypeBool,
			Description: "When [Session Replay opt-in mode](https://dt-url.net/sr-opt-in-mode) is turned on, Session Replay is deactivated until explicitly activated via an API call.",
			Optional:    true,
		},
		"masking_presets": {
			Type:        hcl.TypeList,
			Description: "To protect your end users' privacy, select or customize [predefined masking options](https://dt-url.net/sr-masking-preset-options) that suit your content recording and playback requirements.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MaskingPresetConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"url_exclusion_pattern_list": {
			Type:        hcl.TypeSet,
			Description: "Exclude webpages or views from Session Replay recording by adding [URL exclusion rules](https://dt-url.net/sr-url-exclusion)",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enable_opt_in_mode":         me.EnableOptInMode,
		"masking_presets":            me.MaskingPresets,
		"url_exclusion_pattern_list": me.UrlExclusionPatternList,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enable_opt_in_mode":         &me.EnableOptInMode,
		"masking_presets":            &me.MaskingPresets,
		"url_exclusion_pattern_list": &me.UrlExclusionPatternList,
	})
}
