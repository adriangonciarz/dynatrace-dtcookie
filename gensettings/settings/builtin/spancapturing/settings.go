package spancapturing

import "github.com/dtcookie/hcl"

type Settings struct {
	SpanCaptureRule *SpanCaptureRule `json:"spanCaptureRule"` // Span Capture Rule
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"span_capture_rule": {
			Type:        hcl.TypeList,
			Description: "Span Capture Rule",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SpanCaptureRule).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"span_capture_rule": me.SpanCaptureRule,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"span_capture_rule": &me.SpanCaptureRule,
	})
}
