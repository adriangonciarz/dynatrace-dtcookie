package space

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// DetectionConfig Configuration of low disk space detection.
type DetectionConfig struct {
	Enabled          bool        `json:"enabled"`                    // The detection is enabled (`true`) or disabled (`false`).
	CustomThresholds *Thresholds `json:"customThresholds,omitempty"` // Custom thresholds for low disk space. If not set, automatic mode is used.
}

func (me *DetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Required:    true,
			Description: "The detection is enabled (`true`) or disabled (`false`)",
		},
		"thresholds": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Custom thresholds for low disk space. If not set, automatic mode is used",
			Elem:        &hcl.Resource{Schema: new(Thresholds).Schema()},
		},
	}
}

func (me *DetectionConfig) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["enabled"] = me.Enabled
	if me.CustomThresholds != nil {
		if marshalled, err := me.CustomThresholds.MarshalHCL(); err == nil {
			result["thresholds"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *DetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	adapter := hcl.Adapt(decoder)
	me.Enabled = opt.Bool(adapter.GetBool("enabled"))
	if _, ok := decoder.GetOk("thresholds.#"); ok {
		me.CustomThresholds = new(Thresholds)
		if err := me.CustomThresholds.UnmarshalHCL(hcl.NewDecoder(decoder, "thresholds", 0)); err != nil {
			return err
		}
	}
	return nil
}
