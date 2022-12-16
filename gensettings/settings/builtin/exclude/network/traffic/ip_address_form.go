package traffic

import (
	"github.com/dtcookie/hcl"
)

type IpAddressForms []*IpAddressForm

func (me *IpAddressForms) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"ip_address_form": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(IpAddressForm).Schema()},
		},
	}
}

func (me IpAddressForms) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("ip_address_form", me)
}

func (me *IpAddressForms) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("ip_address_form", me)
}

type IpAddressForm struct {
	IpAddress string `json:"ipAddress"` // IP address
}

func (me *IpAddressForm) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"ip_address": {
			Type:        hcl.TypeString,
			Description: "IP address",
			Required:    true,
		},
	}
}

func (me *IpAddressForm) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"ip_address": me.IpAddress,
	})
}

func (me *IpAddressForm) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"ip_address": &me.IpAddress,
	})
}
