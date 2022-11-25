package dimensions

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// String A filter for the metrics string dimensions.
type String struct {
	BaseDimension
	TextFilter *Filter `json:"textFilter"` // A filter for a string value based on the given operator.
}

func (me *String) GetType() FilterType {
	return FilterTypes.String
}

func (me *String) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The dimensions key on the metric",
		},
		"name": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "No documentation available",
		},
		"index": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "No documentation available",
		},
		"filter": {
			Type:        hcl.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "A filter for a string value based on the given operator",
			Elem:        &hcl.Resource{Schema: new(Filter).Schema()},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *String) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.Key != nil {
		result["key"] = *me.Key
	}
	if me.Name != nil {
		result["name"] = *me.Name
	}
	if me.Index != nil {
		result["index"] = *me.Index
	}
	if me.TextFilter != nil {
		if marshalled, err := me.TextFilter.MarshalHCL(); err == nil {
			result["filter"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *String) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "key")
		delete(me.Unknowns, "name")
		delete(me.Unknowns, "index")
		delete(me.Unknowns, "filterType")
		delete(me.Unknowns, "textFilter")

		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("key"); ok {
		me.Key = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("index"); ok {
		me.Index = opt.NewInt(value.(int))
	}
	if _, ok := decoder.GetOk("filter.#"); ok {
		me.TextFilter = new(Filter)
		if err := me.TextFilter.UnmarshalHCL(hcl.NewDecoder(decoder, "filter", 0)); err != nil {
			return err
		}
	}
	return nil
}

func (me *String) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"filterType": me.GetType(),
		"key":        me.Key,
		"name":       me.Name,
		"index":      me.Index,
		"textFilter": me.TextFilter,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *String) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"filterType": &me.FilterType,
		"key":        &me.Key,
		"name":       &me.Name,
		"index":      &me.Index,
		"textFilter": &me.TextFilter,
	}); err != nil {
		return err
	}
	return nil
}
