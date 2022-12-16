package namespace

import "github.com/dtcookie/hcl"

type Settings struct {
	CpuLimitsQuotaSaturation      *CpuLimitsQuotaSaturation      `json:"cpuLimitsQuotaSaturation"`      // Alerts if almost no CPU limits quota left in namespace
	CpuRequestsQuotaSaturation    *CpuRequestsQuotaSaturation    `json:"cpuRequestsQuotaSaturation"`    // Alerts if almost no CPU requests quota left in namespace
	MemoryLimitsQuotaSaturation   *MemoryLimitsQuotaSaturation   `json:"memoryLimitsQuotaSaturation"`   // Alerts if almost no memory limits quota left in namespace
	MemoryRequestsQuotaSaturation *MemoryRequestsQuotaSaturation `json:"memoryRequestsQuotaSaturation"` // Alerts if almost no memory requests quota left in namespace
	PodsQuotaSaturation           *PodsQuotaSaturation           `json:"podsQuotaSaturation"`           // Alerts if almost no pod number limits quota left in namespace
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cpu_limits_quota_saturation": {
			Type:        hcl.TypeList,
			Description: "Alerts if almost no CPU limits quota left in namespace",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CpuLimitsQuotaSaturation).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"cpu_requests_quota_saturation": {
			Type:        hcl.TypeList,
			Description: "Alerts if almost no CPU requests quota left in namespace",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CpuRequestsQuotaSaturation).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"memory_limits_quota_saturation": {
			Type:        hcl.TypeList,
			Description: "Alerts if almost no memory limits quota left in namespace",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MemoryLimitsQuotaSaturation).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"memory_requests_quota_saturation": {
			Type:        hcl.TypeList,
			Description: "Alerts if almost no memory requests quota left in namespace",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MemoryRequestsQuotaSaturation).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"pods_quota_saturation": {
			Type:        hcl.TypeList,
			Description: "Alerts if almost no pod number limits quota left in namespace",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(PodsQuotaSaturation).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"cpu_limits_quota_saturation":      me.CpuLimitsQuotaSaturation,
		"cpu_requests_quota_saturation":    me.CpuRequestsQuotaSaturation,
		"memory_limits_quota_saturation":   me.MemoryLimitsQuotaSaturation,
		"memory_requests_quota_saturation": me.MemoryRequestsQuotaSaturation,
		"pods_quota_saturation":            me.PodsQuotaSaturation,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cpu_limits_quota_saturation":      &me.CpuLimitsQuotaSaturation,
		"cpu_requests_quota_saturation":    &me.CpuRequestsQuotaSaturation,
		"memory_limits_quota_saturation":   &me.MemoryLimitsQuotaSaturation,
		"memory_requests_quota_saturation": &me.MemoryRequestsQuotaSaturation,
		"pods_quota_saturation":            &me.PodsQuotaSaturation,
	})
}
