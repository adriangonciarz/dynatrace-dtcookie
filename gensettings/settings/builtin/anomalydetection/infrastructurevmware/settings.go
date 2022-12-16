package infrastructurevmware

import "github.com/dtcookie/hcl"

type Settings struct {
	DroppedPacketsDetection      *DroppedPacketsDetectionConfig      `json:"droppedPacketsDetection"`
	EsxiHighCpuDetection         *EsxiHighCpuDetectionConfig         `json:"esxiHighCpuDetection"`
	EsxiHighMemoryDetection      *EsxiHighMemoryDetectionConfig      `json:"esxiHighMemoryDetection"`
	GuestCpuLimitDetection       *GuestCPULimitDetectionConfig       `json:"guestCpuLimitDetection"`
	LowDatastoreSpaceDetection   *LowDatastoreSpaceDetectionConfig   `json:"lowDatastoreSpaceDetection"`
	OverloadedStorageDetection   *OverloadedStorageDetectionConfig   `json:"overloadedStorageDetection"`
	SlowPhysicalStorageDetection *SlowPhysicalStorageDetectionConfig `json:"slowPhysicalStorageDetection"`
	UndersizedStorageDetection   *UndersizedStorageDetectionConfig   `json:"undersizedStorageDetection"`
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dropped_packets_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DroppedPacketsDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"esxi_high_cpu_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EsxiHighCpuDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"esxi_high_memory_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EsxiHighMemoryDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"guest_cpu_limit_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(GuestCPULimitDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"low_datastore_space_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(LowDatastoreSpaceDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"overloaded_storage_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(OverloadedStorageDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"slow_physical_storage_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SlowPhysicalStorageDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"undersized_storage_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(UndersizedStorageDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"dropped_packets_detection":       me.DroppedPacketsDetection,
		"esxi_high_cpu_detection":         me.EsxiHighCpuDetection,
		"esxi_high_memory_detection":      me.EsxiHighMemoryDetection,
		"guest_cpu_limit_detection":       me.GuestCpuLimitDetection,
		"low_datastore_space_detection":   me.LowDatastoreSpaceDetection,
		"overloaded_storage_detection":    me.OverloadedStorageDetection,
		"slow_physical_storage_detection": me.SlowPhysicalStorageDetection,
		"undersized_storage_detection":    me.UndersizedStorageDetection,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"dropped_packets_detection":       &me.DroppedPacketsDetection,
		"esxi_high_cpu_detection":         &me.EsxiHighCpuDetection,
		"esxi_high_memory_detection":      &me.EsxiHighMemoryDetection,
		"guest_cpu_limit_detection":       &me.GuestCpuLimitDetection,
		"low_datastore_space_detection":   &me.LowDatastoreSpaceDetection,
		"overloaded_storage_detection":    &me.OverloadedStorageDetection,
		"slow_physical_storage_detection": &me.SlowPhysicalStorageDetection,
		"undersized_storage_detection":    &me.UndersizedStorageDetection,
	})
}
