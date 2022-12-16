package diskrules

import "github.com/dtcookie/hcl"

type Settings struct {
	DiskNameFilter        *DiskNameFilter `json:"diskNameFilter"`        // Only apply to disks whose name matches
	Enabled               bool            `json:"enabled"`               // Enabled
	Metric                DiskMetric      `json:"metric"`                // Metric to alert on
	Name                  string          `json:"name"`                  // Name
	SampleLimit           *SampleLimit    `json:"sampleLimit"`           // Only alert if the threshold was violated in at least *n* of the last *m* samples
	TagFilters            []string        `json:"tagFilters"`            // Only apply to hosts that have the following tags
	ThresholdMilliseconds float64         `json:"thresholdMilliseconds"` // Alert if higher than
	ThresholdPercent      float64         `json:"thresholdPercent"`      // Alert if lower than
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"disk_name_filter": {
			Type:        hcl.TypeList,
			Description: "Only apply to disks whose name matches",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DiskNameFilter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"metric": {
			Type:        hcl.TypeString,
			Description: "Metric to alert on",
			Required:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "Name",
			Required:    true,
		},
		"sample_limit": {
			Type:        hcl.TypeList,
			Description: "Only alert if the threshold was violated in at least *n* of the last *m* samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SampleLimit).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"tag_filters": {
			Type:        hcl.TypeSet,
			Description: "Only apply to hosts that have the following tags",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"threshold_milliseconds": {
			Type:        hcl.TypeFloat,
			Description: "Alert if higher than",
			Required:    true,
		},
		"threshold_percent": {
			Type:        hcl.TypeFloat,
			Description: "Alert if lower than",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"disk_name_filter":       me.DiskNameFilter,
		"enabled":                me.Enabled,
		"metric":                 me.Metric,
		"name":                   me.Name,
		"sample_limit":           me.SampleLimit,
		"tag_filters":            me.TagFilters,
		"threshold_milliseconds": me.ThresholdMilliseconds,
		"threshold_percent":      me.ThresholdPercent,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"disk_name_filter":       &me.DiskNameFilter,
		"enabled":                &me.Enabled,
		"metric":                 &me.Metric,
		"name":                   &me.Name,
		"sample_limit":           &me.SampleLimit,
		"tag_filters":            &me.TagFilters,
		"threshold_milliseconds": &me.ThresholdMilliseconds,
		"threshold_percent":      &me.ThresholdPercent,
	})
}
