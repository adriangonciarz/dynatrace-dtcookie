package processvisibility

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled      bool `json:"enabled"`      // Enable Process instance snapshots
	MaxProcesses int  `json:"maxProcesses"` // The maximum amount of processes that host may report is **100**
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enable Process instance snapshots",
			Optional:    true,
		},
		"max_processes": {
			Type:        hcl.TypeInt,
			Description: "The maximum amount of processes that host may report is **100**",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":       me.Enabled,
		"max_processes": me.MaxProcesses,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":       &me.Enabled,
		"max_processes": &me.MaxProcesses,
	})
}
