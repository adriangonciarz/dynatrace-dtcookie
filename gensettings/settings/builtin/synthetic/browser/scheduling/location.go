package scheduling

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Locations []*Location

func (me *Locations) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"location": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(Location).Schema()},
		},
	}
}

func (me Locations) MarshalHCL() (map[string]interface{}, error) {
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
		result["location"] = entries
	}
	return result, nil
}

func (me *Locations) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("location"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(Location)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "location", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type Location struct {
	Location string `json:"location"`
}

func (me *Location) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"location": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Optional:    true,
		},
	}
}

func (me *Location) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"location": me.Location,
	})
}

func (me *Location) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"location": &me.Location,
	})
}
