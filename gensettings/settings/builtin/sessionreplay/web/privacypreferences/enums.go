package privacypreferences

type MaskingPreset string

var MaskingPresets = struct {
	BlockList     MaskingPreset
	MaskAll       MaskingPreset
	MaskUserInput MaskingPreset
	AllowList     MaskingPreset
}{
	MaskingPreset("BLOCK_LIST"),
	MaskingPreset("MASK_ALL"),
	MaskingPreset("MASK_USER_INPUT"),
	MaskingPreset("ALLOW_LIST"),
}

type MaskingTargetType string

var MaskingTargetTypes = struct {
	Element   MaskingTargetType
	Attribute MaskingTargetType
}{
	MaskingTargetType("ELEMENT"),
	MaskingTargetType("ATTRIBUTE"),
}
