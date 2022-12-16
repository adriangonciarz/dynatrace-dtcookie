package infrastructurehosts

import "github.com/dtcookie/hcl"

type Settings struct {
	Host    *Host    `json:"host"`    // Hosts
	Network *Network `json:"network"` // Network
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"host": {
			Type:        hcl.TypeList,
			Description: "Hosts",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Host).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"network": {
			Type:        hcl.TypeList,
			Description: "Network",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Network).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"host":    me.Host,
		"network": me.Network,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"host":    &me.Host,
		"network": &me.Network,
	})
}
