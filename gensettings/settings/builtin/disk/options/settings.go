package options

import "github.com/dtcookie/hcl"

type Settings struct {
	Exclusions DiskComplexes `json:"exclusions"` // OneAgent automatically detects and monitors all your mount points, however you can create exception rules to remove disks from the monitoring list.
	NfsShowAll bool          `json:"nfsShowAll"` // When disabled OneAgent will try to deduplicate some of nfs disks. Disabled by default, applies only to Linux hosts. Requires OneAgent 1.209 or later
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"exclusions": {
			Type:        hcl.TypeList,
			Description: "OneAgent automatically detects and monitors all your mount points, however you can create exception rules to remove disks from the monitoring list.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DiskComplexes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"nfs_show_all": {
			Type:        hcl.TypeBool,
			Description: "When disabled OneAgent will try to deduplicate some of nfs disks. Disabled by default, applies only to Linux hosts. Requires OneAgent 1.209 or later",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"exclusions":   me.Exclusions,
		"nfs_show_all": me.NfsShowAll,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"exclusions":   &me.Exclusions,
		"nfs_show_all": &me.NfsShowAll,
	})
}
