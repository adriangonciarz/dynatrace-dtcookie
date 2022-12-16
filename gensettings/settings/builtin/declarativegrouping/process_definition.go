package declarativegrouping

import (
	"github.com/dtcookie/hcl"
)

type ProcessDefinitions []*ProcessDefinition

func (me *ProcessDefinitions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"process_definition": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(ProcessDefinition).Schema()},
		},
	}
}

func (me ProcessDefinitions) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("process_definition", me)
}

func (me *ProcessDefinitions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("process_definition", me)
}

type ProcessDefinition struct {
	ID               string              `json:"id"`               // Process group identifier
	ProcessGroupName string              `json:"processGroupName"` // This identifier is used by Dynatrace to recognize this process group.
	Rules            DetectionConditions `json:"rules"`            // Define process detection rules by selecting a process property and a condition. Each process group can have multiple detection rules associated with it.
}

func (me *ProcessDefinition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"id": {
			Type:        hcl.TypeString,
			Description: "Process group identifier",
			Required:    true,
		},
		"process_group_name": {
			Type:        hcl.TypeString,
			Description: "This identifier is used by Dynatrace to recognize this process group.",
			Required:    true,
		},
		"rules": {
			Type:        hcl.TypeList,
			Description: "Define process detection rules by selecting a process property and a condition. Each process group can have multiple detection rules associated with it.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DetectionConditions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *ProcessDefinition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"id":                 me.ID,
		"process_group_name": me.ProcessGroupName,
		"rules":              me.Rules,
	})
}

func (me *ProcessDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"id":                 &me.ID,
		"process_group_name": &me.ProcessGroupName,
		"rules":              &me.Rules,
	})
}
