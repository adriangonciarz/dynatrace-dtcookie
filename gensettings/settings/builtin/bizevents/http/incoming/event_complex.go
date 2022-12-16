package incoming

import "github.com/dtcookie/hcl"

type EventComplex struct {
	Category *EventCategoryAttributeComplex `json:"category"` // Event category
	Data     EventDataFieldComplexes        `json:"data"`     // Additional attributes for the business event.
	Provider *EventAttributeComplex         `json:"provider"` // Event provider
	Type     *EventAttributeComplex         `json:"type"`     // Event type
}

func (me *EventComplex) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"category": {
			Type:        hcl.TypeList,
			Description: "Event category",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EventCategoryAttributeComplex).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"data": {
			Type:        hcl.TypeList,
			Description: "Additional attributes for the business event.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EventDataFieldComplexes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"provider": {
			Type:        hcl.TypeList,
			Description: "Event provider",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EventAttributeComplex).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"type": {
			Type:        hcl.TypeList,
			Description: "Event type",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EventAttributeComplex).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *EventComplex) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"category": me.Category,
		"data":     me.Data,
		"provider": me.Provider,
		"type":     me.Type,
	})
}

func (me *EventComplex) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"category": &me.Category,
		"data":     &me.Data,
		"provider": &me.Provider,
		"type":     &me.Type,
	})
}
