package cloudapplicationworkloaddetection

import "github.com/dtcookie/hcl"

type MatchFilter struct {
	MatchOperator MatchEnum `json:"matchOperator"` // Match operator
	Namespace     string    `json:"namespace"`     // Namespace name
}

func (me *MatchFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"match_operator": {
			Type:        hcl.TypeString,
			Description: "Match operator",
			Required:    true,
		},
		"namespace": {
			Type:        hcl.TypeString,
			Description: "Namespace name",
			Required:    true,
		},
	}
}

func (me *MatchFilter) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"match_operator": me.MatchOperator,
		"namespace":      me.Namespace,
	})
}

func (me *MatchFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"match_operator": &me.MatchOperator,
		"namespace":      &me.Namespace,
	})
}
