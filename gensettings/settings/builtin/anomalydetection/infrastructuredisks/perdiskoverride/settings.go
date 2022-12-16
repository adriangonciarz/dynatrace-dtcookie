package perdiskoverride

import "github.com/dtcookie/hcl"

type Settings struct {
	DiskLowInodesDetection              *DiskLowInodesDetection          `json:"diskLowInodesDetection"`
	DiskLowSpaceDetection               *DiskLowSpaceDetection           `json:"diskLowSpaceDetection"`
	DiskSlowWritesAndReadsDetection     *DiskSlowWritesAndReadsDetection `json:"diskSlowWritesAndReadsDetection"`
	OverrideDiskLowSpaceDetection       bool                             `json:"overrideDiskLowSpaceDetection"`       // Override low disk space detection settings
	OverrideLowInodesDetection          bool                             `json:"overrideLowInodesDetection"`          // Override low inodes detection settings
	OverrideSlowWritesAndReadsDetection bool                             `json:"overrideSlowWritesAndReadsDetection"` // Override slow writes and reads detection settings
}

func (me *Settings) Schema() map[string]*hcl.Schema {
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
		"override_disk_low_space_detection": {
			Type:        hcl.TypeBool,
			Description: "Override low disk space detection settings",
			Optional:    true,
		},
		"override_low_inodes_detection": {
			Type:        hcl.TypeBool,
			Description: "Override low inodes detection settings",
			Optional:    true,
		},
		"override_slow_writes_and_reads_detection": {
			Type:        hcl.TypeBool,
			Description: "Override slow writes and reads detection settings",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"disk_low_inodes_detection":                me.DiskLowInodesDetection,
		"disk_low_space_detection":                 me.DiskLowSpaceDetection,
		"disk_slow_writes_and_reads_detection":     me.DiskSlowWritesAndReadsDetection,
		"override_disk_low_space_detection":        me.OverrideDiskLowSpaceDetection,
		"override_low_inodes_detection":            me.OverrideLowInodesDetection,
		"override_slow_writes_and_reads_detection": me.OverrideSlowWritesAndReadsDetection,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"disk_low_inodes_detection":                &me.DiskLowInodesDetection,
		"disk_low_space_detection":                 &me.DiskLowSpaceDetection,
		"disk_slow_writes_and_reads_detection":     &me.DiskSlowWritesAndReadsDetection,
		"override_disk_low_space_detection":        &me.OverrideDiskLowSpaceDetection,
		"override_low_inodes_detection":            &me.OverrideLowInodesDetection,
		"override_slow_writes_and_reads_detection": &me.OverrideSlowWritesAndReadsDetection,
	})
}
