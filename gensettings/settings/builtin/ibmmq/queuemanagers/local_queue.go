package queuemanagers

import "github.com/dtcookie/hcl"

type LocalQueue struct {
	LocalQueue string `json:"localQueue"` // Local queue
}

func (me *LocalQueue) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"local_queue": {
			Type:        hcl.TypeString,
			Description: "Local queue",
			Required:    true,
		},
	}
}

func (me *LocalQueue) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"local_queue": me.LocalQueue,
	})
}

func (me *LocalQueue) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"local_queue": &me.LocalQueue,
	})
}
