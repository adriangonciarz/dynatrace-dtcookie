package infrastructurehosts

import "github.com/dtcookie/hcl"

type HighMemoryDetectionThresholds struct {
	EventThresholds                *EventThresholds `json:"eventThresholds"`
	PageFaultsPerSecondNonWindows  int              `json:"pageFaultsPerSecondNonWindows"`  // Alert if the memory page fault rate on Unix systems is higher than this threshold for the defined amount of samples
	PageFaultsPerSecondWindows     int              `json:"pageFaultsPerSecondWindows"`     // Alert if the memory page fault rate on Windows is higher than this threshold for the defined amount of samples
	UsedMemoryPercentageNonWindows int              `json:"usedMemoryPercentageNonWindows"` // Alert if the memory usage on Unix systems is higher than this threshold
	UsedMemoryPercentageWindows    int              `json:"usedMemoryPercentageWindows"`    // Alert if the memory usage on Windows is higher than this threshold
}

func (me *HighMemoryDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"event_thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EventThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"page_faults_per_second_non_windows": {
			Type:        hcl.TypeInt,
			Description: "Alert if the memory page fault rate on Unix systems is higher than this threshold for the defined amount of samples",
			Required:    true,
		},
		"page_faults_per_second_windows": {
			Type:        hcl.TypeInt,
			Description: "Alert if the memory page fault rate on Windows is higher than this threshold for the defined amount of samples",
			Required:    true,
		},
		"used_memory_percentage_non_windows": {
			Type:        hcl.TypeInt,
			Description: "Alert if the memory usage on Unix systems is higher than this threshold",
			Required:    true,
		},
		"used_memory_percentage_windows": {
			Type:        hcl.TypeInt,
			Description: "Alert if the memory usage on Windows is higher than this threshold",
			Required:    true,
		},
	}
}

func (me *HighMemoryDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"event_thresholds":                   me.EventThresholds,
		"page_faults_per_second_non_windows": me.PageFaultsPerSecondNonWindows,
		"page_faults_per_second_windows":     me.PageFaultsPerSecondWindows,
		"used_memory_percentage_non_windows": me.UsedMemoryPercentageNonWindows,
		"used_memory_percentage_windows":     me.UsedMemoryPercentageWindows,
	})
}

func (me *HighMemoryDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"event_thresholds":                   &me.EventThresholds,
		"page_faults_per_second_non_windows": &me.PageFaultsPerSecondNonWindows,
		"page_faults_per_second_windows":     &me.PageFaultsPerSecondWindows,
		"used_memory_percentage_non_windows": &me.UsedMemoryPercentageNonWindows,
		"used_memory_percentage_windows":     &me.UsedMemoryPercentageWindows,
	})
}
