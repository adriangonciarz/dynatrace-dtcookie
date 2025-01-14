package sharing

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// SharePermission represents access permissions of the dashboard
type SharePermission struct {
	ID         *string        `json:"id,omitempty"` // The ID of the user or group to whom the permission is granted.\n\nNot applicable if the **type** is set to `ALL`
	Type       PermissionType `json:"type"`         // The type of the permission: \n\n* `USER`: The dashboard is shared with the specified user. \n* `GROUP`: The dashboard is shared with all users of the specified group. \n* `ALL`: The dashboard is shared via link. Any authenticated user with the link can view the dashboard
	Permission Permission     `json:"permission"`   // The level of the permission: \n \n* `VIEW`: The dashboard is shared with read permission. \n* `EDIT`: The dashboard is shared with edit permission
}

func (me *SharePermission) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"id": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The ID of the user or group to whom the permission is granted.\n\nNot applicable if the **type** is set to `ALL`",
		},
		"type": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The type of the permission: \n\n* `USER`: The dashboard is shared with the specified user. \n* `GROUP`: The dashboard is shared with all users of the specified group. \n* `ALL`: The dashboard is shared via link. Any authenticated user with the link can view the dashboard",
		},
		"level": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The level of the permission: \n \n* `VIEW`: The dashboard is shared with read permission. \n* `EDIT`: The dashboard is shared with edit permission",
		},
	}
}

// MarshalHCL has no documentation
func (me *SharePermission) MarshalHCL() (map[string]interface{}, error) {
	m := map[string]interface{}{}
	if me.ID != nil {
		m["id"] = *me.ID
	}
	m["type"] = string(me.Type)
	m["level"] = string(me.Permission)
	return m, nil
}

// UnmarshalHCL has no documentation
func (me *SharePermission) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("id"); ok {
		me.ID = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("type"); ok {
		me.Type = PermissionType(value.(string))
	}
	if value, ok := decoder.GetOk("level"); ok {
		me.Permission = Permission(value.(string))
	}
	return nil
}

type SharePermissions []*SharePermission

func (me *SharePermissions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"permission": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(SharePermission).Schema()},
			Description: "Access permissions of the dashboard",
		},
	}
}

// MarshalHCL has no documentation
func (me SharePermissions) MarshalHCL() (map[string]interface{}, error) {
	props := hcl.Properties{}
	return props.EncodeSlice("permission", me)
}

// UnmarshalHCL has no documentation
func (me *SharePermissions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("permission", me)
}

type PermissionType string

var PermissionTypes = struct {
	All   PermissionType
	Group PermissionType
	User  PermissionType
}{
	All:   PermissionType("ALL"),
	Group: PermissionType("GROUP"),
	User:  PermissionType("USER"),
}

type Permission string

var Permissions = struct {
	Edit Permission
	View Permission
}{
	Edit: Permission("EDIT"),
	View: Permission("VIEW"),
}
