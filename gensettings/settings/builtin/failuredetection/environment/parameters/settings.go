package parameters

import "github.com/dtcookie/hcl"

type Settings struct {
	BrokenLinks       *BrokenLinks       `json:"brokenLinks"`       // HTTP 404 response codes are thrown when a web server can't find a certain page. 404s are classified as broken links on the client side and therefore aren't considered to be service failures. By enabling this setting, you can have 404s treated as server-side service failures.
	Description       *string            `json:"description"`       // Description
	ExceptionRules    *ExceptionRules    `json:"exceptionRules"`    // Customize failure detection for specific exceptions and errors
	HttpResponseCodes *HttpResponseCodes `json:"httpResponseCodes"` // HTTP response codes
	Name              string             `json:"name"`              // Name
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"broken_links": {
			Type:        hcl.TypeList,
			Description: "HTTP 404 response codes are thrown when a web server can't find a certain page. 404s are classified as broken links on the client side and therefore aren't considered to be service failures. By enabling this setting, you can have 404s treated as server-side service failures.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(BrokenLinks).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"description": {
			Type:        hcl.TypeString,
			Description: "Description",
			Optional:    true,
		},
		"exception_rules": {
			Type:        hcl.TypeList,
			Description: "Customize failure detection for specific exceptions and errors",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ExceptionRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"http_response_codes": {
			Type:        hcl.TypeList,
			Description: "HTTP response codes",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(HttpResponseCodes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "Name",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"broken_links":        me.BrokenLinks,
		"description":         me.Description,
		"exception_rules":     me.ExceptionRules,
		"http_response_codes": me.HttpResponseCodes,
		"name":                me.Name,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"broken_links":        &me.BrokenLinks,
		"description":         &me.Description,
		"exception_rules":     &me.ExceptionRules,
		"http_response_codes": &me.HttpResponseCodes,
		"name":                &me.Name,
	})
}
