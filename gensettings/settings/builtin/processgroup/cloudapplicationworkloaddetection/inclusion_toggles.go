package cloudapplicationworkloaddetection

import "github.com/dtcookie/hcl"

type InclusionToggles struct {
	IncBasepod   bool `json:"incBasepod"`   // E.g. \"cloud-credential-operator-\" for \"cloud-credential-operator-5ff6dbff57-gszgq\"
	IncContainer bool `json:"incContainer"` // Container name
	IncNamespace bool `json:"incNamespace"` // Namespace name
	IncProduct   bool `json:"incProduct"`   // If Product is enabled and has no value, it defaults to Base pod name
	IncStage     bool `json:"incStage"`     // Stage
}

func (me *InclusionToggles) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"inc_basepod": {
			Type:        hcl.TypeBool,
			Description: "E.g. \"cloud-credential-operator-\" for \"cloud-credential-operator-5ff6dbff57-gszgq\"",
			Optional:    true,
		},
		"inc_container": {
			Type:        hcl.TypeBool,
			Description: "Container name",
			Optional:    true,
		},
		"inc_namespace": {
			Type:        hcl.TypeBool,
			Description: "Namespace name",
			Optional:    true,
		},
		"inc_product": {
			Type:        hcl.TypeBool,
			Description: "If Product is enabled and has no value, it defaults to Base pod name",
			Optional:    true,
		},
		"inc_stage": {
			Type:        hcl.TypeBool,
			Description: "Stage",
			Optional:    true,
		},
	}
}

func (me *InclusionToggles) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"inc_basepod":   me.IncBasepod,
		"inc_container": me.IncContainer,
		"inc_namespace": me.IncNamespace,
		"inc_product":   me.IncProduct,
		"inc_stage":     me.IncStage,
	})
}

func (me *InclusionToggles) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"inc_basepod":   &me.IncBasepod,
		"inc_container": &me.IncContainer,
		"inc_namespace": &me.IncNamespace,
		"inc_product":   &me.IncProduct,
		"inc_stage":     &me.IncStage,
	})
}
