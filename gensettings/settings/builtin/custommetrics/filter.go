package custommetrics

import (
	"github.com/dtcookie/hcl"
)

type Filters []*Filter

func (me *Filters) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"filter": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(Filter).Schema()},
		},
	}
}

func (me Filters) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("filter", me)
}

func (me *Filters) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("filter", me)
}

type Filter struct {
	FieldName string   `json:"fieldName"` // Field name
	Operator  Operator `json:"operator"`
	Value     string   `json:"value"`
	ValueIn   []string `json:"valueIn"` // Values
}

func (me *Filter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"field_name": {
			Type:        hcl.TypeString,
			Description: "Field name",
			Required:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"value_in": {
			Type:        hcl.TypeList,
			Description: "Values",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Filter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"field_name": me.FieldName,
		"operator":   me.Operator,
		"value":      me.Value,
		"value_in":   me.ValueIn,
	})
}

func (me *Filter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"field_name": &me.FieldName,
		"operator":   &me.Operator,
		"value":      &me.Value,
		"value_in":   &me.ValueIn,
	})
}
