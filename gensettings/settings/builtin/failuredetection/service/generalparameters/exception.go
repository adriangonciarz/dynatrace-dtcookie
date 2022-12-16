package generalparameters

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Exceptions []*Exception

func (me *Exceptions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"customHandledException": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(Exception).Schema()},
		},
	}
}

func (me Exceptions) MarshalHCL() (map[string]interface{}, error) {
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
		result["customHandledException"] = entries
	}
	return result, nil
}

func (me *Exceptions) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("customHandledException"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(Exception)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "customHandledException", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type Exception struct {
	ClassPattern   string `json:"classPattern"`   // The pattern will match if it is contained within the actual class name.
	MessagePattern string `json:"messagePattern"` // Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.
}

func (me *Exception) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"class_pattern": {
			Type:        hcl.TypeString,
			Description: "The pattern will match if it is contained within the actual class name.",
			Optional:    true,
		},
		"message_pattern": {
			Type:        hcl.TypeString,
			Description: "Optionally, define an exception message pattern. The pattern will match if the actual exception message contains the pattern.",
			Optional:    true,
		},
	}
}

func (me *Exception) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"class_pattern":   me.ClassPattern,
		"message_pattern": me.MessagePattern,
	})
}

func (me *Exception) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"class_pattern":   &me.ClassPattern,
		"message_pattern": &me.MessagePattern,
	})
}
