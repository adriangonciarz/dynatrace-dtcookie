package detectionrules

import "github.com/dtcookie/hcl"

type Settings struct {
	ApiColor      string      `json:"apiColor"`      // This color will be used to highlight APIs when viewing code level data, such as distributed traces or method hotspots.
	ApiName       string      `json:"apiName"`       // API name
	Conditions    ApiRules    `json:"conditions"`    // List of conditions
	Technology    *Technology `json:"technology"`    // Restrict this rule to a specific technology.
	ThirdPartyApi bool        `json:"thirdPartyApi"` // This API defines a third party library
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"api_color": {
			Type:        hcl.TypeString,
			Description: "This color will be used to highlight APIs when viewing code level data, such as distributed traces or method hotspots.",
			Required:    true,
		},
		"api_name": {
			Type:        hcl.TypeString,
			Description: "API name",
			Required:    true,
		},
		"conditions": {
			Type:        hcl.TypeList,
			Description: "List of conditions",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ApiRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"technology": {
			Type:        hcl.TypeString,
			Description: "Restrict this rule to a specific technology.",
			Optional:    true,
		},
		"third_party_api": {
			Type:        hcl.TypeBool,
			Description: "This API defines a third party library",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"api_color":       me.ApiColor,
		"api_name":        me.ApiName,
		"conditions":      me.Conditions,
		"technology":      me.Technology,
		"third_party_api": me.ThirdPartyApi,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"api_color":       &me.ApiColor,
		"api_name":        &me.ApiName,
		"conditions":      &me.Conditions,
		"technology":      &me.Technology,
		"third_party_api": &me.ThirdPartyApi,
	})
}
