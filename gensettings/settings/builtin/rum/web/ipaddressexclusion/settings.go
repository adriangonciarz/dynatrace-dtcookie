package ipaddressexclusion

import "github.com/dtcookie/hcl"

type Settings struct {
	IpAddressExclusionInclude bool                    `json:"ipAddressExclusionInclude"` // These are the only IP addresses that should be monitored
	IpExclusionList           IpAddressExclusionRules `json:"ipExclusionList"`           // **Examples:**\n\n   - 84.112.10.5\n   - fe80::10a1:c6b2:5f68:785d
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"ip_address_exclusion_include": {
			Type:        hcl.TypeBool,
			Description: "These are the only IP addresses that should be monitored",
			Optional:    true,
		},
		"ip_exclusion_list": {
			Type:        hcl.TypeList,
			Description: "**Examples:**\n\n   - 84.112.10.5\n   - fe80::10a1:c6b2:5f68:785d",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(IpAddressExclusionRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"ip_address_exclusion_include": me.IpAddressExclusionInclude,
		"ip_exclusion_list":            me.IpExclusionList,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"ip_address_exclusion_include": &me.IpAddressExclusionInclude,
		"ip_exclusion_list":            &me.IpExclusionList,
	})
}
