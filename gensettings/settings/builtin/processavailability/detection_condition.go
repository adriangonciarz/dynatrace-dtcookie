package processavailability

import (
	"github.com/dtcookie/hcl"
)

type DetectionConditions []*DetectionCondition

func (me *DetectionConditions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rule": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(DetectionCondition).Schema()},
		},
	}
}

func (me DetectionConditions) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("rule", me)
}

func (me *DetectionConditions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("rule", me)
}

type DetectionCondition struct {
	Condition string      `json:"condition"` // - $contains(svc) – Matches if svc appears anywhere in the process property value.\n- $eq(svc.exe) – Matches if svc.exe matches the process property value exactly.\n- $prefix(svc) – Matches if app matches the prefix of the process property value.\n- $suffix(svc.py) – Matches if svc.py matches the suffix of the process property value.\n\nFor example, $suffix(svc.py) would detect processes named loyaltysvc.py and paymentssvc.py.\n\nFor more details, see [Process availability](https://dt-url.net/v923x37).
	Property  ProcessItem `json:"property"`  // Select process property
}

func (me *DetectionCondition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeString,
			Description: "- $contains(svc) – Matches if svc appears anywhere in the process property value.\n- $eq(svc.exe) – Matches if svc.exe matches the process property value exactly.\n- $prefix(svc) – Matches if app matches the prefix of the process property value.\n- $suffix(svc.py) – Matches if svc.py matches the suffix of the process property value.\n\nFor example, $suffix(svc.py) would detect processes named loyaltysvc.py and paymentssvc.py.\n\nFor more details, see [Process availability](https://dt-url.net/v923x37).",
			Required:    true,
		},
		"property": {
			Type:        hcl.TypeString,
			Description: "Select process property",
			Required:    true,
		},
	}
}

func (me *DetectionCondition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"condition": me.Condition,
		"property":  me.Property,
	})
}

func (me *DetectionCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"condition": &me.Condition,
		"property":  &me.Property,
	})
}
