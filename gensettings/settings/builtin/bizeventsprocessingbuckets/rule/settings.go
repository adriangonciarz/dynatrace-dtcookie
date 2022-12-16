package rule

import "github.com/dtcookie/hcl"

type Settings struct {
	BucketName string `json:"bucketName"` // Events will be stored in the selected bucket. Analyze bucket contents in the [log & event viewer.](/ui/logs-events?advancedQueryMode=true&query=fetch+bizevents)
	Enabled    bool   `json:"enabled"`    // Enabled
	Matcher    string `json:"matcher"`    // [See our documentation](https://dt-url.net/bp234rv)
	RuleName   string `json:"ruleName"`   // Rule name
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"bucket_name": {
			Type:        hcl.TypeString,
			Description: "Events will be stored in the selected bucket. Analyze bucket contents in the [log & event viewer.](/ui/logs-events?advancedQueryMode=true&query=fetch+bizevents)",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"matcher": {
			Type:        hcl.TypeString,
			Description: "[See our documentation](https://dt-url.net/bp234rv)",
			Required:    true,
		},
		"rule_name": {
			Type:        hcl.TypeString,
			Description: "Rule name",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"bucket_name": me.BucketName,
		"enabled":     me.Enabled,
		"matcher":     me.Matcher,
		"rule_name":   me.RuleName,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"bucket_name": &me.BucketName,
		"enabled":     &me.Enabled,
		"matcher":     &me.Matcher,
		"rule_name":   &me.RuleName,
	})
}
