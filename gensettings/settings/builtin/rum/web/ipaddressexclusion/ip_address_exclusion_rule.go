package ipaddressexclusion

import (
	"github.com/dtcookie/hcl"
)

type IpAddressExclusionRules []*IpAddressExclusionRule

func (me *IpAddressExclusionRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"ipExclusion": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(IpAddressExclusionRule).Schema()},
		},
	}
}

func (me IpAddressExclusionRules) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("ipExclusion", me)
}

func (me *IpAddressExclusionRules) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("ipExclusion", me)
}

type IpAddressExclusionRule struct {
	Ip   string `json:"ip"`   // Single IP or IP range start address
	IpTo string `json:"ipTo"` // IP range end
}

func (me *IpAddressExclusionRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"ip": {
			Type:        hcl.TypeString,
			Description: "Single IP or IP range start address",
			Required:    true,
		},
		"ip_to": {
			Type:        hcl.TypeString,
			Description: "IP range end",
			Optional:    true,
		},
	}
}

func (me *IpAddressExclusionRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"ip":    me.Ip,
		"ip_to": me.IpTo,
	})
}

func (me *IpAddressExclusionRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"ip":    &me.Ip,
		"ip_to": &me.IpTo,
	})
}
