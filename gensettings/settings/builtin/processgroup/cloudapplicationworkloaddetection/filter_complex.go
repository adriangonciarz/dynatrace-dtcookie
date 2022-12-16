package cloudapplicationworkloaddetection

import (
	"github.com/dtcookie/hcl"
)

type FilterComplexes []*FilterComplex

func (me *FilterComplexes) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"filter": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(FilterComplex).Schema()},
		},
	}
}

func (me FilterComplexes) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("filter", me)
}

func (me *FilterComplexes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("filter", me)
}

type FilterComplex struct {
	Enabled          bool              `json:"enabled"`
	InclusionToggles *InclusionToggles `json:"inclusionToggles"` // ID calculation based on
	MatchFilter      *MatchFilter      `json:"matchFilter"`      // When namespace
}

func (me *FilterComplex) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "no documentation available",
			Optional:    true,
		},
		"inclusion_toggles": {
			Type:        hcl.TypeList,
			Description: "ID calculation based on",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(InclusionToggles).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"match_filter": {
			Type:        hcl.TypeList,
			Description: "When namespace",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MatchFilter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *FilterComplex) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":           me.Enabled,
		"inclusion_toggles": me.InclusionToggles,
		"match_filter":      me.MatchFilter,
	})
}

func (me *FilterComplex) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":           &me.Enabled,
		"inclusion_toggles": &me.InclusionToggles,
		"match_filter":      &me.MatchFilter,
	})
}
