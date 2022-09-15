package alerting

import (
	"github.com/dtcookie/hcl"
)

type EventFilters []*EventFilter

func (me *EventFilters) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"filter": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "A conditions for the metric usage",
			Elem:        &hcl.Resource{Schema: new(EventFilter).Schema()},
		},
	}
}

func (me EventFilters) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("filter", me)
}

func (me *EventFilters) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("filter", me); err != nil {
		return err
	}
	return nil
}

type EventFilter struct {
	Type       EventFilterType        `json:"type"`             // The type of event to filter by
	Predefined *PredefinedEventFilter `json:"predefinedFilter"` // The predefined filter. Only valid if `type` is `PREDEFINED`
	Custom     *CustomEventFilter     `json:"customFilter"`     // The custom filter. Only valid if `type` is `CUSTOM`
}

func (me *EventFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom": {
			Type:        hcl.TypeList,
			Description: "Configuration of a custom event filter. Filters custom events by title or description. If both specified, the AND logic applies",
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(CustomEventFilter).Schema()},
		},
		"predefined": {
			Type:        hcl.TypeList,
			Description: "Configuration of a custom event filter. Filters custom events by title or description. If both specified, the AND logic applies",
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(PredefinedEventFilter).Schema()},
		},
	}
}

func (me *EventFilter) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if me.Custom != nil {
		if marshalled, err := me.Custom.MarshalHCL(); err == nil {
			result["custom"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Predefined != nil {
		if marshalled, err := me.Predefined.MarshalHCL(); err == nil {
			result["predefined"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *EventFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	if _, ok := decoder.GetOk("custom.#"); ok {
		me.Custom = new(CustomEventFilter)
		if err := me.Custom.UnmarshalHCL(hcl.NewDecoder(decoder, "custom", 0)); err != nil {
			return err
		}
		me.Type = EventFilterTypes.Custom
	}
	if _, ok := decoder.GetOk("predefined.#"); ok {
		me.Predefined = new(PredefinedEventFilter)
		if err := me.Predefined.UnmarshalHCL(hcl.NewDecoder(decoder, "predefined", 0)); err != nil {
			return err
		}
		me.Type = EventFilterTypes.Predefined
	}
	return nil
}
