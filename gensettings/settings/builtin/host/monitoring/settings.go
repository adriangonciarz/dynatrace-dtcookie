package monitoring

import "github.com/dtcookie/hcl"

type Settings struct {
	AutoInjection bool `json:"autoInjection"` // An auto-injection disabled with [oneagentctl](https://dt-url.net/oneagentctl) takes precedence over this setting and cannot be changed from the Dynatrace web UI.
	Enabled       bool `json:"enabled"`       // Turn on monitoring to gain visibility into this host, its processes, services, and applications.
	FullStack     bool `json:"fullStack"`     // Dynatrace uses full-stack monitoring by default, to monitor every aspect of your environment, including all processes, services, and applications detected on your hosts. \n\nIf you turn off full-stack monitoring, Dynatrace will only monitor your infrastructure. You will lose access to application performance, user experience data, code-level visibility and PurePath insights. \n\nTo learn more, visit [Infrastructure Monitoring mode](https://www.dynatrace.com/support/help/shortlink/infrastructure).\n\nPlease note that changing the monitoring mode will impact the license consumption of this OneAgent. To learn more, visit [Host units](https://dt-url.net/hi03uns).
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"auto_injection": {
			Type:        hcl.TypeBool,
			Description: "An auto-injection disabled with [oneagentctl](https://dt-url.net/oneagentctl) takes precedence over this setting and cannot be changed from the Dynatrace web UI.",
			Optional:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Turn on monitoring to gain visibility into this host, its processes, services, and applications.",
			Optional:    true,
		},
		"full_stack": {
			Type:        hcl.TypeBool,
			Description: "Dynatrace uses full-stack monitoring by default, to monitor every aspect of your environment, including all processes, services, and applications detected on your hosts. \n\nIf you turn off full-stack monitoring, Dynatrace will only monitor your infrastructure. You will lose access to application performance, user experience data, code-level visibility and PurePath insights. \n\nTo learn more, visit [Infrastructure Monitoring mode](https://www.dynatrace.com/support/help/shortlink/infrastructure).\n\nPlease note that changing the monitoring mode will impact the license consumption of this OneAgent. To learn more, visit [Host units](https://dt-url.net/hi03uns).",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"auto_injection": me.AutoInjection,
		"enabled":        me.Enabled,
		"full_stack":     me.FullStack,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"auto_injection": &me.AutoInjection,
		"enabled":        &me.Enabled,
		"full_stack":     &me.FullStack,
	})
}
