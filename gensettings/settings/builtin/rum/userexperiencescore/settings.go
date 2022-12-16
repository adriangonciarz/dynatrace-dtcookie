package userexperiencescore

import "github.com/dtcookie/hcl"

type Settings struct {
	ConsiderLastAction                bool `json:"considerLastAction"`                // If last user action in a session is classified as Frustrating, classify the entire session as Frustrating
	ConsiderRageClick                 bool `json:"considerRageClick"`                 // Consider rage clicks / rage taps in score calculation
	MaxFrustratedUserActionsThreshold int  `json:"maxFrustratedUserActionsThreshold"` // User experience is considered Frustrating when the selected percentage or more of the user actions in a session are rated as Frustrating.
	MinSatisfiedUserActionsThreshold  int  `json:"minSatisfiedUserActionsThreshold"`  // User experience is considered Satisfying when at least the selected percentage of the user actions in a session are rated as Satisfying.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"consider_last_action": {
			Type:        hcl.TypeBool,
			Description: "If last user action in a session is classified as Frustrating, classify the entire session as Frustrating",
			Optional:    true,
		},
		"consider_rage_click": {
			Type:        hcl.TypeBool,
			Description: "Consider rage clicks / rage taps in score calculation",
			Optional:    true,
		},
		"max_frustrated_user_actions_threshold": {
			Type:        hcl.TypeInt,
			Description: "User experience is considered Frustrating when the selected percentage or more of the user actions in a session are rated as Frustrating.",
			Required:    true,
		},
		"min_satisfied_user_actions_threshold": {
			Type:        hcl.TypeInt,
			Description: "User experience is considered Satisfying when at least the selected percentage of the user actions in a session are rated as Satisfying.",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"consider_last_action":                  me.ConsiderLastAction,
		"consider_rage_click":                   me.ConsiderRageClick,
		"max_frustrated_user_actions_threshold": me.MaxFrustratedUserActionsThreshold,
		"min_satisfied_user_actions_threshold":  me.MinSatisfiedUserActionsThreshold,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"consider_last_action":                  &me.ConsiderLastAction,
		"consider_rage_click":                   &me.ConsiderRageClick,
		"max_frustrated_user_actions_threshold": &me.MaxFrustratedUserActionsThreshold,
		"min_satisfied_user_actions_threshold":  &me.MinSatisfiedUserActionsThreshold,
	})
}
