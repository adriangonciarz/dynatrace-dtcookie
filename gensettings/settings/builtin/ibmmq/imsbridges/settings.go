package imsbridges

import "github.com/dtcookie/hcl"

type Settings struct {
	Name          string        `json:"name"`          // IMS bridge name
	QueueManagers QueueManagers `json:"queueManagers"` // Queue managers
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "IMS bridge name",
			Required:    true,
		},
		"queue_managers": {
			Type:        hcl.TypeList,
			Description: "Queue managers",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(QueueManagers).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"name":           me.Name,
		"queue_managers": me.QueueManagers,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"name":           &me.Name,
		"queue_managers": &me.QueueManagers,
	})
}
