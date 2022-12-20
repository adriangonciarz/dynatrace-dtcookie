package monitors

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
)

// OutageHandlingPolicy Outage handling configuration
type OutageHandlingPolicy struct {
	GlobalOutage       bool                `json:"globalOutage"`       // When enabled (`true`), generate a problem and send an alert when the monitor is unavailable at all configured locations
	GlobalOutagePolicy *GlobalOutagePolicy `json:"globalOutagePolicy"` // Global outage handling configuration. \n\n Alert if **consecutiveRuns** times consecutively
	LocalOutage        bool                `json:"localOutage"`        // When enabled (`true`), generate a problem and send an alert when the monitor is unavailable for one or more consecutive runs at any location
	LocalOutagePolicy  *LocalOutagePolicy  `json:"localOutagePolicy"`  // Local outage handling configuration. \n\n Alert if **affectedLocations** of locations are unable to access the web application **consecutiveRuns** times consecutively
	RetryOnError       bool                `json:"retryOnError"`       // Schedule retry if browser monitor execution results in a fail. For HTTP monitors this property is ignored
}

func (me *OutageHandlingPolicy) MarshalJSON() ([]byte, error) {
	lop := me.LocalOutagePolicy
	if lop == nil {
		lop = new(LocalOutagePolicy)
	}
	gop := me.GlobalOutagePolicy
	if gop == nil {
		gop = new(GlobalOutagePolicy)
	}
	return json.Marshal(struct {
		GlobalOutage       bool                `json:"globalOutage"`
		GlobalOutagePolicy *GlobalOutagePolicy `json:"globalOutagePolicy"`
		LocalOutage        bool                `json:"localOutage"`
		LocalOutagePolicy  *LocalOutagePolicy  `json:"localOutagePolicy"`
		RetryOnError       bool                `json:"retryOnError"`
	}{
		me.GlobalOutage,
		gop,
		me.LocalOutage,
		lop,
		me.RetryOnError,
	})
}

func (me *OutageHandlingPolicy) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"global_outage": {
			Type:        hcl.TypeBool,
			Description: "When enabled (`true`), generate a problem and send an alert when the monitor is unavailable at all configured locations",
			Optional:    true,
		},
		"local_outage": {
			Type:        hcl.TypeBool,
			Description: "When enabled (`true`), generate a problem and send an alert when the monitor is unavailable for one or more consecutive runs at any location",
			Optional:    true,
		},
		"retry_on_error": {
			Type:        hcl.TypeBool,
			Description: "Schedule retry if browser monitor execution results in a fail. For HTTP monitors this property is ignored",
			Optional:    true,
		},
		"local_outage_policy": {
			Type:        hcl.TypeList,
			Description: "Local outage handling configuration. \n\n Alert if **affectedLocations** of locations are unable to access the web application **consecutiveRuns** times consecutively",
			Optional:    true,
			Elem:        &hcl.Resource{Schema: new(LocalOutagePolicy).Schema()},
		},
		"global_outage_policy": {
			Type:        hcl.TypeList,
			Description: "Global outage handling configuration. \n\n Alert if **consecutiveRuns** times consecutively",
			Optional:    true,
			Elem:        &hcl.Resource{Schema: new(GlobalOutagePolicy).Schema()},
		},
	}
}

func (me *OutageHandlingPolicy) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["global_outage"] = me.GlobalOutage
	result["local_outage"] = me.LocalOutage
	result["retry_on_error"] = me.RetryOnError
	if me.LocalOutagePolicy != nil && (me.LocalOutagePolicy.AffectedLocations != nil || me.LocalOutagePolicy.ConsecutiveRuns != nil) {
		if marshalled, err := me.LocalOutagePolicy.MarshalHCL(); err == nil {
			result["local_outage_policy"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.GlobalOutagePolicy != nil && me.GlobalOutagePolicy.ConsecutiveRuns != nil {
		if marshalled, err := me.GlobalOutagePolicy.MarshalHCL(); err == nil {
			result["global_outage_policy"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *OutageHandlingPolicy) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("global_outage", &me.GlobalOutage); err != nil {
		return err
	}
	if err := decoder.Decode("local_outage", &me.LocalOutage); err != nil {
		return err
	}
	if err := decoder.Decode("retry_on_error", &me.RetryOnError); err != nil {
		return err
	}
	if err := decoder.Decode("local_outage_policy", &me.LocalOutagePolicy); err != nil {
		return err
	}
	if err := decoder.Decode("global_outage_policy", &me.GlobalOutagePolicy); err != nil {
		return err
	}
	return nil
}
