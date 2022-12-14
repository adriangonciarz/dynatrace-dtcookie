package dashboards

import (
	"encoding/json"
	"sort"

	"github.com/dtcookie/hcl"
)

type JSONDashboard struct {
	Name     string
	Contents string
}

func (me *JSONDashboard) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"contents": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "Contains the JSON Code of the Dashboard",
		},
	}
}

func (me *JSONDashboard) MarshalHCL() (map[string]interface{}, error) {
	m := map[string]interface{}{}
	m["contents"] = me.Contents
	return m, nil
}

func (me *JSONDashboard) UnmarshalHCL(decoder hcl.Decoder) error {
	v, ok := decoder.GetOk("contents")
	if ok {
		me.Contents = v.(string)
	}
	return nil
}

func (me *JSONDashboard) MarshalJSON() ([]byte, error) {
	return []byte(me.Contents), nil
}

func (me *JSONDashboard) UnmarshalJSON(data []byte) error {
	reduced := struct {
		Metadata *DashboardMetadata       `json:"dashboardMetadata"`
		Tiles    []map[string]interface{} `json:"tiles"`
	}{}

	if err := json.Unmarshal(data, &reduced); err != nil {
		return err
	}

	var err error
	for _, tile := range reduced.Tiles {
		var untypedQueries interface{}
		var found bool
		if untypedQueries, found = tile["queries"]; found {
			queriesArr := untypedQueries.([]interface{})
			for _, untypedQuery := range queriesArr {
				query := untypedQuery.(map[string]interface{})
				var untypedFilterBy interface{}
				if untypedFilterBy, found = query["filterBy"]; found {
					filterBy := untypedFilterBy.(map[string]interface{})
					var untypedNestedFilters interface{}
					if untypedNestedFilters, found = filterBy["nestedFilters"]; found {
						nestedFilters := untypedNestedFilters.([]interface{})
						strs := []string{}
						for _, nestedFilter := range nestedFilters {
							var data []byte
							if data, err = json.Marshal(nestedFilter); err != nil {
								return err
							}
							strs = append(strs, string(data))
							sort.Strings(strs)
							nestedFilters = []interface{}{}
							for _, str := range strs {
								nestedFilters = append(nestedFilters, json.RawMessage([]byte(str)))
							}
							filterBy["nestedFilters"] = nestedFilters
						}
					}

				}
			}
		}
	}

	if data, err = json.Marshal(reduced); err != nil {
		return err
	}

	md := struct {
		Metadata *DashboardMetadata `json:"dashboardMetadata"`
	}{}
	err = json.Unmarshal(data, &md)
	if err != nil {
		return err
	}
	me.Name = md.Metadata.Name
	me.Contents = string(data)
	return nil
}
