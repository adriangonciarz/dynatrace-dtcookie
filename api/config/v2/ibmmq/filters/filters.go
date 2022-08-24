package filters

import (
	"github.com/dtcookie/hcl"
)

// Filters TODO: documentation
type Filters struct {
	CICSMQQueueIdIncludes []string `json:"cicsMqQueueIdIncludes,omitempty"`
	CICSMQQueueIdExcludes []string `json:"cicsMqQueueIdExcludes,omitempty"`
	IMSMQQueueIdIncludes  []string `json:"imsMqQueueIdIncludes,omitempty"`
	IMSMQQueueIdExcludes  []string `json:"imsMqQueueIdExcludes,omitempty"`
	IMSCrTrnIdIncludes    []string `json:"imsCrTrnIdIncludes,omitempty"`
	IMSCrTrnIdExcludes    []string `json:"imsCrTrnIdExcludes,omitempty"`
}

func (me *Filters) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cics_mq_queue_id_includes": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "CICS: Included MQ queues",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"cics_mq_queue_id_excludes": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "CICS: Excluded MQ queues",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"ims_mq_queue_id_includes": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "IMS: Included MQ queues",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"ims_mq_queue_id_excludes": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "IMS: Excluded MQ queues",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"ims_cr_trn_id_includes": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "IMS bridge: Included transaction IDs",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"ims_cr_trn_id_excludes": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "IMS bridge: Excluded transaction IDs",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Filters) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"cics_mq_queue_id_includes": me.CICSMQQueueIdIncludes,
		"cics_mq_queue_id_excludes": me.CICSMQQueueIdExcludes,
		"ims_mq_queue_id_includes":  me.IMSMQQueueIdIncludes,
		"ims_mq_queue_id_excludes":  me.IMSMQQueueIdExcludes,
		"ims_cr_trn_id_includes":    me.IMSCrTrnIdIncludes,
		"ims_cr_trn_id_excludes":    me.IMSCrTrnIdExcludes,
	})
}

func (me *Filters) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"cics_mq_queue_id_includes": &me.CICSMQQueueIdIncludes,
		"cics_mq_queue_id_excludes": &me.CICSMQQueueIdExcludes,
		"ims_mq_queue_id_includes":  &me.IMSMQQueueIdIncludes,
		"ims_mq_queue_id_excludes":  &me.IMSMQQueueIdExcludes,
		"ims_cr_trn_id_includes":    &me.IMSCrTrnIdIncludes,
		"ims_cr_trn_id_excludes":    &me.IMSCrTrnIdExcludes,
	})
}
