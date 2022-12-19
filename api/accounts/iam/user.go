package iam

import (
	"github.com/dtcookie/hcl"
)

type User struct {
	Email  string   `json:"email"`
	Groups []string `json:"-"`
}

func (me *User) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"email": {
			Type:     hcl.TypeString,
			Required: true,
		},
		"groups": {
			Type:     hcl.TypeSet,
			Optional: true,
			MinItems: 1,
			Elem:     &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *User) MarshalHCL() (map[string]interface{}, error) {
	return map[string]interface{}{
		"email":  me.Email,
		"groups": me.Groups,
	}, nil
}

func (me *User) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"email":  &me.Email,
		"groups": &me.Groups,
	})
}
