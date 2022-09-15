package alerting

import (
	"github.com/dtcookie/hcl"
)

type CustomEventFilter struct {
	Title       *TextFilter `json:"titleFilter,omitempty"`       // Title filter
	Description *TextFilter `json:"descriptionFilter,omitempty"` // Description filter
}

func (me *CustomEventFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"description": {
			Type:        hcl.TypeList,
			Description: "Configuration of a matching filter",
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(TextFilter).Schema()},
		},
		"title": {
			Type:        hcl.TypeList,
			Description: "Configuration of a matching filter",
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(TextFilter).Schema()},
		},
	}
}

func (me *CustomEventFilter) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if me.Description != nil {
		if marshalled, err := me.Description.MarshalHCL(); err == nil {
			result["description"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Title != nil {
		if marshalled, err := me.Title.MarshalHCL(); err == nil {
			result["title"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *CustomEventFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	if _, ok := decoder.GetOk("description.#"); ok {
		me.Description = new(TextFilter)
		if err := me.Description.UnmarshalHCL(hcl.NewDecoder(decoder, "description", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("title.#"); ok {
		me.Title = new(TextFilter)
		if err := me.Title.UnmarshalHCL(hcl.NewDecoder(decoder, "title", 0)); err != nil {
			return err
		}
	}
	return nil
}
