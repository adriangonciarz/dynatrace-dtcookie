package simpledetectionrule

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled            bool              `json:"enabled"`            // Enabled
	GroupIdentifier    string            `json:"groupIdentifier"`    // If Dynatrace detects this property at startup of a process, it will use its value to identify process groups more granular.
	InstanceIdentifier string            `json:"instanceIdentifier"` // Use a variable to identify instances within a process group.\n\nThe type of variable is the same as selected in 'Property source'.
	ProcessType        *string           `json:"processType"`        // Note: Not all types can be detected at startup.
	RuleType           DetectionRuleType `json:"ruleType"`           // Source to use to separate processes into multiple process groups.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"group_identifier": {
			Type:        hcl.TypeString,
			Description: "If Dynatrace detects this property at startup of a process, it will use its value to identify process groups more granular.",
			Required:    true,
		},
		"instance_identifier": {
			Type:        hcl.TypeString,
			Description: "Use a variable to identify instances within a process group.\n\nThe type of variable is the same as selected in 'Property source'.",
			Required:    true,
		},
		"process_type": {
			Type:        hcl.TypeString,
			Description: "Note: Not all types can be detected at startup.",
			Optional:    true,
		},
		"rule_type": {
			Type:        hcl.TypeString,
			Description: "Source to use to separate processes into multiple process groups.",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":             me.Enabled,
		"group_identifier":    me.GroupIdentifier,
		"instance_identifier": me.InstanceIdentifier,
		"process_type":        me.ProcessType,
		"rule_type":           me.RuleType,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":             &me.Enabled,
		"group_identifier":    &me.GroupIdentifier,
		"instance_identifier": &me.InstanceIdentifier,
		"process_type":        &me.ProcessType,
		"rule_type":           &me.RuleType,
	})
}
