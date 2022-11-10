package retransmission

import "github.com/dtcookie/hcl"

// Thresholds Custom thresholds for high retransmission rate. If not set, automatic mode is used.
//  **All** of these conditions must be met to trigger an alert.
type Thresholds struct {
	RetransmissionRatePercentage        int32 `json:"retransmissionRatePercentage"`        // Retransmission rate is higher than *X*% in 3 out of 5 samples.
	RetransmittedPacketsNumberPerMinute int32 `json:"retransmittedPacketsNumberPerMinute"` // Number of retransmitted packets is higher than *X* packets per minute in 3 out of 5 samples.
}

func (me *Thresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"retransmission_rate": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Retransmission rate is higher than *X*% in 3 out of 5 samples",
		},
		"retransmitted_packets": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Number of retransmitted packets is higher than *X* packets per minute in 3 out of 5 samples",
		},
	}
}

func (me *Thresholds) MarshalHCL() (map[string]interface{}, error) {
	return map[string]interface{}{
		"retransmission_rate":   int(me.RetransmissionRatePercentage),
		"retransmitted_packets": int(me.RetransmittedPacketsNumberPerMinute),
	}, nil
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("retransmission_rate"); ok {
		me.RetransmissionRatePercentage = int32(value.(int))
	}
	if value, ok := decoder.GetOk("retransmitted_packets"); ok {
		me.RetransmittedPacketsNumberPerMinute = int32(value.(int))
	}
	return nil
}
