package monitoringrule

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled  bool              `json:"enabled"`  // Enabled
	Mode     MonitoringMode    `json:"mode"`     // Mode
	Operator ConditionOperator `json:"operator"` // Condition operator
	Property ContainerItem     `json:"property"` // Container property
	Value    string            `json:"value"`    // Condition value
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"mode": {
			Type:        hcl.TypeString,
			Description: "Mode",
			Required:    true,
		},
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

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":  me.Enabled,
		"mode":     me.Mode,
		"operator": me.Operator,
		"property": me.Property,
		"value":    me.Value,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":  &me.Enabled,
		"mode":     &me.Mode,
		"operator": &me.Operator,
		"property": &me.Property,
		"value":    &me.Value,
	})
}
