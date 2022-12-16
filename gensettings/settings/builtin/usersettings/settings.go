package usersettings

import "github.com/dtcookie/hcl"

type Settings struct {
	Language Language `json:"language"` // Language
	Region   string   `json:"region"`   // Region
	Theme    Theme    `json:"theme"`    // Page refresh required to view changes
	Timezone string   `json:"timezone"` // Timezone
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"language": {
			Type:        hcl.TypeString,
			Description: "Language",
			Required:    true,
		},
		"region": {
			Type:        hcl.TypeString,
			Description: "Region",
			Required:    true,
		},
		"theme": {
			Type:        hcl.TypeString,
			Description: "Page refresh required to view changes",
			Required:    true,
		},
		"timezone": {
			Type:        hcl.TypeString,
			Description: "Timezone",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"language": me.Language,
		"region":   me.Region,
		"theme":    me.Theme,
		"timezone": me.Timezone,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"language": &me.Language,
		"region":   &me.Region,
		"theme":    &me.Theme,
		"timezone": &me.Timezone,
	})
}
