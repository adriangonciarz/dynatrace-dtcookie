package privacy

import "github.com/dtcookie/hcl"

type Settings struct {
	DataCollection *DataCollection `json:"dataCollection"` // To provide your end users with the ability to decide for themselves if their activities should be tracked to measure application performance and usage, enable opt-in mode.
	DoNotTrack     *DoNotTrack     `json:"doNotTrack"`     // Most modern web browsers have a privacy feature called [\"Do Not Track\"](https://dt-url.net/sb3n0pnl) that individual users may have enabled on their devices. Customize how Dynatrace should behave when it encounters this setting.
	Masking        *Masking        `json:"masking"`
	UserTracking   *UserTracking   `json:"userTracking"` // User tracking
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"data_collection": {
			Type:        hcl.TypeList,
			Description: "To provide your end users with the ability to decide for themselves if their activities should be tracked to measure application performance and usage, enable opt-in mode.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DataCollection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"do_not_track": {
			Type:        hcl.TypeList,
			Description: "Most modern web browsers have a privacy feature called [\"Do Not Track\"](https://dt-url.net/sb3n0pnl) that individual users may have enabled on their devices. Customize how Dynatrace should behave when it encounters this setting.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DoNotTrack).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"masking": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Masking).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"user_tracking": {
			Type:        hcl.TypeList,
			Description: "User tracking",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(UserTracking).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"data_collection": me.DataCollection,
		"do_not_track":    me.DoNotTrack,
		"masking":         me.Masking,
		"user_tracking":   me.UserTracking,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"data_collection": &me.DataCollection,
		"do_not_track":    &me.DoNotTrack,
		"masking":         &me.Masking,
		"user_tracking":   &me.UserTracking,
	})
}
