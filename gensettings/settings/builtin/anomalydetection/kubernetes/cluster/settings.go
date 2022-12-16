package cluster

import "github.com/dtcookie/hcl"

type Settings struct {
	CpuRequestsSaturation    *CpuRequestsSaturation    `json:"cpuRequestsSaturation"`
	MemoryRequestsSaturation *MemoryRequestsSaturation `json:"memoryRequestsSaturation"`
	ReadinessIssues          *ReadinessIssues          `json:"readinessIssues"` // Alerts if cluster has not been ready for a given amount of time
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cpu_requests_saturation": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CpuRequestsSaturation).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"memory_requests_saturation": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MemoryRequestsSaturation).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"readiness_issues": {
			Type:        hcl.TypeList,
			Description: "Alerts if cluster has not been ready for a given amount of time",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ReadinessIssues).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"cpu_requests_saturation":    me.CpuRequestsSaturation,
		"memory_requests_saturation": me.MemoryRequestsSaturation,
		"readiness_issues":           me.ReadinessIssues,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cpu_requests_saturation":    &me.CpuRequestsSaturation,
		"memory_requests_saturation": &me.MemoryRequestsSaturation,
		"readiness_issues":           &me.ReadinessIssues,
	})
}
