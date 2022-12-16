package xhrexclusion

import "github.com/dtcookie/hcl"

type Settings struct {
	XhrExclusionRule string `json:"xhrExclusionRule"` // **Examples:**\n\n   - \\\\/segment1\\\\/segment2\n   - dynatrace\\\\.com\n   - www\\\\.dynatrace\\\\.com\\\\/segment1\\\\/.*[a-zA-Z]\n   - www\\\\.dynatrace\\\\.com:8080\n   - www\\\\.dynatrace\\\\.com:([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])\n   - www\\\\.dynatrace\\\\.com\\\\?param1=value1&param2=.*[a-zA-Z]
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"xhr_exclusion_rule": {
			Type:        hcl.TypeString,
			Description: "**Examples:**\n\n   - \\\\/segment1\\\\/segment2\n   - dynatrace\\\\.com\n   - www\\\\.dynatrace\\\\.com\\\\/segment1\\\\/.*[a-zA-Z]\n   - www\\\\.dynatrace\\\\.com:8080\n   - www\\\\.dynatrace\\\\.com:([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])\n   - www\\\\.dynatrace\\\\.com\\\\?param1=value1&param2=.*[a-zA-Z]",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"xhr_exclusion_rule": me.XhrExclusionRule,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"xhr_exclusion_rule": &me.XhrExclusionRule,
	})
}
