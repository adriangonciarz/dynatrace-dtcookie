package externalwebservice

import "github.com/dtcookie/hcl"

type IdContributorsType struct {
	DetectAsWebRequestService bool                  `json:"detectAsWebRequestService"` // Detect the matching requests as web request services instead of web services.\n\nThis prevents detecting of matching requests as opaque web services. An opaque web request service is created instead. If you need to further modify the resulting web request service, you need to create a separate [Opaque/external web request rule](builtin:service-detection.full-web-request).
	PortForServiceID          bool                  `json:"portForServiceId"`          // Let the port contribute to the service identifier calculation.
	UrlPath                   *ServiceIdContributor `json:"urlPath"`                   // Let the request's URL contribute to the service identifier calculation.
}

func (me *IdContributorsType) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"detect_as_web_request_service": {
			Type:        hcl.TypeBool,
			Description: "Detect the matching requests as web request services instead of web services.\n\nThis prevents detecting of matching requests as opaque web services. An opaque web request service is created instead. If you need to further modify the resulting web request service, you need to create a separate [Opaque/external web request rule](builtin:service-detection.full-web-request).",
			Optional:    true,
		},
		"port_for_service_id": {
			Type:        hcl.TypeBool,
			Description: "Let the port contribute to the service identifier calculation.",
			Optional:    true,
		},
		"url_path": {
			Type:        hcl.TypeList,
			Description: "Let the request's URL contribute to the service identifier calculation.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ServiceIdContributor).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *IdContributorsType) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"detect_as_web_request_service": me.DetectAsWebRequestService,
		"port_for_service_id":           me.PortForServiceID,
		"url_path":                      me.UrlPath,
	})
}

func (me *IdContributorsType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"detect_as_web_request_service": &me.DetectAsWebRequestService,
		"port_for_service_id":           &me.PortForServiceID,
		"url_path":                      &me.UrlPath,
	})
}
