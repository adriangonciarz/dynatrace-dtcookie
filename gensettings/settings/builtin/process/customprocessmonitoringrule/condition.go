package customprocessmonitoringrule

import "github.com/dtcookie/hcl"

type Condition struct {
	EnvVar   string            `json:"envVar"`   // supported only with OneAgent 1.167+
	Item     AgentItemName     `json:"item"`     // Condition target
	Operator ConditionOperator `json:"operator"` // Condition operator
	Value    string            `json:"value"`    // Condition value
}

func (me *Condition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"env_var": {
			Type:        hcl.TypeString,
			Description: "supported only with OneAgent 1.167+",
			Required:    true,
		},
		"item": {
			Type:        hcl.TypeString,
			Description: "Condition target",
			Required:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "Condition operator",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "Condition value",
			Required:    true,
		},
	}
}

func (me *Condition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"env_var":  me.EnvVar,
		"item":     me.Item,
		"operator": me.Operator,
		"value":    me.Value,
	})
}

func (me *Condition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"env_var":  &me.EnvVar,
		"item":     &me.Item,
		"operator": &me.Operator,
		"value":    &me.Value,
	})
}
