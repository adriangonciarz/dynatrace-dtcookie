package managementzones

import "github.com/dtcookie/hcl"

// No documentation available
type DimensionCondition struct {
	ConditionType DimensionConditionType `json:"conditionType" hcl:"condition_type"` // Type
	Key           string                 `json:"key" hcl:"key"`                      // Key
	RuleMatcher   DimensionOperator      `json:"ruleMatcher" hcl:"rule_matcher"`     // Operator
	Value         string                 `json:"value" hcl:"value"`                  // Value
}

func (me *DimensionCondition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition_type": {
			Type:        hcl.TypeString,
			Description: "Type",
			Required:    true,
		},
		"key": {
			Type:        hcl.TypeString,
			Description: "Key",
			Required:    true,
		},
		"rule_matcher": {
			Type:        hcl.TypeString,
			Description: "Operator",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "Value",
			Required:    true,
		},
	}
}
