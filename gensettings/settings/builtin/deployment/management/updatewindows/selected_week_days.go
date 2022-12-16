package updatewindows

import "github.com/dtcookie/hcl"

type SelectedWeekDays struct {
	Friday    bool `json:"friday"`
	Monday    bool `json:"monday"`
	Saturday  bool `json:"saturday"`
	Sunday    bool `json:"sunday"`
	Thursday  bool `json:"thursday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
}

func (me *SelectedWeekDays) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"friday": {
			Type:        hcl.TypeBool,
			Description: "no documentation available",
			Optional:    true,
		},
		"monday": {
			Type:        hcl.TypeBool,
			Description: "no documentation available",
			Optional:    true,
		},
		"saturday": {
			Type:        hcl.TypeBool,
			Description: "no documentation available",
			Optional:    true,
		},
		"sunday": {
			Type:        hcl.TypeBool,
			Description: "no documentation available",
			Optional:    true,
		},
		"thursday": {
			Type:        hcl.TypeBool,
			Description: "no documentation available",
			Optional:    true,
		},
		"tuesday": {
			Type:        hcl.TypeBool,
			Description: "no documentation available",
			Optional:    true,
		},
		"wednesday": {
			Type:        hcl.TypeBool,
			Description: "no documentation available",
			Optional:    true,
		},
	}
}

func (me *SelectedWeekDays) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"friday":    me.Friday,
		"monday":    me.Monday,
		"saturday":  me.Saturday,
		"sunday":    me.Sunday,
		"thursday":  me.Thursday,
		"tuesday":   me.Tuesday,
		"wednesday": me.Wednesday,
	})
}

func (me *SelectedWeekDays) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"friday":    &me.Friday,
		"monday":    &me.Monday,
		"saturday":  &me.Saturday,
		"sunday":    &me.Sunday,
		"thursday":  &me.Thursday,
		"tuesday":   &me.Tuesday,
		"wednesday": &me.Wednesday,
	})
}
