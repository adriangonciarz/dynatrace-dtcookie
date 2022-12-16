package customunit

import "github.com/dtcookie/hcl"

type Settings struct {
	AdvancedSettings      bool             `json:"advancedSettings"`      // If not enabled a simple custom unit with the information above is created. If the advanced custom unit option is enabled, dependencies to already existing units can be defined.
	FirstUnit             string           `json:"firstUnit"`             // The new unit depends on existing unit.
	NewUnitDescription    string           `json:"newUnitDescription"`    // Unit description should provide additional information about the new unit
	NewUnitName           string           `json:"newUnitName"`           // Unit name has to be unique and is used as identifier.
	NewUnitPluralName     string           `json:"newUnitPluralName"`     // Unit plural name represent the plural form of the unit name.
	NewUnitSymbol         string           `json:"newUnitSymbol"`         // Unit symbol has to be unique.
	PowerFactor           int              `json:"powerFactor"`           // (existing unit) ^ (exponent) = (new unit)
	ScalingFactor         float64          `json:"scalingFactor"`         // (existing unit) * (Scaling factor) = new unit
	SecondUnit            string           `json:"secondUnit"`            // New unit depends on existing unit (based on selection at unit composition).
	UnitCombinationSelect UnitCombinations `json:"unitCombinationSelect"` // Unit composition defines the way the new unit depends on already existing ones.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"advanced_settings": {
			Type:        hcl.TypeBool,
			Description: "If not enabled a simple custom unit with the information above is created. If the advanced custom unit option is enabled, dependencies to already existing units can be defined.",
			Optional:    true,
		},
		"first_unit": {
			Type:        hcl.TypeString,
			Description: "The new unit depends on existing unit.",
			Required:    true,
		},
		"new_unit_description": {
			Type:        hcl.TypeString,
			Description: "Unit description should provide additional information about the new unit",
			Required:    true,
		},
		"new_unit_name": {
			Type:        hcl.TypeString,
			Description: "Unit name has to be unique and is used as identifier.",
			Required:    true,
		},
		"new_unit_plural_name": {
			Type:        hcl.TypeString,
			Description: "Unit plural name represent the plural form of the unit name.",
			Required:    true,
		},
		"new_unit_symbol": {
			Type:        hcl.TypeString,
			Description: "Unit symbol has to be unique.",
			Required:    true,
		},
		"power_factor": {
			Type:        hcl.TypeInt,
			Description: "(existing unit) ^ (exponent) = (new unit)",
			Required:    true,
		},
		"scaling_factor": {
			Type:        hcl.TypeFloat,
			Description: "(existing unit) * (Scaling factor) = new unit",
			Required:    true,
		},
		"second_unit": {
			Type:        hcl.TypeString,
			Description: "New unit depends on existing unit (based on selection at unit composition).",
			Required:    true,
		},
		"unit_combination_select": {
			Type:        hcl.TypeString,
			Description: "Unit composition defines the way the new unit depends on already existing ones.",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"advanced_settings":       me.AdvancedSettings,
		"first_unit":              me.FirstUnit,
		"new_unit_description":    me.NewUnitDescription,
		"new_unit_name":           me.NewUnitName,
		"new_unit_plural_name":    me.NewUnitPluralName,
		"new_unit_symbol":         me.NewUnitSymbol,
		"power_factor":            me.PowerFactor,
		"scaling_factor":          me.ScalingFactor,
		"second_unit":             me.SecondUnit,
		"unit_combination_select": me.UnitCombinationSelect,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"advanced_settings":       &me.AdvancedSettings,
		"first_unit":              &me.FirstUnit,
		"new_unit_description":    &me.NewUnitDescription,
		"new_unit_name":           &me.NewUnitName,
		"new_unit_plural_name":    &me.NewUnitPluralName,
		"new_unit_symbol":         &me.NewUnitSymbol,
		"power_factor":            &me.PowerFactor,
		"scaling_factor":          &me.ScalingFactor,
		"second_unit":             &me.SecondUnit,
		"unit_combination_select": &me.UnitCombinationSelect,
	})
}
