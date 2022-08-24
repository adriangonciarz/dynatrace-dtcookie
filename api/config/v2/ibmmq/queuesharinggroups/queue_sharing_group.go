package queuesharinggroups

import (
	"github.com/dtcookie/hcl"
)

// Filters TODO: documentation
type QueueSharingGroup struct {
	Name          string   `json:"name"`
	QueueManagers []string `json:"queueManagers,omitempty"`
	SharedQueues  []string `json:"sharedQueues,omitempty"`
}

func (me *QueueSharingGroup) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the queue sharing group",
		},
		"queue_managers": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "Queue manager(s) that belong to the queue sharing group",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"shared_queues": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "Shared queue(s) that belong to the queue sharing group",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *QueueSharingGroup) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"name":           me.Name,
		"queue_managers": me.QueueManagers,
		"shared_queues":  me.SharedQueues,
	})
}

func (me *QueueSharingGroup) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":           &me.Name,
		"queue_managers": &me.QueueManagers,
		"shared_queues":  &me.SharedQueues,
	})
}
