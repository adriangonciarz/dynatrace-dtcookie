package queuemanagers

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ClusterQueues []*ClusterQueue

func (me *ClusterQueues) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"clusterQueue": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(ClusterQueue).Schema()},
		},
	}
}

func (me ClusterQueues) MarshalHCL() (map[string]interface{}, error) {
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
		result["clusterQueue"] = entries
	}
	return result, nil
}

func (me *ClusterQueues) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("clusterQueue"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(ClusterQueue)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "clusterQueue", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type ClusterQueue struct {
	ClusterVisibility []string `json:"clusterVisibility"` // Name of the cluster(s) this local queue should be visible in
	LocalQueue        string   `json:"localQueue"`        // Local queue name
}

func (me *ClusterQueue) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cluster_visibility": {
			Type:        hcl.TypeSet,
			Description: "Name of the cluster(s) this local queue should be visible in",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"local_queue": {
			Type:        hcl.TypeString,
			Description: "Local queue name",
			Required:    true,
		},
	}
}

func (me *ClusterQueue) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"cluster_visibility": me.ClusterVisibility,
		"local_queue":        me.LocalQueue,
	})
}

func (me *ClusterQueue) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cluster_visibility": &me.ClusterVisibility,
		"local_queue":        &me.LocalQueue,
	})
}
