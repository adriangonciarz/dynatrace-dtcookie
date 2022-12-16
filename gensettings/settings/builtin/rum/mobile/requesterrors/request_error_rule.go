package requesterrors

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type RequestErrorRules []*RequestErrorRule

func (me *RequestErrorRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"errorRule": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(RequestErrorRule).Schema()},
		},
	}
}

func (me RequestErrorRules) MarshalHCL() (map[string]interface{}, error) {
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
		result["errorRule"] = entries
	}
	return result, nil
}

func (me *RequestErrorRules) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("errorRule"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(RequestErrorRule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "errorRule", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type RequestErrorRule struct {
	ErrorCodes string `json:"errorCodes"` // Exclude response codes
}

func (me *RequestErrorRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"error_codes": {
			Type:        hcl.TypeString,
			Description: "Exclude response codes",
			Required:    true,
		},
	}
}

func (me *RequestErrorRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"error_codes": me.ErrorCodes,
	})
}

func (me *RequestErrorRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"error_codes": &me.ErrorCodes,
	})
}
