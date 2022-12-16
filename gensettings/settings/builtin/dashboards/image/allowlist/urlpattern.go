package allowlist

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type URLPatterns []*URLPattern

func (me *URLPatterns) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"urlpattern": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(URLPattern).Schema()},
		},
	}
}

func (me URLPatterns) MarshalHCL() (map[string]interface{}, error) {
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
		result["urlpattern"] = entries
	}
	return result, nil
}

func (me *URLPatterns) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("urlpattern"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(URLPattern)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "urlpattern", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type URLPattern struct {
	Rule     RuleEnum `json:"rule"`
	Template string   `json:"template"` // Pattern
}

func (me *URLPattern) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rule": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"template": {
			Type:        hcl.TypeString,
			Description: "Pattern",
			Required:    true,
		},
	}
}

func (me *URLPattern) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"rule":     me.Rule,
		"template": me.Template,
	})
}

func (me *URLPattern) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"rule":     &me.Rule,
		"template": &me.Template,
	})
}
