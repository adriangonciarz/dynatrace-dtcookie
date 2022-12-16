package osservicesmonitoring

import (
	"github.com/dtcookie/hcl"
)

type LinuxDetectionConditions []*LinuxDetectionCondition

func (me *LinuxDetectionConditions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"linux_detection_condition": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(LinuxDetectionCondition).Schema()},
		},
	}
}

func (me LinuxDetectionConditions) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("linux_detection_condition", me)
}

func (me *LinuxDetectionConditions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("linux_detection_condition", me)
}

type LinuxDetectionCondition struct {
	Condition        string           `json:"condition"`        // This string has to match a required format. See [OS services monitoring](https://dt-url.net/vl03xzk).\n\n- `$contains(ssh)` – Matches if `ssh` appears anywhere in the service's property value.\n- `$eq(sshd)` – Matches if `sshd` matches the service's property value exactly.\n- `$prefix(ss)` – Matches if `ss` matches the prefix of the service's property value.\n- `$suffix(hd)` – Matches if `hd` matches the suffix of the service's property value.\n\nAvailable logic operations:\n- `$not($eq(sshd))` – Matches if the service's property value is different from `sshd`.\n- `$and($prefix(ss),$suffix(hd))` – Matches if service's property value starts with `ss` and ends with `hd`.\n- `$or($prefix(ss),$suffix(hd))` – Matches if service's property value starts with `ss` or ends with `hd`.
	Property         LinuxServiceProp `json:"property"`         // Service property
	StartupCondition string           `json:"startupCondition"` // This string has to match a required format. See [OS services monitoring](https://dt-url.net/vl03xzk).\n\n- `$eq(enabled)` – Matches services with startup type equal to enabled.\n\nAvailable logic operations:\n- `$not($eq(enabled))` – Matches services with startup type different from enabled.\n- `$or($eq(enabled),$eq(disabled))` - Matches services that are either enabled or disabled.\n\nUse one of the following values as a parameter for this condition:\n\n- `enabled`\n- `enabled-runtime`\n- `static`\n- `disabled`
}

func (me *LinuxDetectionCondition) Schema() map[string]*hcl.Schema {
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
			Description: "This string has to match a required format. See [OS services monitoring](https://dt-url.net/vl03xzk).\n\n- `$eq(enabled)` – Matches services with startup type equal to enabled.\n\nAvailable logic operations:\n- `$not($eq(enabled))` – Matches services with startup type different from enabled.\n- `$or($eq(enabled),$eq(disabled))` - Matches services that are either enabled or disabled.\n\nUse one of the following values as a parameter for this condition:\n\n- `enabled`\n- `enabled-runtime`\n- `static`\n- `disabled`",
			Required:    true,
		},
	}
}

func (me *LinuxDetectionCondition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"condition":         me.Condition,
		"property":          me.Property,
		"startup_condition": me.StartupCondition,
	})
}

func (me *LinuxDetectionCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"condition":         &me.Condition,
		"property":          &me.Property,
		"startup_condition": &me.StartupCondition,
	})
}
