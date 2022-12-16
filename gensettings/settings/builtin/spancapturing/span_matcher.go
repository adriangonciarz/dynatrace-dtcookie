package spancapturing

import (
	"github.com/dtcookie/hcl"
)

type SpanMatchers []*SpanMatcher

func (me *SpanMatchers) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"matcher": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(SpanMatcher).Schema()},
		},
	}
}

func (me SpanMatchers) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("matcher", me)
}

func (me *SpanMatchers) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("matcher", me)
}

type SpanMatcher struct {
	CaseSensitive bool              `json:"caseSensitive"` // affects value and key
	Source        SpanMatcherSource `json:"source"`
	SourceKey     string            `json:"sourceKey"`     // Key
	SpanKindValue SpanKind          `json:"spanKindValue"` // Value
	Type          SpanMatcherType   `json:"type"`          // affects value
	Value         string            `json:"value"`         // evaluated at span start
}

func (me *SpanMatcher) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"case_sensitive": {
			Type:        hcl.TypeBool,
			Description: "affects value and key",
			Optional:    true,
		},
		"source": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"source_key": {
			Type:        hcl.TypeString,
			Description: "Key",
			Required:    true,
		},
		"span_kind_value": {
			Type:        hcl.TypeString,
			Description: "Value",
			Required:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "affects value",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "evaluated at span start",
			Required:    true,
		},
	}
}

func (me *SpanMatcher) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"case_sensitive":  me.CaseSensitive,
		"source":          me.Source,
		"source_key":      me.SourceKey,
		"span_kind_value": me.SpanKindValue,
		"type":            me.Type,
		"value":           me.Value,
	})
}

func (me *SpanMatcher) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"case_sensitive":  &me.CaseSensitive,
		"source":          &me.Source,
		"source_key":      &me.SourceKey,
		"span_kind_value": &me.SpanKindValue,
		"type":            &me.Type,
		"value":           &me.Value,
	})
}
