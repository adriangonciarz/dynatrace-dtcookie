package rumweb

import "github.com/dtcookie/hcl"

type TrafficDrops struct {
	AbnormalStateAbnormalState float64 `json:"abnormalStateAbnormalState"` // Minutes the observed traffic has to stay in abnormal state before alert
	TrafficDropPercentage      float64 `json:"trafficDropPercentage"`      // Alert if the observed traffic is less than this percentage of the expected value
}

func (me *TrafficDrops) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"abnormal_state_abnormal_state": {
			Type:        hcl.TypeFloat,
			Description: "Minutes the observed traffic has to stay in abnormal state before alert",
			Required:    true,
		},
		"traffic_drop_percentage": {
			Type:        hcl.TypeFloat,
			Description: "Alert if the observed traffic is less than this percentage of the expected value",
			Required:    true,
		},
	}
}

func (me *TrafficDrops) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"abnormal_state_abnormal_state": me.AbnormalStateAbnormalState,
		"traffic_drop_percentage":       me.TrafficDropPercentage,
	})
}

func (me *TrafficDrops) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"abnormal_state_abnormal_state": &me.AbnormalStateAbnormalState,
		"traffic_drop_percentage":       &me.TrafficDropPercentage,
	})
}
