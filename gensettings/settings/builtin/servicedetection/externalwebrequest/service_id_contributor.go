package externalwebrequest

import "github.com/dtcookie/hcl"

type ServiceIdContributor struct {
	EnableIdContributor  bool               `json:"enableIdContributor"`  // Enable service identifier contributor
	ServiceIdContributor *TransformationSet `json:"serviceIdContributor"` // Service identifier contributor
}

func (me *ServiceIdContributor) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enable_id_contributor": {
			Type:        hcl.TypeBool,
			Description: "Enable service identifier contributor",
			Optional:    true,
		},
		"service_id_contributor": {
			Type:        hcl.TypeList,
			Description: "Service identifier contributor",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(TransformationSet).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *ServiceIdContributor) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enable_id_contributor":  me.EnableIdContributor,
		"service_id_contributor": me.ServiceIdContributor,
	})
}

func (me *ServiceIdContributor) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enable_id_contributor":  &me.EnableIdContributor,
		"service_id_contributor": &me.ServiceIdContributor,
	})
}
