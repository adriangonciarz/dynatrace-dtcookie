package requesterrors

import "github.com/dtcookie/hcl"

type FilterSettings struct {
	Filter UrlFilter `json:"filter"` // Filter by URL
	Url    string    `json:"url"`
}

func (me *FilterSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"filter": {
			Type:        hcl.TypeString,
			Description: "Filter by URL",
			Optional:    true,
		},
		"url": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *FilterSettings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"filter": me.Filter,
		"url":    me.Url,
	})
}

func (me *FilterSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"filter": &me.Filter,
		"url":    &me.Url,
	})
}
