package perdiskoverride

import "github.com/dtcookie/hcl"

type DiskLowInodesDetectionThresholds struct {
	FreeInodesPercentage int `json:"freeInodesPercentage"` // Alert if the percentage of available inodes is lower than this threshold in 3 out of 5 samples
}

func (me *DiskLowInodesDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"free_inodes_percentage": {
			Type:        hcl.TypeInt,
			Description: "Alert if the percentage of available inodes is lower than this threshold in 3 out of 5 samples",
			Required:    true,
		},
	}
}

func (me *DiskLowInodesDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"free_inodes_percentage": me.FreeInodesPercentage,
	})
}

func (me *DiskLowInodesDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"free_inodes_percentage": &me.FreeInodesPercentage,
	})
}
