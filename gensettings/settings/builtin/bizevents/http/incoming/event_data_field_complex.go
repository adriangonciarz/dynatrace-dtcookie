package incoming

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type EventDataFieldComplexes []*EventDataFieldComplex

func (me *EventDataFieldComplexes) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"event_data_field_complex": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(EventDataFieldComplex).Schema()},
		},
	}
}

func (me EventDataFieldComplexes) MarshalHCL() (map[string]interface{}, error) {
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
		result["event_data_field_complex"] = entries
	}
	return result, nil
}

func (me *EventDataFieldComplexes) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("event_data_field_complex"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(EventDataFieldComplex)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "event_data_field_complex", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type EventDataFieldComplex struct {
	Name   string                     `json:"name"` // Field name to be added to data.
	Source *EventDataAttributeComplex `json:"source"`
}

func (me *EventDataFieldComplex) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "Field name to be added to data.",
			Required:    true,
		},
		"source": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EventDataAttributeComplex).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *EventDataFieldComplex) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"name":   me.Name,
		"source": me.Source,
	})
}

func (me *EventDataFieldComplex) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"name":   &me.Name,
		"source": &me.Source,
	})
}
