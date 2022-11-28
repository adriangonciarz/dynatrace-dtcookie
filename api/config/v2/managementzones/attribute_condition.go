package managementzones

import "github.com/dtcookie/hcl"

// No documentation available
type AttributeCondition struct {
	EnumValue        string    `json:"enumValue" hcl:"enum_value"`                // Value
	Tag              string    `json:"tag" hcl:"tag"`                             // Tag. Format: `[CONTEXT]tagKey:tagValue`
	Key              Attribute `json:"key" hcl:"key"`                             // Property
	DynamicKey       string    `json:"dynamicKey" hcl:"dynamic_key"`              // Dynamic key
	CaseSensitive    bool      `json:"caseSensitive" hcl:"case_sensitive"`        // Case sensitive
	IntegerValue     int       `json:"integerValue" hcl:"integer_value"`          // Value
	EntityID         string    `json:"entityId" hcl:"entity_id"`                  // Value
	Operator         Operator  `json:"operator" hcl:"operator"`                   // Operator
	StringValue      string    `json:"stringValue" hcl:"string_value"`            // Value
	DynamicKeySource string    `json:"dynamicKeySource" hcl:"dynamic_key_source"` // Key source
}

func (me *AttributeCondition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enum_value": {
			Type:        hcl.TypeString,
			Description: "Value",
			Required:    true,
		},
		"tag": {
			Type:        hcl.TypeString,
			Description: "Tag. Format: `[CONTEXT]tagKey:tagValue`",
			Required:    true,
		},
		"key": {
			Type:        hcl.TypeString,
			Description: "Property",
			Required:    true,
		},
		"dynamic_key": {
			Type:        hcl.TypeString,
			Description: "Dynamic key",
			Required:    true,
		},
		"case_sensitive": {
			Type:        hcl.TypeBool,
			Description: "Case sensitive",
			Optional:    true,
		},
		"integer_value": {
			Type:        hcl.TypeInt,
			Description: "Value",
			Required:    true,
		},
		"entity_id": {
			Type:        hcl.TypeString,
			Description: "Value",
			Required:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "Operator",
			Required:    true,
		},
		"string_value": {
			Type:        hcl.TypeString,
			Description: "Value",
			Required:    true,
		},
		"dynamic_key_source": {
			Type:        hcl.TypeString,
			Description: "Key source",
			Required:    true,
		},
	}
}
