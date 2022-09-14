package alerting

import (
	"log"

	"github.com/dtcookie/hcl"
)

type SeverityRules []*SeverityRule

func (me *SeverityRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rule": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "A conditions for the metric usage",
			Elem:        &hcl.Resource{Schema: new(SeverityRule).Schema()},
		},
	}
}

func (me SeverityRules) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("rule", me)
}

func (me *SeverityRules) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("rule", me)
}

type SeverityRule struct {
	SeverityLevel        SeverityLevel        `json:"severityLevel"`        // Problem severity level
	DelayInMinutes       int32                `json:"delayInMinutes"`       // Send a notification if a problem remains open longer than X minutes. Must be between 0 and 10000.
	TagFilterIncludeMode TagFilterIncludeMode `json:"tagFilterIncludeMode"` // Possible values are `NONE`, `INCLUDE_ANY` and `INCLUDE_ALL`
	Tags                 []string             `json:"tagFilter"`            // SET / no documentation available
}

func (me *SeverityRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"severity_level": {
			Type:        hcl.TypeString,
			Description: "The severity level to trigger the alert. Possible values are `AVAILABILITY`,	`CUSTOM_ALERT`,	`ERRORS`,`MONITORING_UNAVAILABLE`,`PERFORMANCE` and `RESOURCE_CONTENTION`.",
			Required:    true,
		},
		"delay_in_minutes": {
			Type:        hcl.TypeInt,
			Description: "Send a notification if a problem remains open longer than *X* minutes",
			Required:    true,
		},
		"include_mode": {
			Type:        hcl.TypeString,
			Description: "The filtering mode:  * `INCLUDE_ANY`: The rule applies to monitored entities that have at least one of the specified tags. You can specify up to 100 tags.  * `INCLUDE_ALL`: The rule applies to monitored entities that have **all** of the specified tags. You can specify up to 10 tags.  * `NONE`: The rule applies to all monitored entities",
			Required:    true,
		},
		"tags": {
			Type:        hcl.TypeSet,
			Description: "A set of tags you want to filter by. You can also specify a tag value alongside the tag name using the syntax `name:value`.",
			MinItems:    1,
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString}},
	}
}

func (me *SeverityRule) MarshalHCL() (map[string]interface{}, error) {
	log.Println("SeverityRule", "MarshalHCL")
	result := map[string]interface{}{}

	result["delay_in_minutes"] = int(me.DelayInMinutes)
	result["severity_level"] = string(me.SeverityLevel)
	result["include_mode"] = string(me.TagFilterIncludeMode)

	if len(me.Tags) > 0 {
		result["tags"] = me.Tags
	}
	return result, nil
}

func (me *SeverityRule) UnmarshalHCL(decoder hcl.Decoder) error {
	log.Println("SeverityRule", "UnmarshalHCL")
	if value, ok := decoder.GetOk("severity_level"); ok {
		me.SeverityLevel = SeverityLevel(value.(string))
	}
	if value, ok := decoder.GetOk("include_mode"); ok {
		me.TagFilterIncludeMode = TagFilterIncludeMode(value.(string))
	}
	if value, ok := decoder.GetOk("delay_in_minutes"); ok {
		me.DelayInMinutes = int32(value.(int))
	}
	if value, ok := decoder.GetOk("tags"); ok {
		me.Tags = []string{}
		for _, entry := range value.(hcl.Set).List() {
			me.Tags = append(me.Tags, entry.(string))
		}
	}
	return nil
}
