package browserexclusion

import (
	"github.com/dtcookie/hcl"
)

type BrowserExclusionListObjects []*BrowserExclusionListObject

func (me *BrowserExclusionListObjects) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"browserExclusion": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(BrowserExclusionListObject).Schema()},
		},
	}
}

func (me BrowserExclusionListObjects) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("browserExclusion", me)
}

func (me *BrowserExclusionListObjects) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("browserExclusion", me)
}

type BrowserExclusionListObject struct {
	BrowserName       Browser           `json:"browserName"` // Browser
	Platform          Platform          `json:"platform"`
	Version           int               `json:"version"`
	VersionComparator VersionComparator `json:"versionComparator"` // Browser version comparator
}

func (me *BrowserExclusionListObject) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"browser_name": {
			Type:        hcl.TypeString,
			Description: "Browser",
			Required:    true,
		},
		"platform": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"version": {
			Type:        hcl.TypeInt,
			Description: "no documentation available",
			Required:    true,
		},
		"version_comparator": {
			Type:        hcl.TypeString,
			Description: "Browser version comparator",
			Required:    true,
		},
	}
}

func (me *BrowserExclusionListObject) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"browser_name":       me.BrowserName,
		"platform":           me.Platform,
		"version":            me.Version,
		"version_comparator": me.VersionComparator,
	})
}

func (me *BrowserExclusionListObject) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"browser_name":       &me.BrowserName,
		"platform":           &me.Platform,
		"version":            &me.Version,
		"version_comparator": &me.VersionComparator,
	})
}
