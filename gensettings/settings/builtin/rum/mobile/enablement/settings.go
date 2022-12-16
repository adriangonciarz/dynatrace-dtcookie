package enablement

import "github.com/dtcookie/hcl"

type Settings struct {
	Rum           *Rum           `json:"rum"`           // Capture and analyze all user actions within your application. Enable [Real User Monitoring (RUM)](https://dt-url.net/1n2b0prq) to monitor and improve your application's performance, identify errors, and gain insight into your user's behavior and experience.
	SessionReplay *SessionReplay `json:"sessionReplay"` // [Session Replay on crashes](https://dt-url.net/session-replay) gives you additional context for crash analysis in the form of video-like screen recordings that replay user actions immediately preceding a detected crash, while providing [best-in-class security and data protection](https://dt-url.net/b303zxj).
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rum": {
			Type:        hcl.TypeList,
			Description: "Capture and analyze all user actions within your application. Enable [Real User Monitoring (RUM)](https://dt-url.net/1n2b0prq) to monitor and improve your application's performance, identify errors, and gain insight into your user's behavior and experience.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Rum).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"session_replay": {
			Type:        hcl.TypeList,
			Description: "[Session Replay on crashes](https://dt-url.net/session-replay) gives you additional context for crash analysis in the form of video-like screen recordings that replay user actions immediately preceding a detected crash, while providing [best-in-class security and data protection](https://dt-url.net/b303zxj).",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SessionReplay).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"rum":            me.Rum,
		"session_replay": me.SessionReplay,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"rum":            &me.Rum,
		"session_replay": &me.SessionReplay,
	})
}
