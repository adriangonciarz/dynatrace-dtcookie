package processavailability

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled  bool                `json:"enabled"`  // Enabled
	Metadata MetadataItems       `json:"metadata"` // Set of additional key-value properties to be attached to the triggered event.
	Name     string              `json:"name"`     // Monitored rule name
	Rules    DetectionConditions `json:"rules"`    // Define process detection rules by selecting a process property and a condition. Each monitoring rule can have multiple detection rules associated with it.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"metadata": {
			Type:        hcl.TypeList,
			Description: "Set of additional key-value properties to be attached to the triggered event.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MetadataItems).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "Monitored rule name",
			Required:    true,
		},
		"rules": {
			Type:        hcl.TypeList,
			Description: "Define process detection rules by selecting a process property and a condition. Each monitoring rule can have multiple detection rules associated with it.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DetectionConditions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":  me.Enabled,
		"metadata": me.Metadata,
		"name":     me.Name,
		"rules":    me.Rules,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":  &me.Enabled,
		"metadata": &me.Metadata,
		"name":     &me.Name,
		"rules":    &me.Rules,
	})
}
