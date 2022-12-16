package queuemanagers

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AliasQueues []*AliasQueue

func (me *AliasQueues) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"aliasQueue": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(AliasQueue).Schema()},
		},
	}
}

func (me AliasQueues) MarshalHCL() (map[string]interface{}, error) {
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
		result["aliasQueue"] = entries
	}
	return result, nil
}

func (me *AliasQueues) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("aliasQueue"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(AliasQueue)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "aliasQueue", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

// AliasQueue. Alias queue
type AliasQueue struct {
	AliasQueue        string   `json:"aliasQueue"`        // Alias queue name
	BaseQueue         string   `json:"baseQueue"`         // Base queue name
	ClusterVisibility []string `json:"clusterVisibility"` // Name of the cluster(s) this alias should be visible in
}

func (me *AliasQueue) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"alias_queue": {
			Type:        hcl.TypeString,
			Description: "Alias queue name",
			Required:    true,
		},
		"base_queue": {
			Type:        hcl.TypeString,
			Description: "Base queue name",
			Required:    true,
		},
		"cluster_visibility": {
			Type:        hcl.TypeSet,
			Description: "Name of the cluster(s) this alias should be visible in",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *AliasQueue) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"alias_queue":        me.AliasQueue,
		"base_queue":         me.BaseQueue,
		"cluster_visibility": me.ClusterVisibility,
	})
}

func (me *AliasQueue) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"alias_queue":        &me.AliasQueue,
		"base_queue":         &me.BaseQueue,
		"cluster_visibility": &me.ClusterVisibility,
	})
}
