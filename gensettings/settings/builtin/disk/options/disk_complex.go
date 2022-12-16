package options

import (
	"github.com/dtcookie/hcl"
)

type DiskComplexes []*DiskComplex

func (me *DiskComplexes) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"exclusion": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(DiskComplex).Schema()},
		},
	}
}

func (me DiskComplexes) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("exclusion", me)
}

func (me *DiskComplexes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("exclusion", me)
}

type DiskComplex struct {
	Filesystem string     `json:"filesystem"` // **File system type field:** the type of the file system to be excluded from monitoring. Examples:\n\n* ext4\n* ext3\n* btrfs\n* ext*\n\n⚠️ File system types are case sensitive! \n\nThe wildcard in the last example means to exclude matching file systems such as types ext4 and ext3
	Mountpoint string     `json:"mountpoint"` // **Disk or mount point path field:** the path to where the disk to be excluded from monitoring is mounted. Examples:\n\n* /mnt/my_disk\n* /staff/emp1\n* C:\\\n* /staff/*\n* /disk*\n\n ⚠️ Mount point paths are case sensitive! \n\nThe wildcard in **/staff/*** means to exclude every child folder of /staff.\n\nThe wildcard in **/disk*** means to exclude every mount point starting with /disk, for example /disk1, /disk99,  /diskabc
	Os         OsTypeEnum `json:"os"`         // Operating system
}

func (me *DiskComplex) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"filesystem": {
			Type:        hcl.TypeString,
			Description: "**File system type field:** the type of the file system to be excluded from monitoring. Examples:\n\n* ext4\n* ext3\n* btrfs\n* ext*\n\n⚠️ File system types are case sensitive! \n\nThe wildcard in the last example means to exclude matching file systems such as types ext4 and ext3",
			Optional:    true,
		},
		"mountpoint": {
			Type:        hcl.TypeString,
			Description: "**Disk or mount point path field:** the path to where the disk to be excluded from monitoring is mounted. Examples:\n\n* /mnt/my_disk\n* /staff/emp1\n* C:\\\n* /staff/*\n* /disk*\n\n ⚠️ Mount point paths are case sensitive! \n\nThe wildcard in **/staff/*** means to exclude every child folder of /staff.\n\nThe wildcard in **/disk*** means to exclude every mount point starting with /disk, for example /disk1, /disk99,  /diskabc",
			Optional:    true,
		},
		"os": {
			Type:        hcl.TypeString,
			Description: "Operating system",
			Required:    true,
		},
	}
}

func (me *DiskComplex) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"filesystem": me.Filesystem,
		"mountpoint": me.Mountpoint,
		"os":         me.Os,
	})
}

func (me *DiskComplex) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"filesystem": &me.Filesystem,
		"mountpoint": &me.Mountpoint,
		"os":         &me.Os,
	})
}
