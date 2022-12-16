package httpparameters

import "github.com/dtcookie/hcl"

type BrokenLinks struct {
	BrokenLinkDomains       []string `json:"brokenLinkDomains"`       // If your application relies on other hosts at other domains, add the associated domain names here. Once configured, Dynatrace will consider 404s thrown by hosts at these domains to be service failures related to your application.
	Http404NotFoundFailures bool     `json:"http404NotFoundFailures"` // Consider 404 HTTP response codes as failures
}

func (me *BrokenLinks) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"broken_link_domains": {
			Type:        hcl.TypeSet,
			Description: "If your application relies on other hosts at other domains, add the associated domain names here. Once configured, Dynatrace will consider 404s thrown by hosts at these domains to be service failures related to your application.",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"http_404_not_found_failures": {
			Type:        hcl.TypeBool,
			Description: "Consider 404 HTTP response codes as failures",
			Optional:    true,
		},
	}
}

func (me *BrokenLinks) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"broken_link_domains":         me.BrokenLinkDomains,
		"http_404_not_found_failures": me.Http404NotFoundFailures,
	})
}

func (me *BrokenLinks) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"broken_link_domains":         &me.BrokenLinkDomains,
		"http_404_not_found_failures": &me.Http404NotFoundFailures,
	})
}
