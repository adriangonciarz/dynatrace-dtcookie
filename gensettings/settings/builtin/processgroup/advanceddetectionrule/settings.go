package advanceddetectionrule

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled            bool                       `json:"enabled"`            // Enabled
	GroupExtraction    *ProcessGroupExtraction    `json:"groupExtraction"`    // You can define the properties that should be used to identify your process groups.
	InstanceExtraction *ProcessInstanceExtraction `json:"instanceExtraction"` // You can define the properties that should be used to identify your process instances.
	ProcessDetection   *ProcessGroupDetection     `json:"processDetection"`   // Apply this rule to processes where the selected property contains the specified string.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"group_extraction": {
			Type:        hcl.TypeList,
			Description: "You can define the properties that should be used to identify your process groups.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ProcessGroupExtraction).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"instance_extraction": {
			Type:        hcl.TypeList,
			Description: "You can define the properties that should be used to identify your process instances.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ProcessInstanceExtraction).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"process_detection": {
			Type:        hcl.TypeList,
			Description: "Apply this rule to processes where the selected property contains the specified string.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ProcessGroupDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":             me.Enabled,
		"group_extraction":    me.GroupExtraction,
		"instance_extraction": me.InstanceExtraction,
		"process_detection":   me.ProcessDetection,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":             &me.Enabled,
		"group_extraction":    &me.GroupExtraction,
		"instance_extraction": &me.InstanceExtraction,
		"process_detection":   &me.ProcessDetection,
	})
}
