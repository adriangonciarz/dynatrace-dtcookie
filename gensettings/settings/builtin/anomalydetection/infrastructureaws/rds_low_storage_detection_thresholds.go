package infrastructureaws

import "github.com/dtcookie/hcl"

// RdsLowStorageDetectionThresholds. Alert if the condition is met in 3 out of 5 samples
type RdsLowStorageDetectionThresholds struct {
	FreeStoragePercentage int `json:"freeStoragePercentage"` // Free storage space divided by allocated storage is lower than
}

func (me *RdsLowStorageDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"free_storage_percentage": {
			Type:        hcl.TypeInt,
			Description: "Free storage space divided by allocated storage is lower than",
			Required:    true,
		},
	}
}

func (me *RdsLowStorageDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"free_storage_percentage": me.FreeStoragePercentage,
	})
}

func (me *RdsLowStorageDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"free_storage_percentage": &me.FreeStoragePercentage,
	})
}
