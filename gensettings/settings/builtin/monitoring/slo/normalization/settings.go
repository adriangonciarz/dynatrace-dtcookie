package normalization

import "github.com/dtcookie/hcl"

type Settings struct {
	Normalize bool `json:"normalize"` // When set to true, the error budget left will be shown in percent of the total error budget. For more details see [SLO normalization help](https://dt-url.net/slo-normalize-error-budget).
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"normalize": {
			Type:        hcl.TypeBool,
			Description: "When set to true, the error budget left will be shown in percent of the total error budget. For more details see [SLO normalization help](https://dt-url.net/slo-normalize-error-budget).",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"normalize": me.Normalize,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"normalize": &me.Normalize,
	})
}
