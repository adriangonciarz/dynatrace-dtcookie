package infrastructuredisks

import "github.com/dtcookie/hcl"

type DiskSlowWritesAndReadsDetectionThresholds struct {
	WriteAndReadTime int `json:"writeAndReadTime"` // Alert if disk read time and write time is higher than this threshold in 3 out of 5 samples
}

func (me *DiskSlowWritesAndReadsDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"write_and_read_time": {
			Type:        hcl.TypeInt,
			Description: "Alert if disk read time and write time is higher than this threshold in 3 out of 5 samples",
			Required:    true,
		},
	}
}

func (me *DiskSlowWritesAndReadsDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"write_and_read_time": me.WriteAndReadTime,
	})
}

func (me *DiskSlowWritesAndReadsDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"write_and_read_time": &me.WriteAndReadTime,
	})
}
