package traffic

import "github.com/dtcookie/hcl"

type Settings struct {
	NetTracer bool `json:"netTracer"` // When disabled, OneAgent won't use NetTracer to collect network data from containers. Disabled by default. Applies only to Linux hosts. Requires OneAgent 1.231+.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"net_tracer": {
			Type:        hcl.TypeBool,
			Description: "When disabled, OneAgent won't use NetTracer to collect network data from containers. Disabled by default. Applies only to Linux hosts. Requires OneAgent 1.231+.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"net_tracer": me.NetTracer,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"net_tracer": &me.NetTracer,
	})
}
