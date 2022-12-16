package spaneventattribute

type MaskingType string

var MaskingTypes = struct {
	MaskEntireValue          MaskingType
	MaskOnlyConfidentialData MaskingType
	NotMasked                MaskingType
}{
	MaskingType("MASK_ENTIRE_VALUE"),
	MaskingType("MASK_ONLY_CONFIDENTIAL_DATA"),
	MaskingType("NOT_MASKED"),
}
