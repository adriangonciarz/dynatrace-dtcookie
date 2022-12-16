package osservicesmonitoring

import (
	"github.com/dtcookie/hcl"
)

type WindowsDetectionConditions []*WindowsDetectionCondition

func (me *WindowsDetectionConditions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"detectionConditionsWindow": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(WindowsDetectionCondition).Schema()},
		},
	}
}

func (me WindowsDetectionConditions) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("detectionConditionsWindow", me)
}

func (me *WindowsDetectionConditions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("detectionConditionsWindow", me)
}

type WindowsDetectionCondition struct {
	Condition        string              `json:"condition"`        // This string has to match a required format. See [OS services monitoring](https://dt-url.net/vl03xzk).\n\n- `$contains(ssh)` – Matches if `ssh` appears anywhere in the service's property value.\n- `$eq(sshd)` – Matches if `sshd` matches the service's property value exactly.\n- `$prefix(ss)` – Matches if `ss` matches the prefix of the service's property value.\n- `$suffix(hd)` – Matches if `hd` matches the suffix of the service's property value.\n\nAvailable logic operations:\n- `$not($eq(sshd))` – Matches if the service's property value is different from `sshd`.\n- `$and($prefix(ss),$suffix(hd))` – Matches if service's property value starts with `ss` and ends with `hd`.\n- `$or($prefix(ss),$suffix(hd))` – Matches if service's property value starts with `ss` or ends with `hd`.
	Property         WindowsServiceProps `json:"property"`         // Service property
	StartupCondition string              `json:"startupCondition"` // This string has to match a required format. See [OS services monitoring](https://dt-url.net/vl03xzk).\n\n- `$eq(manual)` – Matches services that are started manually.\n\nAvailable logic operations:\n- `$not($eq(auto))` – Matches services with startup type different from Automatic.\n- `$or($eq(auto),$eq(manual))` – Matches if service's startup type is either Automatic or Manual.\n\nUse one of the following values as a parameter for this condition:\n\n- `manual` for Manual\n- `manual_trigger` for Manual (Trigger Start)\n- `auto` for Automatic\n- `auto_delay` for Automatic (Delayed Start)\n- `auto_trigger` for Automatic (Trigger Start)\n- `auto_delay_trigger` for Automatic (Delayed Start, Trigger Start)\n- `disabled` for Disabled
}

func (me *WindowsDetectionCondition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeString,
			Description: "This string has to match a required format. See [OS services monitoring](https://dt-url.net/vl03xzk).\n\n- `$contains(ssh)` – Matches if `ssh` appears anywhere in the service's property value.\n- `$eq(sshd)` – Matches if `sshd` matches the service's property value exactly.\n- `$prefix(ss)` – Matches if `ss` matches the prefix of the service's property value.\n- `$suffix(hd)` – Matches if `hd` matches the suffix of the service's property value.\n\nAvailable logic operations:\n- `$not($eq(sshd))` – Matches if the service's property value is different from `sshd`.\n- `$and($prefix(ss),$suffix(hd))` – Matches if service's property value starts with `ss` and ends with `hd`.\n- `$or($prefix(ss),$suffix(hd))` – Matches if service's property value starts with `ss` or ends with `hd`.",
			Required:    true,
		},
		"property": {
			Type:        hcl.TypeString,
			Description: "Service property",
			Required:    true,
		},
		"startup_condition": {
			Type:        hcl.TypeString,
			Description: "This string has to match a required format. See [OS services monitoring](https://dt-url.net/vl03xzk).\n\n- `$eq(manual)` – Matches services that are started manually.\n\nAvailable logic operations:\n- `$not($eq(auto))` – Matches services with startup type different from Automatic.\n- `$or($eq(auto),$eq(manual))` – Matches if service's startup type is either Automatic or Manual.\n\nUse one of the following values as a parameter for this condition:\n\n- `manual` for Manual\n- `manual_trigger` for Manual (Trigger Start)\n- `auto` for Automatic\n- `auto_delay` for Automatic (Delayed Start)\n- `auto_trigger` for Automatic (Trigger Start)\n- `auto_delay_trigger` for Automatic (Delayed Start, Trigger Start)\n- `disabled` for Disabled",
			Required:    true,
		},
	}
}

func (me *WindowsDetectionCondition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"condition":         me.Condition,
		"property":          me.Property,
		"startup_condition": me.StartupCondition,
	})
}

func (me *WindowsDetectionCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"condition":         &me.Condition,
		"property":          &me.Property,
		"startup_condition": &me.StartupCondition,
	})
}
