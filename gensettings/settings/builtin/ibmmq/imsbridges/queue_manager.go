package imsbridges

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type QueueManagers []*QueueManager

func (me *QueueManagers) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"queueManager": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(QueueManager).Schema()},
		},
	}
}

func (me QueueManagers) MarshalHCL() (map[string]interface{}, error) {
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
		result["queueManager"] = entries
	}
	return result, nil
}

func (me *QueueManagers) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("queueManager"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(QueueManager)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "queueManager", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type QueueManager struct {
	Name              string   `json:"name"`              // Queue manager name
	QueueManagerQueue []string `json:"queueManagerQueue"` // Queues
}

func (me *QueueManager) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "Queue manager name",
			Required:    true,
		},
		"queue_manager_queue": {
			Type:        hcl.TypeSet,
			Description: "Queues",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
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
	return decoder.DecodeAll(map[string]any{
		"name":                &me.Name,
		"queue_manager_queue": &me.QueueManagerQueue,
	})
}
