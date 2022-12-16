package slo

import "github.com/dtcookie/hcl"

type Settings struct {
	CustomDescription   *string              `json:"customDescription"`   // The description of the SLO
	Enabled             bool                 `json:"enabled"`             // Enabled
	ErrorBudgetBurnRate *ErrorBudgetBurnRate `json:"errorBudgetBurnRate"` // ### Error budget burn rate
	EvaluationType      SloEvaluationType    `json:"evaluationType"`      // Select \"Aggregate\" to have the measurements of this success metric aggregated into a single percentage-rate metric.
	EvaluationWindow    string               `json:"evaluationWindow"`    // Define the timeframe during which the SLO is to be evaluated. For the timeframe you can enter expressions like -1h (last hour), -1w (last week) or complex expressions like -2d to now (last two days), -1d/d to now/d (beginning of yesterday to beginning of today).
	Filter              string               `json:"filter"`              // Set a filter parameter (entitySelector) on any GET call to evaluate this SLO against specific services only (for example, type(\"SERVICE\")).  For details, see the [Entity Selector documentation](https://dt-url.net/entityselector).
	MetricExpression    string               `json:"metricExpression"`    // For details, see the [Metrics page](/ui/metrics \"Metrics page\").
	MetricName          string               `json:"metricName"`          // Once created, metric keys cannot be changed
	Name                string               `json:"name"`                // SLO name
	TargetSuccess       float64              `json:"targetSuccess"`       // Set the target value of the SLO. A percentage below this value indicates a failure.
	TargetWarning       float64              `json:"targetWarning"`       // Set the warning value of the SLO. At the warning state the SLO is fulfilled. However, it is getting close to a failure.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_description": {
			Type:        hcl.TypeString,
			Description: "The description of the SLO",
			Optional:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"error_budget_burn_rate": {
			Type:        hcl.TypeList,
			Description: "### Error budget burn rate",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ErrorBudgetBurnRate).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"evaluation_type": {
			Type:        hcl.TypeString,
			Description: "Select \"Aggregate\" to have the measurements of this success metric aggregated into a single percentage-rate metric.",
			Required:    true,
		},
		"evaluation_window": {
			Type:        hcl.TypeString,
			Description: "Define the timeframe during which the SLO is to be evaluated. For the timeframe you can enter expressions like -1h (last hour), -1w (last week) or complex expressions like -2d to now (last two days), -1d/d to now/d (beginning of yesterday to beginning of today).",
			Required:    true,
		},
		"filter": {
			Type:        hcl.TypeString,
			Description: "Set a filter parameter (entitySelector) on any GET call to evaluate this SLO against specific services only (for example, type(\"SERVICE\")).  For details, see the [Entity Selector documentation](https://dt-url.net/entityselector).",
			Required:    true,
		},
		"metric_expression": {
			Type:        hcl.TypeString,
			Description: "For details, see the [Metrics page](/ui/metrics \"Metrics page\").",
			Required:    true,
		},
		"metric_name": {
			Type:        hcl.TypeString,
			Description: "Once created, metric keys cannot be changed",
			Required:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "SLO name",
			Required:    true,
		},
		"target_success": {
			Type:        hcl.TypeFloat,
			Description: "Set the target value of the SLO. A percentage below this value indicates a failure.",
			Required:    true,
		},
		"target_warning": {
			Type:        hcl.TypeFloat,
			Description: "Set the warning value of the SLO. At the warning state the SLO is fulfilled. However, it is getting close to a failure.",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_description":     me.CustomDescription,
		"enabled":                me.Enabled,
		"error_budget_burn_rate": me.ErrorBudgetBurnRate,
		"evaluation_type":        me.EvaluationType,
		"evaluation_window":      me.EvaluationWindow,
		"filter":                 me.Filter,
		"metric_expression":      me.MetricExpression,
		"metric_name":            me.MetricName,
		"name":                   me.Name,
		"target_success":         me.TargetSuccess,
		"target_warning":         me.TargetWarning,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_description":     &me.CustomDescription,
		"enabled":                &me.Enabled,
		"error_budget_burn_rate": &me.ErrorBudgetBurnRate,
		"evaluation_type":        &me.EvaluationType,
		"evaluation_window":      &me.EvaluationWindow,
		"filter":                 &me.Filter,
		"metric_expression":      &me.MetricExpression,
		"metric_name":            &me.MetricName,
		"name":                   &me.Name,
		"target_success":         &me.TargetSuccess,
		"target_warning":         &me.TargetWarning,
	})
}
