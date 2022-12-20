package iam

import "github.com/dtcookie/hcl"

type Group struct {
	Name                     string      `json:"name"`
	Description              string      `json:"description"`
	FederatedAttributeValues []string    `json:"federatedAttributeValues"`
	Permissions              Permissions `json:"-"`
}

func (me *Group) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:     hcl.TypeString,
			Required: true,
		},
		"description": {
			Type:     hcl.TypeString,
			Optional: true,
		},
		"federated_attribute_values": {
			Type:     hcl.TypeSet,
			Optional: true,
			MinItems: 1,
			Elem:     &hcl.Schema{Type: hcl.TypeString},
		},
		"permissions": {
			Type:     hcl.TypeList,
			Optional: true,
			MinItems: 1,
			MaxItems: 1,
			Elem:     &hcl.Resource{Schema: new(Permissions).Schema()},
		},
	}
}

func (me *Group) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"name":                       me.Name,
		"description":                me.Description,
		"federated_attribute_values": me.FederatedAttributeValues,
		"permissions":                me.Permissions,
	})
}

func (me *Group) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"name":                       &me.Name,
		"description":                &me.Description,
		"federated_attribute_values": &me.FederatedAttributeValues,
		"permissions":                &me.Permissions,
	})
}
