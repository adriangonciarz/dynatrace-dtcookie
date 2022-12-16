package privacypreferences

type MaskingPreset string

var MaskingPresets = struct {
	AllowList     MaskingPreset
	BlockList     MaskingPreset
	MaskAll       MaskingPreset
	MaskUserInput MaskingPreset
}{
	MaskingPreset("ALLOW_LIST"),
	MaskingPreset("BLOCK_LIST"),
	MaskingPreset("MASK_ALL"),
	MaskingPreset("MASK_USER_INPUT"),
}

type MaskingTargetType string

var MaskingTargetTypes = struct {
	Attribute MaskingTargetType
	Element   MaskingTargetType
}{
	MaskingTargetType("ATTRIBUTE"),
	MaskingTargetType("ELEMENT"),
}
