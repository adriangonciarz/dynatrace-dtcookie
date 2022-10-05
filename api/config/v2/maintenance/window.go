package maintenance

import (
	"encoding/json"
	"sort"

	"github.com/dtcookie/hcl"
)

// MaintenanceWindow TODO: documentation
type MaintenanceWindow struct {
	Enabled           bool               `json:"enabled"`           // The maintenance window is enabled or disabled
	GeneralProperties *GeneralProperties `json:"generalProperties"` // The general properties of the maintenance window
	Schedule          *Schedule          `json:"schedule"`          // The schedule of the maintenance window
	Filters           Filters            `json:"filters,omitempty"` // The filters of the maintenance window
}

func (me *MaintenanceWindow) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "The maintenance window is enabled or disabled",
			Default:     true,
			Optional:    true,
		},
		"general_properties": {
			Type:        hcl.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "The general properties of the maintenance window",
			Elem: &hcl.Resource{
				Schema: new(GeneralProperties).Schema(),
			},
		},
		"schedule": {
			Type:        hcl.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "The schedule of the maintenance window",
			Elem: &hcl.Resource{
				Schema: new(Schedule).Schema(),
			},
		},
		"filters": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "The filters of the maintenance window",
			Elem: &hcl.Resource{
				Schema: new(Filters).Schema(),
			},
		},
	}
}

func (me *MaintenanceWindow) EnsurePredictableOrder() {
	if len(me.Filters) > 0 {
		conds := []*Filter{}
		condStrings := sort.StringSlice{}
		for _, entry := range me.Filters {
			condBytes, _ := json.Marshal(entry)
			condStrings = append(condStrings, string(condBytes))
		}
		condStrings.Sort()
		for _, condString := range condStrings {
			cond := Filter{}
			json.Unmarshal([]byte(condString), &cond)
			conds = append(conds, &cond)
		}
		me.Filters = conds
	}
}

func (me *MaintenanceWindow) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	me.EnsurePredictableOrder()

	return properties.EncodeAll(map[string]interface{}{
		"enabled":            me.Enabled,
		"general_properties": me.GeneralProperties,
		"schedule":           me.Schedule,
		"filters":            me.Filters,
	})
}

func (me *MaintenanceWindow) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"enabled":            &me.Enabled,
		"general_properties": &me.GeneralProperties,
		"schedule":           &me.Schedule,
		"filters":            &me.Filters,
	})
}
