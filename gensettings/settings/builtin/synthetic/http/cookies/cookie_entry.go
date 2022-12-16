package cookies

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CookieEntries []*CookieEntry

func (me *CookieEntries) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cookie": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(CookieEntry).Schema()},
		},
	}
}

func (me CookieEntries) MarshalHCL() (map[string]interface{}, error) {
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
		result["cookie"] = entries
	}
	return result, nil
}

func (me *CookieEntries) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("cookie"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(CookieEntry)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "cookie", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type CookieEntry struct {
	Domain string `json:"domain"` // Enclose placeholder values in brackets, for example \\{email\\}
	Name   string `json:"name"`   // Enclose placeholder values in brackets, for example \\{email\\}
	Path   string `json:"path"`   // Enclose placeholder values in brackets, for example \\{email\\}
	Value  string `json:"value"`  // Enclose placeholder values in brackets, for example \\{email\\}
}

func (me *CookieEntry) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"domain": {
			Type:        hcl.TypeString,
			Description: "Enclose placeholder values in brackets, for example \\{email\\}",
			Required:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "Enclose placeholder values in brackets, for example \\{email\\}",
			Required:    true,
		},
		"path": {
			Type:        hcl.TypeString,
			Description: "Enclose placeholder values in brackets, for example \\{email\\}",
			Optional:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "Enclose placeholder values in brackets, for example \\{email\\}",
			Required:    true,
		},
	}
}

func (me *CookieEntry) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"domain": me.Domain,
		"name":   me.Name,
		"path":   me.Path,
		"value":  me.Value,
	})
}

func (me *CookieEntry) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"domain": &me.Domain,
		"name":   &me.Name,
		"path":   &me.Path,
		"value":  &me.Value,
	})
}
