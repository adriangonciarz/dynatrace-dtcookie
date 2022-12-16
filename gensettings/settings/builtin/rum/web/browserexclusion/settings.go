package browserexclusion

import "github.com/dtcookie/hcl"

type Settings struct {
	BrowserExclusionInclude bool                        `json:"browserExclusionInclude"` // These are the only browsers that should be monitored
	BrowserExclusionList    BrowserExclusionListObjects `json:"browserExclusionList"`    // Browser exclusion list
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"browser_exclusion_include": {
			Type:        hcl.TypeBool,
			Description: "These are the only browsers that should be monitored",
			Optional:    true,
		},
		"browser_exclusion_list": {
			Type:        hcl.TypeList,
			Description: "Browser exclusion list",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(BrowserExclusionListObjects).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"browser_exclusion_include": me.BrowserExclusionInclude,
		"browser_exclusion_list":    me.BrowserExclusionList,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"browser_exclusion_include": &me.BrowserExclusionInclude,
		"browser_exclusion_list":    &me.BrowserExclusionList,
	})
}
