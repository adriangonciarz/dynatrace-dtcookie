package privacypreferences

import "github.com/dtcookie/hcl"

// Content masking preferences. To protect your end users' privacy, select or customize [predefined masking options](https://dt-url.net/sr-masking-preset-options) that suit your content recording and playback requirements.// The recording masking settings are applied at record time to all webpages that your users navigate to. The playback masking settings are applied when replaying recorded sessions, including those that were recorded before your masking settings were applied.// Note: When you set the recording masking settings to a more restrictive option, the same option is also enabled for playback masking settings, which affects all past recorded sessions as well.
type MaskingPresetConfig struct {
	PlaybackMaskingAllowListRules  AllowListRules `json:"playbackMaskingAllowListRules"`  // The elements are defined by the CSS selector or attribute name.
	PlaybackMaskingBlockListRules  BlockListRules `json:"playbackMaskingBlockListRules"`  // The elements are defined by the CSS selector or attribute name.
	PlaybackMaskingPreset          MaskingPreset  `json:"playbackMaskingPreset"`          // Playback masking settings are applied during playback of recorded sessions, including playback of sessions that were recorded before these settings were applied.
	RecordingMaskingAllowListRules AllowListRules `json:"recordingMaskingAllowListRules"` // The elements are defined by the CSS selector or attribute name.
	RecordingMaskingBlockListRules BlockListRules `json:"recordingMaskingBlockListRules"` // The elements are defined by the CSS selector or attribute name.
	RecordingMaskingPreset         MaskingPreset  `json:"recordingMaskingPreset"`         // Recording masking settings are applied at record time. When you set these settings to a more restrictive option, the same option is also enabled for the playback masking settings.
}

func (me *MaskingPresetConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"playback_masking_allow_list_rules": {
			Type:        hcl.TypeList,
			Description: "The elements are defined by the CSS selector or attribute name.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(AllowListRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"playback_masking_block_list_rules": {
			Type:        hcl.TypeList,
			Description: "The elements are defined by the CSS selector or attribute name.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(BlockListRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"playback_masking_preset": {
			Type:        hcl.TypeString,
			Description: "Playback masking settings are applied during playback of recorded sessions, including playback of sessions that were recorded before these settings were applied.",
			Required:    true,
		},
		"recording_masking_allow_list_rules": {
			Type:        hcl.TypeList,
			Description: "The elements are defined by the CSS selector or attribute name.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(AllowListRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"recording_masking_block_list_rules": {
			Type:        hcl.TypeList,
			Description: "The elements are defined by the CSS selector or attribute name.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(BlockListRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"recording_masking_preset": {
			Type:        hcl.TypeString,
			Description: "Recording masking settings are applied at record time. When you set these settings to a more restrictive option, the same option is also enabled for the playback masking settings.",
			Required:    true,
		},
	}
}

func (me *MaskingPresetConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"playback_masking_allow_list_rules":  me.PlaybackMaskingAllowListRules,
		"playback_masking_block_list_rules":  me.PlaybackMaskingBlockListRules,
		"playback_masking_preset":            me.PlaybackMaskingPreset,
		"recording_masking_allow_list_rules": me.RecordingMaskingAllowListRules,
		"recording_masking_block_list_rules": me.RecordingMaskingBlockListRules,
		"recording_masking_preset":           me.RecordingMaskingPreset,
	})
}

func (me *MaskingPresetConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"playback_masking_allow_list_rules":  &me.PlaybackMaskingAllowListRules,
		"playback_masking_block_list_rules":  &me.PlaybackMaskingBlockListRules,
		"playback_masking_preset":            &me.PlaybackMaskingPreset,
		"recording_masking_allow_list_rules": &me.RecordingMaskingAllowListRules,
		"recording_masking_block_list_rules": &me.RecordingMaskingBlockListRules,
		"recording_masking_preset":           &me.RecordingMaskingPreset,
	})
}
