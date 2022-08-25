package imsbridges

import (
	"github.com/dtcookie/hcl"
)

type QueueManager struct {
	Name              string   `json:"name"`
	QueueManagerQueue []string `json:"queueManagerQueue,omitempty"`
}

type QueueManagers []*QueueManager

func (me *QueueManager) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the queue manager",
		},
		"queue_manager_queue": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "Queue(s) that belong to the queue manager",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *QueueManagers) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"queue_manager": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "Queue manager definition for IMS bridge",
			Elem:        &hcl.Resource{Schema: new(QueueManager).Schema()},
		},
	}
}

func (me *QueueManager) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"name":                me.Name,
		"queue_manager_queue": me.QueueManagerQueue,
	})
}

func (me *QueueManager) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":                &me.Name,
		"queue_manager_queue": &me.QueueManagerQueue,
	})
}

func (me QueueManagers) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("queue_manager", me)
}

func (me *QueueManagers) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("queue_manager", me)
}
