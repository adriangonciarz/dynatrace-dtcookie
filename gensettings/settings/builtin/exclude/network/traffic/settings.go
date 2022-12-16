package traffic

import "github.com/dtcookie/hcl"

type Settings struct {
	ExcludeIp  IpAddressForms `json:"excludeIp"`  // Providing a host IP address, you will exclude network traffic only in calculating connectivity (other metrics will still be calculated).
	ExcludeNic NicForms       `json:"excludeNic"` // Selecting a network interface, you will exclude all network traffic on that interface from being monitored. You can select from the list below what to not monitor, or input it manually using the \"other one\" option.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"exclude_ip": {
			Type:        hcl.TypeList,
			Description: "Providing a host IP address, you will exclude network traffic only in calculating connectivity (other metrics will still be calculated).",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(IpAddressForms).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"exclude_nic": {
			Type:        hcl.TypeList,
			Description: "Selecting a network interface, you will exclude all network traffic on that interface from being monitored. You can select from the list below what to not monitor, or input it manually using the \"other one\" option.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(NicForms).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"exclude_ip":  me.ExcludeIp,
		"exclude_nic": me.ExcludeNic,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"exclude_ip":  &me.ExcludeIp,
		"exclude_nic": &me.ExcludeNic,
	})
}
