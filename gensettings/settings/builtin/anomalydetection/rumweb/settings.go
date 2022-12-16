package rumweb

import "github.com/dtcookie/hcl"

type Settings struct {
	ErrorRate     *ErrorRate        `json:"errorRate"`     // Error rate
	ResponseTime  *ResponseTime     `json:"responseTime"`  // Response time
	TrafficDrops  *AppTrafficDrops  `json:"trafficDrops"`  // Detect traffic drops
	TrafficSpikes *AppTrafficSpikes `json:"trafficSpikes"` // Detect traffic spikes
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"error_rate": {
			Type:        hcl.TypeList,
			Description: "Error rate",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ErrorRate).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"response_time": {
			Type:        hcl.TypeList,
			Description: "Response time",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ResponseTime).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"traffic_drops": {
			Type:        hcl.TypeList,
			Description: "Detect traffic drops",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(AppTrafficDrops).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"traffic_spikes": {
			Type:        hcl.TypeList,
			Description: "Detect traffic spikes",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(AppTrafficSpikes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"error_rate":     me.ErrorRate,
		"response_time":  me.ResponseTime,
		"traffic_drops":  me.TrafficDrops,
		"traffic_spikes": me.TrafficSpikes,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"error_rate":     &me.ErrorRate,
		"response_time":  &me.ResponseTime,
		"traffic_drops":  &me.TrafficDrops,
		"traffic_spikes": &me.TrafficSpikes,
	})
}
