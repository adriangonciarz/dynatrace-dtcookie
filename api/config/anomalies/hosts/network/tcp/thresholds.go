package tcp

import "github.com/dtcookie/hcl"

// NetworkTcpProblemsThresholds Custom thresholds for TCP connection problems. If not set, automatic mode is used.
//  **All** of these conditions must be met to trigger an alert.
type Thresholds struct {
	NewConnectionFailuresPercentage  int32 `json:"newConnectionFailuresPercentage"`  // Percentage of new connection failures is higher than *X*% in 3 out of 5 samples.
	FailedConnectionsNumberPerMinute int32 `json:"failedConnectionsNumberPerMinute"` // Number of failed connections is higher than *X* connections per minute in 3 out of 5 samples.
}

func (me *Thresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"new_connection_failures": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Percentage of new connection failures is higher than *X*% in 3 out of 5 samples",
		},
		"failed_connections": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Number of failed connections is higher than *X* connections per minute in 3 out of 5 samples",
		},
	}
}

func (me *Thresholds) MarshalHCL() (map[string]interface{}, error) {
	return map[string]interface{}{
		"new_connection_failures": int(me.NewConnectionFailuresPercentage),
		"failed_connections":      int(me.FailedConnectionsNumberPerMinute),
	}, nil
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("new_connection_failures"); ok {
		me.NewConnectionFailuresPercentage = int32(value.(int))
	}
	if value, ok := decoder.GetOk("failed_connections"); ok {
		me.FailedConnectionsNumberPerMinute = int32(value.(int))
	}
	return nil
}
