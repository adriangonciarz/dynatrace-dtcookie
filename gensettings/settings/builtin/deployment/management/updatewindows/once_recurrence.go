package updatewindows

import "github.com/dtcookie/hcl"

type OnceRecurrence struct {
	RecurrenceRange *OnceWindow `json:"recurrenceRange"` // Update time
}

func (me *OnceRecurrence) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"recurrence_range": {
			Type:        hcl.TypeList,
			Description: "Update time",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(OnceWindow).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *OnceRecurrence) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"recurrence_range": me.RecurrenceRange,
	})
}

func (me *OnceRecurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"recurrence_range": &me.RecurrenceRange,
	})
}
