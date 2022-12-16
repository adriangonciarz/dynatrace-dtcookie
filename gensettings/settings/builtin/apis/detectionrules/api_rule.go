package detectionrules

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ApiRules []*ApiRule

func (me *ApiRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(ApiRule).Schema()},
		},
	}
}

func (me ApiRules) MarshalHCL() (map[string]interface{}, error) {
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
		result["condition"] = entries
	}
	return result, nil
}

func (me *ApiRules) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("condition"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(ApiRule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "condition", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type ApiRule struct {
	Base    Base    `json:"base"`
	Matcher Matcher `json:"matcher"`
	Pattern string  `json:"pattern"`
}

func (me *ApiRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"base": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"matcher": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"pattern": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *ApiRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"base":    me.Base,
		"matcher": me.Matcher,
		"pattern": me.Pattern,
	})
}

func (me *ApiRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"base":    &me.Base,
		"matcher": &me.Matcher,
		"pattern": &me.Pattern,
	})
}
