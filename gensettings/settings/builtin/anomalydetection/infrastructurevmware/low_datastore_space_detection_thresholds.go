package infrastructurevmware

import "github.com/dtcookie/hcl"

// LowDatastoreSpaceDetectionThresholds. Alert if the condition is met in 1 out of 5 samples
type LowDatastoreSpaceDetectionThresholds struct {
	FreeSpacePercentage int `json:"freeSpacePercentage"` // Datastore free space is lower than
}

func (me *LowDatastoreSpaceDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"free_space_percentage": {
			Type:        hcl.TypeInt,
			Description: "Datastore free space is lower than",
			Required:    true,
		},
	}
}

func (me *LowDatastoreSpaceDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"free_space_percentage": me.FreeSpacePercentage,
	})
}

func (me *LowDatastoreSpaceDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"free_space_percentage": &me.FreeSpacePercentage,
	})
}
