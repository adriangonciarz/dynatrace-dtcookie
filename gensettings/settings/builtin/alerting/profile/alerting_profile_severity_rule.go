package profile

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AlertingProfileSeverityRules []*AlertingProfileSeverityRule

func (me *AlertingProfileSeverityRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"severityRule": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(AlertingProfileSeverityRule).Schema()},
		},
	}
}

func (me AlertingProfileSeverityRules) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me) > 0 {
		entries := []interface{}{}
		for _, entry := range me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["severityRule"] = entries
	}
	return result, nil
}

func (me *AlertingProfileSeverityRules) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("severityRule"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(AlertingProfileSeverityRule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "severityRule", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type AlertingProfileSeverityRule struct {
	DelayInMinutes       int                  `json:"delayInMinutes"`       // Send a notification if a problem remains open longer than X minutes.
	SeverityLevel        SeverityLevel        `json:"severityLevel"`        // Problem severity level
	TagFilter            []string             `json:"tagFilter"`            // Tags
	TagFilterIncludeMode TagFilterIncludeMode `json:"tagFilterIncludeMode"` // Filter problems by tag
}

func (me *AlertingProfileSeverityRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"delay_in_minutes": {
			Type:        hcl.TypeInt,
			Description: "Send a notification if a problem remains open longer than X minutes.",
			Required:    true,
		},
		"severity_level": {
			Type:        hcl.TypeString,
			Description: "Problem severity level",
			Required:    true,
		},
		"tag_filter": {
			Type:        hcl.TypeSet,
			Description: "Tags",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"tag_filter_include_mode": {
			Type:        hcl.TypeString,
			Description: "Filter problems by tag",
			Required:    true,
		},
	}
}

func (me *AlertingProfileSeverityRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"delay_in_minutes":        me.DelayInMinutes,
		"severity_level":          me.SeverityLevel,
		"tag_filter":              me.TagFilter,
		"tag_filter_include_mode": me.TagFilterIncludeMode,
	})
}

func (me *AlertingProfileSeverityRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"delay_in_minutes":        &me.DelayInMinutes,
		"severity_level":          &me.SeverityLevel,
		"tag_filter":              &me.TagFilter,
		"tag_filter_include_mode": &me.TagFilterIncludeMode,
	})
}
