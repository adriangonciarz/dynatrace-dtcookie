package slo

import "github.com/dtcookie/hcl"

type ErrorBudgetBurnRate struct {
	BurnRateVisualizationEnabled bool    `json:"burnRateVisualizationEnabled"` // Burn rate visualization enabled
	FastBurnThreshold            float64 `json:"fastBurnThreshold"`            // The threshold defines when a burn rate is marked as fast-burning (high-emergency). Burn rates lower than this threshold (and greater than 1) are highlighted as slow-burn (low-emergency).
}

func (me *ErrorBudgetBurnRate) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"burn_rate_visualization_enabled": {
			Type:        hcl.TypeBool,
			Description: "Burn rate visualization enabled",
			Optional:    true,
		},
		"fast_burn_threshold": {
			Type:        hcl.TypeFloat,
			Description: "The threshold defines when a burn rate is marked as fast-burning (high-emergency). Burn rates lower than this threshold (and greater than 1) are highlighted as slow-burn (low-emergency).",
			Required:    true,
		},
	}
}

func (me *ErrorBudgetBurnRate) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"burn_rate_visualization_enabled": me.BurnRateVisualizationEnabled,
		"fast_burn_threshold":             me.FastBurnThreshold,
	})
}

func (me *ErrorBudgetBurnRate) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"burn_rate_visualization_enabled": &me.BurnRateVisualizationEnabled,
		"fast_burn_threshold":             &me.FastBurnThreshold,
	})
}
