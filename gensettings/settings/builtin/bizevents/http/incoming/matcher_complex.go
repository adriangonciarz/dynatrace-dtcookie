package incoming

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type MatcherComplexes []*MatcherComplex

func (me *MatcherComplexes) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"trigger": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(MatcherComplex).Schema()},
		},
	}
}

func (me MatcherComplexes) MarshalHCL() (map[string]interface{}, error) {
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
		result["trigger"] = entries
	}
	return result, nil
}

func (me *MatcherComplexes) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("trigger"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(MatcherComplex)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "trigger", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

// Matcher. Rule must match
type MatcherComplex struct {
	CaseSensitive bool               `json:"caseSensitive"` // Case sensitive
	Source        *DataSourceComplex `json:"source"`
	Type          ComparisonEnum     `json:"type"` // Operator
	Value         string             `json:"value"`
}

func (me *MatcherComplex) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"case_sensitive": {
			Type:        hcl.TypeBool,
			Description: "Case sensitive",
			Optional:    true,
		},
		"source": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DataSourceComplex).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "Operator",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *MatcherComplex) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"case_sensitive": me.CaseSensitive,
		"source":         me.Source,
		"type":           me.Type,
		"value":          me.Value,
	})
}

func (me *MatcherComplex) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"case_sensitive": &me.CaseSensitive,
		"source":         &me.Source,
		"type":           &me.Type,
		"value":          &me.Value,
	})
}
