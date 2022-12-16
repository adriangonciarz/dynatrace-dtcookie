package infrastructurehosts

import "github.com/dtcookie/hcl"

type StrictEventThresholds struct {
	DealertingEvaluationWindow int `json:"dealertingEvaluationWindow"` // The number of **10-second samples** that form the sliding evaluation window for dealerting.
	DealertingSamples          int `json:"dealertingSamples"`          // The number of **10-second samples** within the evaluation window that must be lower the threshold to close an event
	ViolatingEvaluationWindow  int `json:"violatingEvaluationWindow"`  // The number of **10-second samples** that form the sliding evaluation window to detect violating samples.
	ViolatingSamples           int `json:"violatingSamples"`           // The number of **10-second samples** within the evaluation window that must exceed the threshold to trigger an event
}

func (me *StrictEventThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dealerting_evaluation_window": {
			Type:        hcl.TypeInt,
			Description: "The number of **10-second samples** that form the sliding evaluation window for dealerting.",
			Required:    true,
		},
		"dealerting_samples": {
			Type:        hcl.TypeInt,
			Description: "The number of **10-second samples** within the evaluation window that must be lower the threshold to close an event",
			Required:    true,
		},
		"violating_evaluation_window": {
			Type:        hcl.TypeInt,
			Description: "The number of **10-second samples** that form the sliding evaluation window to detect violating samples.",
			Required:    true,
		},
		"violating_samples": {
			Type:        hcl.TypeInt,
			Description: "The number of **10-second samples** within the evaluation window that must exceed the threshold to trigger an event",
			Required:    true,
		},
	}
}

func (me *StrictEventThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"dealerting_evaluation_window": me.DealertingEvaluationWindow,
		"dealerting_samples":           me.DealertingSamples,
		"violating_evaluation_window":  me.ViolatingEvaluationWindow,
		"violating_samples":            me.ViolatingSamples,
	})
}

func (me *StrictEventThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"dealerting_evaluation_window": &me.DealertingEvaluationWindow,
		"dealerting_samples":           &me.DealertingSamples,
		"violating_evaluation_window":  &me.ViolatingEvaluationWindow,
		"violating_samples":            &me.ViolatingSamples,
	})
}
