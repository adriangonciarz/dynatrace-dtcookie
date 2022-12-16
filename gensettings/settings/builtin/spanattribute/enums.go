package spanattribute

type MaskingType string

var MaskingTypes = struct {
	NotMasked                MaskingType
	MaskOnlyConfidentialData MaskingType
	MaskEntireValue          MaskingType
}{
	MaskingType("NOT_MASKED"),
	MaskingType("MASK_ONLY_CONFIDENTIAL_DATA"),
	MaskingType("MASK_ENTIRE_VALUE"),
}
