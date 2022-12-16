package environment

import "github.com/dtcookie/hcl"

type Settings struct {
	Name         string `json:"name"`         // Name
	NetworkScope Scope  `json:"networkScope"` // - External: The remote environment is located in an another network.\n- Internal: The remote environment is located in the same network.\n- Cluster: The remote environment is located in the same cluster.\n\nDynatrace SaaS can only use External.
	Token        string `json:"token"`        // Provide a valid token created on the remote environment.
	Uri          string `json:"uri"`          // Specify the full URI to the remote environment. Your local environment will have to be able to connect this URI on a network level.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "Name",
			Required:    true,
		},
		"network_scope": {
			Type:        hcl.TypeString,
			Description: "- External: The remote environment is located in an another network.\n- Internal: The remote environment is located in the same network.\n- Cluster: The remote environment is located in the same cluster.\n\nDynatrace SaaS can only use External.",
			Required:    true,
		},
		"token": {
			Type:        hcl.TypeString,
			Description: "Provide a valid token created on the remote environment.",
			Required:    true,
		},
		"uri": {
			Type:        hcl.TypeString,
			Description: "Specify the full URI to the remote environment. Your local environment will have to be able to connect this URI on a network level.",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"name":          me.Name,
		"network_scope": me.NetworkScope,
		"token":         me.Token,
		"uri":           me.Uri,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"name":          &me.Name,
		"network_scope": &me.NetworkScope,
		"token":         &me.Token,
		"uri":           &me.Uri,
	})
}
