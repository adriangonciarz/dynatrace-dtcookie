package infrastructurevmware

import "github.com/dtcookie/hcl"

// EsxiHighMemoryDetectionThresholds. Alert if the condition is met in 3 out of 5 samples
type EsxiHighMemoryDetectionThresholds struct {
	CompressionDecompressionRate float64 `json:"compressionDecompressionRate"` // ESXi host swap IN/OUT or compression/decompression rate is higher than
}

func (me *EsxiHighMemoryDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"compression_decompression_rate": {
			Type:        hcl.TypeFloat,
			Description: "ESXi host swap IN/OUT or compression/decompression rate is higher than",
			Required:    true,
		},
	}
}

func (me *EsxiHighMemoryDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"compression_decompression_rate": me.CompressionDecompressionRate,
	})
}

func (me *EsxiHighMemoryDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"compression_decompression_rate": &me.CompressionDecompressionRate,
	})
}
