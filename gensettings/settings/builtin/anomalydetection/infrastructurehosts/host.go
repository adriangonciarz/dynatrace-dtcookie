package infrastructurehosts

import "github.com/dtcookie/hcl"

type Host struct {
	ConnectionLostDetection    *ConnectionLostDetection    `json:"connectionLostDetection"`
	HighCpuSaturationDetection *HighCpuSaturationDetection `json:"highCpuSaturationDetection"`
	HighGcActivityDetection    *HighGcActivityDetection    `json:"highGcActivityDetection"`
	HighMemoryDetection        *HighMemoryDetection        `json:"highMemoryDetection"`
	OutOfMemoryDetection       *OutOfMemoryDetection       `json:"outOfMemoryDetection"`
	OutOfThreadsDetection      *OutOfThreadsDetection      `json:"outOfThreadsDetection"`
}

func (me *Host) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"connection_lost_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ConnectionLostDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"high_cpu_saturation_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(HighCpuSaturationDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"high_gc_activity_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(HighGcActivityDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"high_memory_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(HighMemoryDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"out_of_memory_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(OutOfMemoryDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"out_of_threads_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(OutOfThreadsDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Host) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"connection_lost_detection":     me.ConnectionLostDetection,
		"high_cpu_saturation_detection": me.HighCpuSaturationDetection,
		"high_gc_activity_detection":    me.HighGcActivityDetection,
		"high_memory_detection":         me.HighMemoryDetection,
		"out_of_memory_detection":       me.OutOfMemoryDetection,
		"out_of_threads_detection":      me.OutOfThreadsDetection,
	})
}

func (me *Host) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"connection_lost_detection":     &me.ConnectionLostDetection,
		"high_cpu_saturation_detection": &me.HighCpuSaturationDetection,
		"high_gc_activity_detection":    &me.HighGcActivityDetection,
		"high_memory_detection":         &me.HighMemoryDetection,
		"out_of_memory_detection":       &me.OutOfMemoryDetection,
		"out_of_threads_detection":      &me.OutOfThreadsDetection,
	})
}
