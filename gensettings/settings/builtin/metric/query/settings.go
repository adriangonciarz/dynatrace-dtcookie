package query

import "github.com/dtcookie/hcl"

type Settings struct {
	MetricSelector string `json:"metricSelector"` // Query
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"metric_selector": {
			Type:        hcl.TypeString,
			Description: "Query",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"metric_selector": me.MetricSelector,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"metric_selector": &me.MetricSelector,
	})
}
