package processgroup

import "github.com/dtcookie/hcl"

type Settings struct {
	Enable bool `json:"enable"` // Allows OneAgent to:\n* automatically inject the RUM JavaScript tag into each page delivered by this process group\n* provide the necessary info to correlate RUM data with server-side PurePaths\n* forward beacons to the cluster\n* deliver the monitoring code
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enable": {
			Type:        hcl.TypeBool,
			Description: "Allows OneAgent to:\n* automatically inject the RUM JavaScript tag into each page delivered by this process group\n* provide the necessary info to correlate RUM data with server-side PurePaths\n* forward beacons to the cluster\n* deliver the monitoring code",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enable": me.Enable,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enable": &me.Enable,
	})
}
