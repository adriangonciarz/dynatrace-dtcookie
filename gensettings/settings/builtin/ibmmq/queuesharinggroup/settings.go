package queuesharinggroup

import "github.com/dtcookie/hcl"

type Settings struct {
	Name          string   `json:"name"`          // Queue sharing group name
	QueueManagers []string `json:"queueManagers"` // Queue managers
	SharedQueues  []string `json:"sharedQueues"`  // Shared queues
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "Queue sharing group name",
			Required:    true,
		},
		"queue_managers": {
			Type:        hcl.TypeSet,
			Description: "Queue managers",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"shared_queues": {
			Type:        hcl.TypeSet,
			Description: "Shared queues",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"name":           me.Name,
		"queue_managers": me.QueueManagers,
		"shared_queues":  me.SharedQueues,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"name":           &me.Name,
		"queue_managers": &me.QueueManagers,
		"shared_queues":  &me.SharedQueues,
	})
}
