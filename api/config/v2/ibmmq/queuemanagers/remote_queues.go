package queuemanagers

import (
	"github.com/dtcookie/hcl"
)

type RemoteQueue struct {
	LocalQueueName     string   `json:"localQueue"`
	RemoteQueueName    string   `json:"remoteQueue"`
	RemoteQueueManager string   `json:"remoteQueueManager"`
	ClusterVisibility  []string `json:"clusterVisibility,omitempty"`
}

type RemoteQueues []*RemoteQueue

func (me *RemoteQueue) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"local_queue_name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the local queue",
		},
		"remote_queue_name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the remote queue",
		},
		"remote_queue_manager": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the remote queue manager",
		},
		"cluster_visibility": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "Name of the cluster(s) this local definition of the remote queue should be visible in",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *RemoteQueues) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"remote_queue": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "Remote queue definitions for queue manager",
			Elem:        &hcl.Resource{Schema: new(RemoteQueue).Schema()},
		},
	}
}

func (me *RemoteQueue) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"local_queue_name":     me.LocalQueueName,
		"remote_queue_name":    me.RemoteQueueName,
		"remote_queue_manager": me.RemoteQueueManager,
		"cluster_visibility":   me.ClusterVisibility,
	})
}

func (me *RemoteQueue) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"local_queue_name":     &me.LocalQueueName,
		"remote_queue_name":    &me.RemoteQueueName,
		"remote_queue_manager": &me.RemoteQueueManager,
		"cluster_visibility":   &me.ClusterVisibility,
	})
}

func (me RemoteQueues) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("remote_queue", me)
}

func (me *RemoteQueues) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("remote_queue", me)
}
