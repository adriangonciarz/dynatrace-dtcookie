package timestampconfiguration

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Matchers []*Matcher

func (me *Matchers) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"matcher": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(Matcher).Schema()},
		},
	}
}

func (me Matchers) MarshalHCL() (map[string]interface{}, error) {
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
		result["matcher"] = entries
	}
	return result, nil
}

func (me *Matchers) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("matcher"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(Matcher)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "matcher", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type Matcher struct {
	Attribute MatcherType `json:"attribute"`
	Operator  Operator    `json:"operator"`
	Values    []string    `json:"values"`
}

func (me *Matcher) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"attribute": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"values": {
			Type:        hcl.TypeSet,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Matcher) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"attribute": me.Attribute,
		"operator":  me.Operator,
		"values":    me.Values,
	})
}

func (me *Matcher) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute": &me.Attribute,
		"operator":  &me.Operator,
		"values":    &me.Values,
	})
}
