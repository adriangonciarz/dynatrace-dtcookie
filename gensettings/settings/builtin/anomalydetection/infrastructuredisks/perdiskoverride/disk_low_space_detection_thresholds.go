package perdiskoverride

import "github.com/dtcookie/hcl"

type DiskLowSpaceDetectionThresholds struct {
	FreeSpacePercentage int `json:"freeSpacePercentage"` // Alert if free disk space is lower than this percentage in 3 out of 5 samples
}

func (me *DiskLowSpaceDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"free_space_percentage": {
			Type:        hcl.TypeInt,
			Description: "Alert if free disk space is lower than this percentage in 3 out of 5 samples",
			Required:    true,
		},
	}
}

func (me *DiskLowSpaceDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"free_space_percentage": me.FreeSpacePercentage,
	})
}

func (me *DiskLowSpaceDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"free_space_percentage": &me.FreeSpacePercentage,
	})
}
