package queuemanagers

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type RemoteQueues []*RemoteQueue

func (me *RemoteQueues) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"remoteQueue": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(RemoteQueue).Schema()},
		},
	}
}

func (me RemoteQueues) MarshalHCL() (map[string]interface{}, error) {
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
		result["remoteQueue"] = entries
	}
	return result, nil
}

func (me *RemoteQueues) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("remoteQueue"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(RemoteQueue)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "remoteQueue", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type RemoteQueue struct {
	ClusterVisibility  []string `json:"clusterVisibility"`  // Name of the cluster(s) this local definition of the remote queue should be visible in
	LocalQueue         string   `json:"localQueue"`         // Local queue name
	RemoteQueue        string   `json:"remoteQueue"`        // Remote queue name
	RemoteQueueManager string   `json:"remoteQueueManager"` // Remote queue manager name
}

func (me *RemoteQueue) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cluster_visibility": {
			Type:        hcl.TypeSet,
			Description: "Name of the cluster(s) this local definition of the remote queue should be visible in",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"local_queue": {
			Type:        hcl.TypeString,
			Description: "Local queue name",
			Required:    true,
		},
		"remote_queue": {
			Type:        hcl.TypeString,
			Description: "Remote queue name",
			Required:    true,
		},
		"remote_queue_manager": {
			Type:        hcl.TypeString,
			Description: "Remote queue manager name",
			Required:    true,
		},
	}
}

func (me *RemoteQueue) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"cluster_visibility":   me.ClusterVisibility,
		"local_queue":          me.LocalQueue,
		"remote_queue":         me.RemoteQueue,
		"remote_queue_manager": me.RemoteQueueManager,
	})
}

func (me *RemoteQueue) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cluster_visibility":   &me.ClusterVisibility,
		"local_queue":          &me.LocalQueue,
		"remote_queue":         &me.RemoteQueue,
		"remote_queue_manager": &me.RemoteQueueManager,
	})
}
