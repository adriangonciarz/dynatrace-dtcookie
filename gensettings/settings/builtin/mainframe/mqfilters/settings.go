package mqfilters

import "github.com/dtcookie/hcl"

type Settings struct {
	CicsMqQueueIdExcludes []string `json:"cicsMqQueueIdExcludes"` // CICS: Excluded MQ queues
	CicsMqQueueIdIncludes []string `json:"cicsMqQueueIdIncludes"` // CICS: Included MQ queues
	ImsCrTrnIdExcludes    []string `json:"imsCrTrnIdExcludes"`    // When you add a transaction ID to the exclude list remaining transactions are still monitored.
	ImsCrTrnIdIncludes    []string `json:"imsCrTrnIdIncludes"`    // When you add a transaction ID to the include list, all the remaining transactions are ignored.
	ImsMqQueueIdExcludes  []string `json:"imsMqQueueIdExcludes"`  // IMS: Excluded MQ queues
	ImsMqQueueIdIncludes  []string `json:"imsMqQueueIdIncludes"`  // IMS: Included MQ queues
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cics_mq_queue_id_excludes": {
			Type:        hcl.TypeSet,
			Description: "CICS: Excluded MQ queues",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"cics_mq_queue_id_includes": {
			Type:        hcl.TypeSet,
			Description: "CICS: Included MQ queues",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"ims_cr_trn_id_excludes": {
			Type:        hcl.TypeSet,
			Description: "When you add a transaction ID to the exclude list remaining transactions are still monitored.",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"ims_cr_trn_id_includes": {
			Type:        hcl.TypeSet,
			Description: "When you add a transaction ID to the include list, all the remaining transactions are ignored.",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"ims_mq_queue_id_excludes": {
			Type:        hcl.TypeSet,
			Description: "IMS: Excluded MQ queues",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"ims_mq_queue_id_includes": {
			Type:        hcl.TypeSet,
			Description: "IMS: Included MQ queues",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"cics_mq_queue_id_excludes": me.CicsMqQueueIdExcludes,
		"cics_mq_queue_id_includes": me.CicsMqQueueIdIncludes,
		"ims_cr_trn_id_excludes":    me.ImsCrTrnIdExcludes,
		"ims_cr_trn_id_includes":    me.ImsCrTrnIdIncludes,
		"ims_mq_queue_id_excludes":  me.ImsMqQueueIdExcludes,
		"ims_mq_queue_id_includes":  me.ImsMqQueueIdIncludes,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cics_mq_queue_id_excludes": &me.CicsMqQueueIdExcludes,
		"cics_mq_queue_id_includes": &me.CicsMqQueueIdIncludes,
		"ims_cr_trn_id_excludes":    &me.ImsCrTrnIdExcludes,
		"ims_cr_trn_id_includes":    &me.ImsCrTrnIdIncludes,
		"ims_mq_queue_id_excludes":  &me.ImsMqQueueIdExcludes,
		"ims_mq_queue_id_includes":  &me.ImsMqQueueIdIncludes,
	})
}
