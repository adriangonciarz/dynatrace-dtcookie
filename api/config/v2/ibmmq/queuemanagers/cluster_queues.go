package queuemanagers

import (
	"github.com/dtcookie/hcl"
)

type ClusterQueue struct {
	LocalQueueName    string   `json:"localQueue"`
	ClusterVisibility []string `json:"clusterVisibility,omitempty"`
}

type ClusterQueues []*ClusterQueue

func (me *ClusterQueue) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"local_queue_name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the local queue",
		},
		"cluster_visibility": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "Name of the cluster(s) this local queue should be visible in",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *ClusterQueues) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cluster_queue": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "Cluster queue definitions for queue manager",
			Elem:        &hcl.Resource{Schema: new(ClusterQueue).Schema()},
		},
	}
}

func (me *ClusterQueue) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"local_queue_name":   me.LocalQueueName,
		"cluster_visibility": me.ClusterVisibility,
	})
}

func (me *ClusterQueue) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"local_queue_name":   &me.LocalQueueName,
		"cluster_visibility": &me.ClusterVisibility,
	})
}

func (me ClusterQueues) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("cluster_queue", me)
}

func (me *ClusterQueues) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("cluster_queue", me)
}
