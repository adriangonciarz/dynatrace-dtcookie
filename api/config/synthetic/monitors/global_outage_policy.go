package monitors

import "github.com/dtcookie/hcl"

// GlobalOutagePolicy Global outage handling configuration. \n\n Alert if **consecutiveRuns** times consecutively
type GlobalOutagePolicy struct {
	ConsecutiveRuns *int32 `json:"consecutiveRuns"` // The number of consecutive fails to trigger an alert
}

func (me *GlobalOutagePolicy) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"consecutive_runs": {
			Type:        hcl.TypeInt,
			Description: "The number of consecutive fails to trigger an alert",
			Required:    true,
		},
	}
}

func (me *GlobalOutagePolicy) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["consecutive_runs"] = *me.ConsecutiveRuns
	return result, nil
}

func (me *GlobalOutagePolicy) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("consecutive_runs", &me.ConsecutiveRuns); err != nil {
		return err
	}
	return nil
}
