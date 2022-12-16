package advanceddetectionrule

import "github.com/dtcookie/hcl"

type Delimiter struct {
	From      string `json:"from"`      // Delimit from
	RemoveIds bool   `json:"removeIds"` // (e.g. versions, hex, dates, and build numbers)
	To        string `json:"to"`        // Delimit to
}

func (me *Delimiter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"from": {
			Type:        hcl.TypeString,
			Description: "Delimit from",
			Required:    true,
		},
		"remove_ids": {
			Type:        hcl.TypeBool,
			Description: "(e.g. versions, hex, dates, and build numbers)",
			Optional:    true,
		},
		"to": {
			Type:        hcl.TypeString,
			Description: "Delimit to",
			Required:    true,
		},
	}
}

func (me *Delimiter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"from":       me.From,
		"remove_ids": me.RemoveIds,
		"to":         me.To,
	})
}

func (me *Delimiter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"from":       &me.From,
		"remove_ids": &me.RemoveIds,
		"to":         &me.To,
	})
}
