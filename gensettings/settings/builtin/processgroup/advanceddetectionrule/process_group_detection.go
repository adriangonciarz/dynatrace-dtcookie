package advanceddetectionrule

import "github.com/dtcookie/hcl"

type ProcessGroupDetection struct {
	ContainedString       string `json:"containedString"` // (case sensitive)
	Property              string `json:"property"`
	RestrictToProcessType string `json:"restrictToProcessType"` // Note: Not all types can be detected at startup.
}

func (me *ProcessGroupDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"contained_string": {
			Type:        hcl.TypeString,
			Description: "(case sensitive)",
			Required:    true,
		},
		"property": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"restrict_to_process_type": {
			Type:        hcl.TypeString,
			Description: "Note: Not all types can be detected at startup.",
			Optional:    true,
		},
	}
}

func (me *ProcessGroupDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"contained_string":         me.ContainedString,
		"property":                 me.Property,
		"restrict_to_process_type": me.RestrictToProcessType,
	})
}

func (me *ProcessGroupDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"contained_string":         &me.ContainedString,
		"property":                 &me.Property,
		"restrict_to_process_type": &me.RestrictToProcessType,
	})
}
