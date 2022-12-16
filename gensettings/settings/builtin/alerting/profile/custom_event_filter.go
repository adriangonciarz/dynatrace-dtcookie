package profile

import "github.com/dtcookie/hcl"

type CustomEventFilter struct {
	DescriptionFilter *TextFilter `json:"descriptionFilter"` // Description filter
	TitleFilter       *TextFilter `json:"titleFilter"`       // Title filter
}

func (me *CustomEventFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"description_filter": {
			Type:        hcl.TypeList,
			Description: "Description filter",
			Optional:    true,
			Elem:        &hcl.Resource{Schema: new(TextFilter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"title_filter": {
			Type:        hcl.TypeList,
			Description: "Title filter",
			Optional:    true,
			Elem:        &hcl.Resource{Schema: new(TextFilter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *CustomEventFilter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"description_filter": me.DescriptionFilter,
		"title_filter":       me.TitleFilter,
	})
}

func (me *CustomEventFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"description_filter": &me.DescriptionFilter,
		"title_filter":       &me.TitleFilter,
	})
}
