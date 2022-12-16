package infrastructurevmware

import "github.com/dtcookie/hcl"

// EsxiHighCpuDetectionThresholds. Alert if **all three** conditions are met in 3 out of 5 samples
type EsxiHighCpuDetectionThresholds struct {
	CpuPeakPercentage    int `json:"cpuPeakPercentage"`    // At least one peak occurred when Hypervisor CPU usage was higher than
	CpuUsagePercentage   int `json:"cpuUsagePercentage"`   // CPU usage is higher than
	VmCpuReadyPercentage int `json:"vmCpuReadyPercentage"` // VM CPU ready is higher than
}

func (me *EsxiHighCpuDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cpu_peak_percentage": {
			Type:        hcl.TypeInt,
			Description: "At least one peak occurred when Hypervisor CPU usage was higher than",
			Required:    true,
		},
		"cpu_usage_percentage": {
			Type:        hcl.TypeInt,
			Description: "CPU usage is higher than",
			Required:    true,
		},
		"vm_cpu_ready_percentage": {
			Type:        hcl.TypeInt,
			Description: "VM CPU ready is higher than",
			Required:    true,
		},
	}
}

func (me *EsxiHighCpuDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"cpu_peak_percentage":     me.CpuPeakPercentage,
		"cpu_usage_percentage":    me.CpuUsagePercentage,
		"vm_cpu_ready_percentage": me.VmCpuReadyPercentage,
	})
}

func (me *EsxiHighCpuDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cpu_peak_percentage":     &me.CpuPeakPercentage,
		"cpu_usage_percentage":    &me.CpuUsagePercentage,
		"vm_cpu_ready_percentage": &me.VmCpuReadyPercentage,
	})
}
