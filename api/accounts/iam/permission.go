package iam

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Permissions []*Permission

func (me *Permissions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"permission": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "A Permission",
			Elem:        &hcl.Resource{Schema: new(Permission).Schema()},
		},
	}
}

func (me *Permissions) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(*me) > 0 {
		entries := []interface{}{}
		for _, entry := range *me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["permission"] = entries
	}
	return result, nil
}

func (me *Permissions) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("permission"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(Permission)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "permission", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type Permission struct {
	Name      string `json:"permissionName"` // The name of the permission. Possible values are account-company-info, account-user-management, account-viewer, tenant-viewer, tenant-manage-settings, tenant-agent-install, tenant-logviewer, tenant-view-sensitive-request-data, tenant-configure-request-capture-data, tenant-replay-sessions-with-masking, tenant-replay-sessions-without-masking, tenant-manage-security-problems, tenant-manage-support-tickets.
	Scope     string `json:"scope"`          // The scope of the permission. Depending on the scope type, it is defined by the UUID of the account (scopeType = `account`), the ID of the environment (scopeType = `tenant`) or the ID of the management zone from an environment in `{environment-id}-{management-zone-id}` format (scopeType = `management-zone`)
	ScopeType string `json:"scopeType"`      // The type of the permission scope. Possible values are `account`, `tenant` and `management-zone`
}

func (me *Permission) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:     hcl.TypeString,
			Required: true,
		},
		"scope": {
			Type:     hcl.TypeString,
			Required: true,
		},
		"type": {
			Type:     hcl.TypeString,
			Required: true,
		},
	}
}

func (me *Permission) MarshalHCL() (map[string]interface{}, error) {
	return map[string]interface{}{
		"name":  me.Name,
		"scope": me.Scope,
		"type":  string(me.ScopeType),
	}, nil
}

func (me *Permission) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"name":  &me.Name,
		"scope": &me.Scope,
		"type":  &me.ScopeType,
	})
}

type ScopeType string

var ScopeTypes = struct {
	Account        ScopeType
	Tenant         ScopeType
	ManagementZone ScopeType
}{
	ScopeType("account"),
	ScopeType("tenant"),
	ScopeType("management-zone"),
}
