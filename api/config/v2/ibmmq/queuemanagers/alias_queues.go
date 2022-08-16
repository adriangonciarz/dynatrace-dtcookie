package queuemanagers

import (
	"github.com/dtcookie/hcl"
)

type AliasQueue struct {
	AliasQueueName    string   `json:"aliasQueue"`
	BaseQueueName     string   `json:"baseQueue"`
	ClusterVisibility []string `json:"clusterVisibility,omitempty"`
}

type AliasQueues []*AliasQueue

func (me *AliasQueue) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"alias_queue_name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the alias queue",
		},
		"base_queue_name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the base queue",
		},
		"cluster_visibility": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "Name of the cluster(s) this alias should be visible in",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *AliasQueues) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"alias_queue": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "Alias queue definitions for queue manager",
			Elem:        &hcl.Resource{Schema: new(AliasQueue).Schema()},
		},
	}
}

func (me *AliasQueue) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"alias_queue_name":   me.AliasQueueName,
		"base_queue_name":    me.BaseQueueName,
		"cluster_visibility": me.ClusterVisibility,
	})
}

func (me *AliasQueue) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"alias_queue_name":   &me.AliasQueueName,
		"base_queue_name":    &me.BaseQueueName,
		"cluster_visibility": &me.ClusterVisibility,
	})
}

func (me AliasQueues) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("alias_queue", me)
}

func (me *AliasQueues) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("alias_queue", me)
}
