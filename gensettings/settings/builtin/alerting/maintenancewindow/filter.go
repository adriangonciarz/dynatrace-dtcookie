package maintenancewindow

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Filters []*Filter

func (me *Filters) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"filter": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(Filter).Schema()},
		},
	}
}

func (me Filters) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me) > 0 {
		entries := []interface{}{}
		for _, entry := range me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["filter"] = entries
	}
	return result, nil
}

func (me *Filters) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("filter"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(Filter)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "filter", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

// Filter. Configured values of one filter are evaluated together (**AND**).. The maintenance window is applied onto an entity if it matches all of the values of at least one filter.
type Filter struct {
	EntityID        string   `json:"entityId"`        // A specific entity that should match this maintenance window.. **Note**: If an entity type filter value is set, it must be equal to the type of the selected entity. Otherwise this maintenance window will not match.
	EntityTags      []string `json:"entityTags"`      // Entities which contain all of the configured tags will match this maintenance window.
	EntityType      string   `json:"entityType"`      // Type of entities this maintenance window should match.. If no entity type is selected all entities regardless of the type will match.
	ManagementZones []string `json:"managementZones"` // Entities which are part of all the configured management zones will match this maintenance window.
}

func (me *Filter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"entity_id": {
			Type:        hcl.TypeString,
			Description: "A specific entity that should match this maintenance window.. **Note**: If an entity type filter value is set, it must be equal to the type of the selected entity. Otherwise this maintenance window will not match.",
			Optional:    true,
		},
		"entity_tags": {
			Type:        hcl.TypeSet,
			Description: "Entities which contain all of the configured tags will match this maintenance window.",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"entity_type": {
			Type:        hcl.TypeString,
			Description: "Type of entities this maintenance window should match.. If no entity type is selected all entities regardless of the type will match.",
			Optional:    true,
		},
		"management_zones": {
			Type:        hcl.TypeSet,
			Description: "Entities which are part of all the configured management zones will match this maintenance window.",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Filter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"entity_id":        me.EntityID,
		"entity_tags":      me.EntityTags,
		"entity_type":      me.EntityType,
		"management_zones": me.ManagementZones,
	})
}

func (me *Filter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"entity_id":        &me.EntityID,
		"entity_tags":      &me.EntityTags,
		"entity_type":      &me.EntityType,
		"management_zones": &me.ManagementZones,
	})
}
