package usabilityanalytics

import "github.com/dtcookie/hcl"

type Settings struct {
	DetectRageClicks bool `json:"detectRageClicks"` // Three or more rapid clicks within the same area of a web page are considered to be rage clicks. Rage clicks commonly reflect slow-loading or failed page resources. Rage click counts are compiled for each session and considered in the [User Experience Score](https://dt-url.net/39034wt) .\nWith this setting enabled, a rage click count is compiled for each monitored user session.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"detect_rage_clicks": {
			Type:        hcl.TypeBool,
			Description: "Three or more rapid clicks within the same area of a web page are considered to be rage clicks. Rage clicks commonly reflect slow-loading or failed page resources. Rage click counts are compiled for each session and considered in the [User Experience Score](https://dt-url.net/39034wt) .\nWith this setting enabled, a rage click count is compiled for each monitored user session.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"detect_rage_clicks": me.DetectRageClicks,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"detect_rage_clicks": &me.DetectRageClicks,
	})
}
