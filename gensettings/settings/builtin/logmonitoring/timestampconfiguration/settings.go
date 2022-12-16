package timestampconfiguration

import "github.com/dtcookie/hcl"

type Settings struct {
	Config_item_title string   `json:"config-item-title"` // Name
	Date_time_pattern string   `json:"date-time-pattern"` // Date-time pattern
	Enabled           bool     `json:"enabled"`           // Active
	Matchers          Matchers `json:"matchers"`
	Timezone          string   `json:"timezone"` // Timezone
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"config_-_item_-_title": {
			Type:        hcl.TypeString,
			Description: "Name",
			Required:    true,
		},
		"date_-_time_-_pattern": {
			Type:        hcl.TypeString,
			Description: "Date-time pattern",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Active",
			Optional:    true,
		},
		"matchers": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Matchers).Schema()},
			MinItems:    1,
			MaxItems:    1,
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
		"config_-_item_-_title": me.Config_item_title,
		"date_-_time_-_pattern": me.Date_time_pattern,
		"enabled":               me.Enabled,
		"matchers":              me.Matchers,
		"timezone":              me.Timezone,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"config_-_item_-_title": &me.Config_item_title,
		"date_-_time_-_pattern": &me.Date_time_pattern,
		"enabled":               &me.Enabled,
		"matchers":              &me.Matchers,
		"timezone":              &me.Timezone,
	})
}
