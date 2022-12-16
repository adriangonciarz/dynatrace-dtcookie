package performancethresholds

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ThresholdEntries []*ThresholdEntry

func (me *ThresholdEntries) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"threshold": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(ThresholdEntry).Schema()},
		},
	}
}

func (me ThresholdEntries) MarshalHCL() (map[string]interface{}, error) {
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
		result["threshold"] = entries
	}
	return result, nil
}

func (me *ThresholdEntries) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("threshold"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(ThresholdEntry)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "threshold", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type ThresholdEntry struct {
	Event     string  `json:"event"`     // Synthetic event
	Threshold float64 `json:"threshold"` // Threshold (in seconds)
}

func (me *ThresholdEntry) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"event": {
			Type:        hcl.TypeString,
			Description: "Synthetic event",
			Required:    true,
		},
		"threshold": {
			Type:        hcl.TypeFloat,
			Description: "Threshold (in seconds)",
			Required:    true,
		},
	}
}

func (me *ThresholdEntry) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"event":     me.Event,
		"threshold": me.Threshold,
	})
}

func (me *ThresholdEntry) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"event":     &me.Event,
		"threshold": &me.Threshold,
	})
}
