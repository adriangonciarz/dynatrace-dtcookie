package node

import "github.com/dtcookie/hcl"

type ReadinessIssues struct {
	Configuration *ReadinessIssuesConfig `json:"configuration"` // Alert if
	Enabled       bool                   `json:"enabled"`       // Evaluates on node condition 'Ready'
}

func (me *ReadinessIssues) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"configuration": {
			Type:        hcl.TypeList,
			Description: "Alert if",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ReadinessIssuesConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Evaluates on node condition 'Ready'",
			Optional:    true,
		},
	}
}

func (me *ReadinessIssues) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"configuration": me.Configuration,
		"enabled":       me.Enabled,
	})
}

func (me *ReadinessIssues) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"configuration": &me.Configuration,
		"enabled":       &me.Enabled,
	})
}
