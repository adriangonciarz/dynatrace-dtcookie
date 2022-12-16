package updatewindows

import "github.com/dtcookie/hcl"

type RecurrenceRange struct {
	End   string `json:"end"`
	Start string `json:"start"`
}

func (me *RecurrenceRange) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"end": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"start": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *RecurrenceRange) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"end":   me.End,
		"start": me.Start,
	})
}

func (me *RecurrenceRange) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"end":   &me.End,
		"start": &me.Start,
	})
}
