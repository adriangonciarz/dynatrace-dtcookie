package providerbreakdown

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DomainNamePatternListObjects []*DomainNamePatternListObject

func (me *DomainNamePatternListObjects) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"domainNamePattern": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(DomainNamePatternListObject).Schema()},
		},
	}
}

func (me DomainNamePatternListObjects) MarshalHCL() (map[string]interface{}, error) {
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
		result["domainNamePattern"] = entries
	}
	return result, nil
}

func (me *DomainNamePatternListObjects) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("domainNamePattern"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(DomainNamePatternListObject)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "domainNamePattern", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type DomainNamePatternListObject struct {
	DomainNamePattern string `json:"domainNamePattern"` // Please type at least part of this content provider's URL. Asterisks (*) can be used as wildcard characters.
}

func (me *DomainNamePatternListObject) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"domain_name_pattern": {
			Type:        hcl.TypeString,
			Description: "Please type at least part of this content provider's URL. Asterisks (*) can be used as wildcard characters.",
			Required:    true,
		},
	}
}

func (me *DomainNamePatternListObject) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"domain_name_pattern": me.DomainNamePattern,
	})
}

func (me *DomainNamePatternListObject) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"domain_name_pattern": &me.DomainNamePattern,
	})
}
