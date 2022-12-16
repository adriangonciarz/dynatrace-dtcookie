package privacy

import "github.com/dtcookie/hcl"

type DoNotTrack struct {
	ComplyWithDoNotTrack bool             `json:"complyWithDoNotTrack"` // Comply with \"Do Not Track\" browser settings
	DoNotTrack           DoNotTrackOption `json:"doNotTrack"`
}

func (me *DoNotTrack) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"comply_with_do_not_track": {
			Type:        hcl.TypeBool,
			Description: "Comply with \"Do Not Track\" browser settings",
			Optional:    true,
		},
		"do_not_track": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *DoNotTrack) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"comply_with_do_not_track": me.ComplyWithDoNotTrack,
		"do_not_track":             me.DoNotTrack,
	})
}

func (me *DoNotTrack) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"comply_with_do_not_track": &me.ComplyWithDoNotTrack,
		"do_not_track":             &me.DoNotTrack,
	})
}
