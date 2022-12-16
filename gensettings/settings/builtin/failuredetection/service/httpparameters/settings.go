package httpparameters

import "github.com/dtcookie/hcl"

type Settings struct {
	BrokenLinks       *BrokenLinks       `json:"brokenLinks"`       // HTTP 404 response codes are thrown when a web server can't find a certain page. 404s are classified as broken links on the client side and therefore aren't considered to be service failures. By enabling this setting, you can have 404s treated as server-side service failures.
	Enabled           bool               `json:"enabled"`           // Override global failure detection settings
	HttpResponseCodes *HttpResponseCodes `json:"httpResponseCodes"` // HTTP response codes
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
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Override global failure detection settings",
			Optional:    true,
		},
		"http_response_codes": {
			Type:        hcl.TypeList,
			Description: "HTTP response codes",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(HttpResponseCodes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"broken_links":        me.BrokenLinks,
		"enabled":             me.Enabled,
		"http_response_codes": me.HttpResponseCodes,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"broken_links":        &me.BrokenLinks,
		"enabled":             &me.Enabled,
		"http_response_codes": &me.HttpResponseCodes,
	})
}
