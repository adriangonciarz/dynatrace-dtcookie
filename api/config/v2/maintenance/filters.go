package maintenance

import "github.com/dtcookie/hcl"

type Filters []*Filter

type Filter struct {
	EntityType      *string  `json:"entityType,omitempty"`      // Type of entities this maintenance window should match
	EntityId        *string  `json:"entityId,omitempty"`        // A specific entity that should match this maintenance window
	EntityTags      []string `json:"entityTags,omitempty"`      // The tags you want to use for matching in the format key or key:value
	ManagementZones []string `json:"managementZones,omitempty"` // The IDs of management zones to which the matched entities must belong
}

func (me *Filter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"entity_type": {
			Type:        hcl.TypeString,
			Description: "Type of entities this maintenance window should match",
			Optional:    true,
		},
		"entity_id": {
			Type:        hcl.TypeString,
			Description: "A specific entity that should match this maintenance window",
			Optional:    true,
		},
		"entity_tags": {
			Type:        hcl.TypeSet,
			Description: "The tags you want to use for matching in the format key or key:value",
			Optional:    true,
			MinItems:    1,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"management_zones": {
			Type:        hcl.TypeSet,
			Description: "The IDs of management zones to which the matched entities must belong",
			Optional:    true,
			MinItems:    1,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Filter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"entity_type":      me.EntityType,
		"entity_id":        me.EntityId,
		"entity_tags":      me.EntityTags,
		"management_zones": me.ManagementZones,
	})
}

func (me *Filter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"entity_type":      &me.EntityType,
		"entity_id":        &me.EntityId,
		"entity_tags":      &me.EntityTags,
		"management_zones": &me.ManagementZones,
	})
}

func (me *Filters) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"filter": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "A list of matching rules for dynamic filter formation.  If several rules are set, the OR logic applies",
			Elem:        &hcl.Resource{Schema: new(Filter).Schema()},
		},
	}
}

func (me Filters) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("filter", me)
}

func (me *Filters) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("filter", me)
}
