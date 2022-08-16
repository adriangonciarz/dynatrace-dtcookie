package queuemanagers

import (
	"encoding/json"
	"sort"

	"github.com/dtcookie/hcl"
)

// QueueManager TODO: documentation
type QueueManager struct {
	Name          string        `json:"name"`
	Clusters      []string      `json:"clusters"`
	AliasQueues   AliasQueues   `json:"aliasQueues,omitempty"`
	RemoteQueues  RemoteQueues  `json:"remoteQueues,omitempty"`
	ClusterQueues ClusterQueues `json:"clusterQueues,omitempty"`
}

func (me *QueueManager) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the queue manager",
		},
		"clusters": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "Name of the cluster(s) this queue manager is part of",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"alias_queues": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "The alias queues in the queue manager",
			Elem: &hcl.Resource{
				Schema: new(AliasQueues).Schema(),
			},
		},
		"remote_queues": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "The alias queues in the queue manager",
			Elem: &hcl.Resource{
				Schema: new(RemoteQueues).Schema(),
			},
		},
		"cluster_queues": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "The alias queues in the queue manager",
			Elem: &hcl.Resource{
				Schema: new(ClusterQueues).Schema(),
			},
		},
	}
}

func (me *QueueManager) EnsurePredictableOrder() {
	if len(me.AliasQueues) > 0 {
		conds := []*AliasQueue{}
		condStrings := sort.StringSlice{}
		for _, entry := range me.AliasQueues {
			condBytes, _ := json.Marshal(entry)
			condStrings = append(condStrings, string(condBytes))
		}
		condStrings.Sort()
		for _, condString := range condStrings {
			cond := AliasQueue{}
			json.Unmarshal([]byte(condString), &cond)
			conds = append(conds, &cond)
		}
		me.AliasQueues = conds
	}
	if len(me.RemoteQueues) > 0 {
		conds := []*RemoteQueue{}
		condStrings := sort.StringSlice{}
		for _, entry := range me.RemoteQueues {
			condBytes, _ := json.Marshal(entry)
			condStrings = append(condStrings, string(condBytes))
		}
		condStrings.Sort()
		for _, condString := range condStrings {
			cond := RemoteQueue{}
			json.Unmarshal([]byte(condString), &cond)
			conds = append(conds, &cond)
		}
		me.RemoteQueues = conds
	}
	if len(me.ClusterQueues) > 0 {
		conds := []*ClusterQueue{}
		condStrings := sort.StringSlice{}
		for _, entry := range me.ClusterQueues {
			condBytes, _ := json.Marshal(entry)
			condStrings = append(condStrings, string(condBytes))
		}
		condStrings.Sort()
		for _, condString := range condStrings {
			cond := ClusterQueue{}
			json.Unmarshal([]byte(condString), &cond)
			conds = append(conds, &cond)
		}
		me.ClusterQueues = conds
	}
}

func (me *QueueManager) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	me.EnsurePredictableOrder()

	return properties.EncodeAll(map[string]interface{}{
		"name":           me.Name,
		"clusters":       me.Clusters,
		"alias_queues":   me.AliasQueues,
		"remote_queues":  me.RemoteQueues,
		"cluster_queues": me.ClusterQueues,
	})
}

func (me *QueueManager) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":           &me.Name,
		"clusters":       &me.Clusters,
		"alias_queues":   &me.AliasQueues,
		"remote_queues":  &me.RemoteQueues,
		"cluster_queues": &me.ClusterQueues,
	})
}
