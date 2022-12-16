package services

import "github.com/dtcookie/hcl"

type Settings struct {
	FailureRate  *FailureRate  `json:"failureRate"`  // Failure rate
	LoadDrops    *LoadDrops    `json:"loadDrops"`    // Alert if the observed load is lower than the expected load by a specified margin for a specified amount of time:
	LoadSpikes   *LoadSpikes   `json:"loadSpikes"`   // Alert if the observed load exceeds the expected load by a specified margin for a specified amount of time:
	ResponseTime *ResponseTime `json:"responseTime"` // Response time
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"failure_rate": {
			Type:        hcl.TypeList,
			Description: "Failure rate",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(FailureRate).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"load_drops": {
			Type:        hcl.TypeList,
			Description: "Alert if the observed load is lower than the expected load by a specified margin for a specified amount of time:",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(LoadDrops).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"load_spikes": {
			Type:        hcl.TypeList,
			Description: "Alert if the observed load exceeds the expected load by a specified margin for a specified amount of time:",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(LoadSpikes).Schema()},
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
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"failure_rate":  me.FailureRate,
		"load_drops":    me.LoadDrops,
		"load_spikes":   me.LoadSpikes,
		"response_time": me.ResponseTime,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"failure_rate":  &me.FailureRate,
		"load_drops":    &me.LoadDrops,
		"load_spikes":   &me.LoadSpikes,
		"response_time": &me.ResponseTime,
	})
}
