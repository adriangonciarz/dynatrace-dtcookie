package monitoringrule

import "github.com/dtcookie/hcl"

type ContainerCondition struct {
	Operator ConditionOperator `json:"operator"` // Condition operator
	Property ContainerItem     `json:"property"` // Container property
	Value    string            `json:"value"`    // Condition value
}

func (me *ContainerCondition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"operator": {
			Type:        hcl.TypeString,
			Description: "Condition operator",
			Required:    true,
		},
		"property": {
			Type:        hcl.TypeString,
			Description: "Container property",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "Condition value",
			Required:    true,
		},
	}
}

func (me *ContainerCondition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"operator": me.Operator,
		"property": me.Property,
		"value":    me.Value,
	})
}

func (me *ContainerCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"operator": &me.Operator,
		"property": &me.Property,
		"value":    &me.Value,
	})
}
