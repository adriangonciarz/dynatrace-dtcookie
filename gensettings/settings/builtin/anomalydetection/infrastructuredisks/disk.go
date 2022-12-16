package infrastructuredisks

import "github.com/dtcookie/hcl"

type Disk struct {
	DiskLowInodesDetection          *DiskLowInodesDetection          `json:"diskLowInodesDetection"`
	DiskLowSpaceDetection           *DiskLowSpaceDetection           `json:"diskLowSpaceDetection"`
	DiskSlowWritesAndReadsDetection *DiskSlowWritesAndReadsDetection `json:"diskSlowWritesAndReadsDetection"`
}

func (me *Disk) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"disk_low_inodes_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DiskLowInodesDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"disk_low_space_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DiskLowSpaceDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"disk_slow_writes_and_reads_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DiskSlowWritesAndReadsDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Disk) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"disk_low_inodes_detection":            me.DiskLowInodesDetection,
		"disk_low_space_detection":             me.DiskLowSpaceDetection,
		"disk_slow_writes_and_reads_detection": me.DiskSlowWritesAndReadsDetection,
	})
}

func (me *Disk) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"disk_low_inodes_detection":            &me.DiskLowInodesDetection,
		"disk_low_space_detection":             &me.DiskLowSpaceDetection,
		"disk_slow_writes_and_reads_detection": &me.DiskSlowWritesAndReadsDetection,
	})
}
