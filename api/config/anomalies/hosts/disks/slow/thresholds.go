package slow

import "github.com/dtcookie/hcl"

// Thresholds Custom thresholds for slow running disks. If not set, the automatic mode is used.
type Thresholds struct {
	WriteAndReadTime int32 `json:"writeAndReadTime"` // Alert if disk read/write time is higher than *X* milliseconds in 3 out of 5 samples.
}

func (me *Thresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"write_and_read_time": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Alert if disk read/write time is higher than *X* milliseconds in 3 out of 5 samples",
		},
	}
}

func (me *Thresholds) MarshalHCL() (map[string]interface{}, error) {
	return map[string]interface{}{
		"write_and_read_time": int(me.WriteAndReadTime),
	}, nil
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("write_and_read_time"); ok {
		me.WriteAndReadTime = int32(value.(int))
	}
	return nil
}
