package cloudapplicationworkloaddetection

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled bool            `json:"enabled"` // Enable cloud application and workload detection
	Filters FilterComplexes `json:"filters"` // Define rules to merge similar Kubernetes workloads into process groups. \n\n You can use workload properties like namespace name, base pod name or container name as well as the [environment variables DT_RELEASE_STAGE and DT_RELEASE_PRODUCT](https://dt-url.net/sb02v2a) for grouping processes of similar workloads. The first applicable rule will be applied. If no rule matches, “Namespace name” + “Base pod name” + “Container name” is used as fallback. \n\n Note: This rule set does not cover Cloud Foundry.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enable cloud application and workload detection",
			Optional:    true,
		},
		"filters": {
			Type:        hcl.TypeList,
			Description: "Define rules to merge similar Kubernetes workloads into process groups. \n\n You can use workload properties like namespace name, base pod name or container name as well as the [environment variables DT_RELEASE_STAGE and DT_RELEASE_PRODUCT](https://dt-url.net/sb02v2a) for grouping processes of similar workloads. The first applicable rule will be applied. If no rule matches, “Namespace name” + “Base pod name” + “Container name” is used as fallback. \n\n Note: This rule set does not cover Cloud Foundry.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(FilterComplexes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled": me.Enabled,
		"filters": me.Filters,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled": &me.Enabled,
		"filters": &me.Filters,
	})
}
