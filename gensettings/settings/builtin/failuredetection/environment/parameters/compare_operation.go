package parameters

import "github.com/dtcookie/hcl"

type CompareOperation struct {
	CaseSensitive        bool    `json:"caseSensitive"`        // Case sensitive
	CompareOperationType string  `json:"compareOperationType"` // Apply this comparison
	DoubleValue          float64 `json:"doubleValue"`          // Value
	IntValue             int     `json:"intValue"`             // Value
	TextValue            string  `json:"textValue"`            // Value
}

func (me *CompareOperation) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"case_sensitive": {
			Type:        hcl.TypeBool,
			Description: "Case sensitive",
			Optional:    true,
		},
		"compare_operation_type": {
			Type:        hcl.TypeString,
			Description: "Apply this comparison",
			Required:    true,
		},
		"double_value": {
			Type:        hcl.TypeFloat,
			Description: "Value",
			Required:    true,
		},
		"int_value": {
			Type:        hcl.TypeInt,
			Description: "Value",
			Required:    true,
		},
		"text_value": {
			Type:        hcl.TypeString,
			Description: "Value",
			Required:    true,
		},
	}
}

func (me *CompareOperation) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"case_sensitive":         me.CaseSensitive,
		"compare_operation_type": me.CompareOperationType,
		"double_value":           me.DoubleValue,
		"int_value":              me.IntValue,
		"text_value":             me.TextValue,
	})
}

func (me *CompareOperation) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"case_sensitive":         &me.CaseSensitive,
		"compare_operation_type": &me.CompareOperationType,
		"double_value":           &me.DoubleValue,
		"int_value":              &me.IntValue,
		"text_value":             &me.TextValue,
	})
}
