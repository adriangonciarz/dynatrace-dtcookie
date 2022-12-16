package profile

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AlertingProfileEventFilters []*AlertingProfileEventFilter

func (me *AlertingProfileEventFilters) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"eventFilter": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(AlertingProfileEventFilter).Schema()},
		},
	}
}

func (me AlertingProfileEventFilters) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me) > 0 {
		entries := []interface{}{}
		for _, entry := range me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["eventFilter"] = entries
	}
	return result, nil
}

func (me *AlertingProfileEventFilters) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("eventFilter"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(AlertingProfileEventFilter)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "eventFilter", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type AlertingProfileEventFilter struct {
	CustomFilter     *CustomEventFilter             `json:"customFilter"`
	PredefinedFilter *PredefinedEventFilter         `json:"predefinedFilter"`
	Type             AlertingProfileEventFilterType `json:"type"` // Filter problems by any event of source
}

func (me *AlertingProfileEventFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_filter": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CustomEventFilter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"predefined_filter": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(PredefinedEventFilter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "Filter problems by any event of source",
			Required:    true,
		},
	}
}

func (me *AlertingProfileEventFilter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_filter":     me.CustomFilter,
		"predefined_filter": me.PredefinedFilter,
		"type":              me.Type,
	})
}

func (me *AlertingProfileEventFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_filter":     &me.CustomFilter,
		"predefined_filter": &me.PredefinedFilter,
		"type":              &me.Type,
	})
}
