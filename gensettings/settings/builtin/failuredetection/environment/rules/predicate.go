package rules

import "github.com/dtcookie/hcl"

type Predicate struct {
	CaseSensitive   bool           `json:"caseSensitive"`   // Case sensitive
	ManagementZones []string       `json:"managementZones"` // Management zones
	PredicateType   string         `json:"predicateType"`   // Predicate type
	ServiceType     []ServiceTypes `json:"serviceType"`     // Service types
	TagKeys         []string       `json:"tagKeys"`         // Tag keys
	Tags            []string       `json:"tags"`            // Tags (exact match)
	TextValues      []string       `json:"textValues"`      // Names
}

func (me *Predicate) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"case_sensitive": {
			Type:        hcl.TypeBool,
			Description: "Case sensitive",
			Optional:    true,
		},
		"management_zones": {
			Type:        hcl.TypeSet,
			Description: "Management zones",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"predicate_type": {
			Type:        hcl.TypeString,
			Description: "Predicate type",
			Required:    true,
		},
		"service_type": {
			Type:        hcl.TypeSet,
			Description: "Service types",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"tag_keys": {
			Type:        hcl.TypeSet,
			Description: "Tag keys",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"tags": {
			Type:        hcl.TypeSet,
			Description: "Tags (exact match)",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"text_values": {
			Type:        hcl.TypeSet,
			Description: "Names",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Predicate) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"case_sensitive":   me.CaseSensitive,
		"management_zones": me.ManagementZones,
		"predicate_type":   me.PredicateType,
		"service_type":     me.ServiceType,
		"tag_keys":         me.TagKeys,
		"tags":             me.Tags,
		"text_values":      me.TextValues,
	})
}

func (me *Predicate) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"case_sensitive":   &me.CaseSensitive,
		"management_zones": &me.ManagementZones,
		"predicate_type":   &me.PredicateType,
		"service_type":     &me.ServiceType,
		"tag_keys":         &me.TagKeys,
		"tags":             &me.Tags,
		"text_values":      &me.TextValues,
	})
}
