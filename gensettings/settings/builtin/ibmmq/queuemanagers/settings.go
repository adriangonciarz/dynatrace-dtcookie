package queuemanagers

import "github.com/dtcookie/hcl"

type Settings struct {
	AliasQueues   AliasQueues   `json:"aliasQueues"`   // Alias queues
	ClusterQueues ClusterQueues `json:"clusterQueues"` // Cluster queues
	Clusters      []string      `json:"clusters"`      // Name of the cluster(s) this queue manager is part of
	Name          string        `json:"name"`          // Queue manager name
	RemoteQueues  RemoteQueues  `json:"remoteQueues"`  // Remote queues
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"alias_queues": {
			Type:        hcl.TypeList,
			Description: "Alias queues",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(AliasQueues).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"cluster_queues": {
			Type:        hcl.TypeList,
			Description: "Cluster queues",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ClusterQueues).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"clusters": {
			Type:        hcl.TypeSet,
			Description: "Name of the cluster(s) this queue manager is part of",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "Queue manager name",
			Required:    true,
		},
		"remote_queues": {
			Type:        hcl.TypeList,
			Description: "Remote queues",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(RemoteQueues).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"alias_queues":   me.AliasQueues,
		"cluster_queues": me.ClusterQueues,
		"clusters":       me.Clusters,
		"name":           me.Name,
		"remote_queues":  me.RemoteQueues,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"alias_queues":   &me.AliasQueues,
		"cluster_queues": &me.ClusterQueues,
		"clusters":       &me.Clusters,
		"name":           &me.Name,
		"remote_queues":  &me.RemoteQueues,
	})
}
