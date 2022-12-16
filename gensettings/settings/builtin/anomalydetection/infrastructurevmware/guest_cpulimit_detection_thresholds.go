package infrastructurevmware

import "github.com/dtcookie/hcl"

// GuestCPULimitDetectionThresholds. Alert if **all three** conditions are met in 3 out of 5 samples
type GuestCPULimitDetectionThresholds struct {
	HostCpuUsagePercentage int `json:"hostCpuUsagePercentage"` // Hypervisor CPU usage is higher than
	VmCpuReadyPercentage   int `json:"vmCpuReadyPercentage"`   // VM CPU ready is higher than
	VmCpuUsagePercentage   int `json:"vmCpuUsagePercentage"`   // VM CPU usage (VM CPU Usage Mhz / VM CPU limit in Mhz) is higher than
}

func (me *GuestCPULimitDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"host_cpu_usage_percentage": {
			Type:        hcl.TypeInt,
			Description: "Hypervisor CPU usage is higher than",
			Required:    true,
		},
		"vm_cpu_ready_percentage": {
			Type:        hcl.TypeInt,
			Description: "VM CPU ready is higher than",
			Required:    true,
		},
		"vm_cpu_usage_percentage": {
			Type:        hcl.TypeInt,
			Description: "VM CPU usage (VM CPU Usage Mhz / VM CPU limit in Mhz) is higher than",
			Required:    true,
		},
	}
}

func (me *GuestCPULimitDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"host_cpu_usage_percentage": me.HostCpuUsagePercentage,
		"vm_cpu_ready_percentage":   me.VmCpuReadyPercentage,
		"vm_cpu_usage_percentage":   me.VmCpuUsagePercentage,
	})
}

func (me *GuestCPULimitDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"host_cpu_usage_percentage": &me.HostCpuUsagePercentage,
		"vm_cpu_ready_percentage":   &me.VmCpuReadyPercentage,
		"vm_cpu_usage_percentage":   &me.VmCpuUsagePercentage,
	})
}
