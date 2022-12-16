package usersessionexportsettingsv2

import "github.com/dtcookie/hcl"

type EndpointDefinition struct {
	ContentType             string `json:"contentType"`             // Content type
	EnableUserSessionExport bool   `json:"enableUserSessionExport"` // Enable user session export
	EndpointUrl             string `json:"endpointUrl"`             // Endpoint URL
	UsePost                 bool   `json:"usePost"`                 // Use POST method (instead of PUT)
}

func (me *EndpointDefinition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"content_type": {
			Type:        hcl.TypeString,
			Description: "Content type",
			Required:    true,
		},
		"enable_user_session_export": {
			Type:        hcl.TypeBool,
			Description: "Enable user session export",
			Optional:    true,
		},
		"endpoint_url": {
			Type:        hcl.TypeString,
			Description: "Endpoint URL",
			Optional:    true,
		},
		"use_post": {
			Type:        hcl.TypeBool,
			Description: "Use POST method (instead of PUT)",
			Optional:    true,
		},
	}
}

func (me *EndpointDefinition) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"content_type":               me.ContentType,
		"enable_user_session_export": me.EnableUserSessionExport,
		"endpoint_url":               me.EndpointUrl,
		"use_post":                   me.UsePost,
	})
}

func (me *EndpointDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"content_type":               &me.ContentType,
		"enable_user_session_export": &me.EnableUserSessionExport,
		"endpoint_url":               &me.EndpointUrl,
		"use_post":                   &me.UsePost,
	})
}
