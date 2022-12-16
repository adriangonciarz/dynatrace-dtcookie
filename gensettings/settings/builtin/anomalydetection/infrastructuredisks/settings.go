package infrastructuredisks

import "github.com/dtcookie/hcl"

type Settings struct {
	Disk *Disk `json:"disk"` // Disk
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"disk": {
			Type:        hcl.TypeList,
			Description: "Disk",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Disk).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"disk": me.Disk,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"disk": &me.Disk,
	})
}
