package imsbridges

import (
	"encoding/json"
	"sort"

	"github.com/dtcookie/hcl"
)

// Filters TODO: documentation
type IMSBridge struct {
	Name          string        `json:"name"`
	QueueManagers QueueManagers `json:"queueManagers,omitempty"`
}

func (me *IMSBridge) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the IMS bridge",
		},
		"queue_managers": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "Queue manager(s) that belong to the IMS bridge",
			Elem: &hcl.Resource{
				Schema: new(QueueManagers).Schema(),
			},
		},
	}
}

func (me *IMSBridge) EnsurePredictableOrder() {
	if len(me.QueueManagers) > 0 {
		conds := []*QueueManager{}
		condStrings := sort.StringSlice{}
		for _, entry := range me.QueueManagers {
			condBytes, _ := json.Marshal(entry)
			condStrings = append(condStrings, string(condBytes))
		}
		condStrings.Sort()
		for _, condString := range condStrings {
			cond := QueueManager{}
			json.Unmarshal([]byte(condString), &cond)
			conds = append(conds, &cond)
		}
		me.QueueManagers = conds
	}
}

func (me *IMSBridge) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	me.EnsurePredictableOrder()

	return properties.EncodeAll(map[string]interface{}{
		"name":           me.Name,
		"queue_managers": me.QueueManagers,
	})
}

func (me *IMSBridge) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":           &me.Name,
		"queue_managers": &me.QueueManagers,
	})
}
