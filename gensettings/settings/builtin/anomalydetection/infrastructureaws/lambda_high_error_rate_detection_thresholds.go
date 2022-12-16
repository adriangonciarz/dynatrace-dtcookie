package infrastructureaws

import "github.com/dtcookie/hcl"

// LambdaHighErrorRateDetectionThresholds. Alert if the condition is met in 3 out of 5 samples
type LambdaHighErrorRateDetectionThresholds struct {
	FailedInvocationsRate int `json:"failedInvocationsRate"` // Failed invocations rate is higher than
}

func (me *LambdaHighErrorRateDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"failed_invocations_rate": {
			Type:        hcl.TypeInt,
			Description: "Failed invocations rate is higher than",
			Required:    true,
		},
	}
}

func (me *LambdaHighErrorRateDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"failed_invocations_rate": me.FailedInvocationsRate,
	})
}

func (me *LambdaHighErrorRateDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"failed_invocations_rate": &me.FailedInvocationsRate,
	})
}
