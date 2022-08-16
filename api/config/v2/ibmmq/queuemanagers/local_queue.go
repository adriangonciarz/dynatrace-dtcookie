package queuemanagers

import (
	"github.com/dtcookie/hcl"
)

type LocalQueue struct {
	LocalQueueName string `json:"localQueue"`
}

func (me *LocalQueue) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"local_queue_name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the local queue",
		},
	}
}

func (me *LocalQueue) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"local_queue_name": me.LocalQueueName,
	})
}

func (me *LocalQueue) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"local_queue_name": &me.LocalQueueName,
	})
}
