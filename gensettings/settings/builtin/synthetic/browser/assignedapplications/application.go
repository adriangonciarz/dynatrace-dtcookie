package assignedapplications

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Applications []*Application

func (me *Applications) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"application": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(Application).Schema()},
		},
	}
}

func (me Applications) MarshalHCL() (map[string]interface{}, error) {
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
		result["application"] = entries
	}
	return result, nil
}

func (me *Applications) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("application"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(Application)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "application", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type Application struct {
	Application string `json:"application"`
}

func (me *Application) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"application": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *Application) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"application": me.Application,
	})
}

func (me *Application) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"application": &me.Application,
	})
}
