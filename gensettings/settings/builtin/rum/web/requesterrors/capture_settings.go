package requesterrors

import "github.com/dtcookie/hcl"

type CaptureSettings struct {
	Capture       bool `json:"capture"`       // Capture this error
	ConsiderForAI bool `json:"considerForAi"` // [View more details](https://dt-url.net/hd580p2k)
	ImpactApdex   bool `json:"impactApdex"`   // Include error in Apdex calculations
}

func (me *CaptureSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"capture": {
			Type:        hcl.TypeBool,
			Description: "Capture this error",
			Optional:    true,
		},
		"consider_for_ai": {
			Type:        hcl.TypeBool,
			Description: "[View more details](https://dt-url.net/hd580p2k)",
			Optional:    true,
		},
		"impact_apdex": {
			Type:        hcl.TypeBool,
			Description: "Include error in Apdex calculations",
			Optional:    true,
		},
	}
}

func (me *CaptureSettings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"capture":         me.Capture,
		"consider_for_ai": me.ConsiderForAI,
		"impact_apdex":    me.ImpactApdex,
	})
}

func (me *CaptureSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"capture":         &me.Capture,
		"consider_for_ai": &me.ConsiderForAI,
		"impact_apdex":    &me.ImpactApdex,
	})
}
