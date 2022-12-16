package fullwebservice

import "github.com/dtcookie/hcl"

type IdContributorsType struct {
	ApplicationID             *ServiceIdContributor `json:"applicationId"`             // Let the detected application identifier contribute to the service identifier calculation
	ContextRoot               *ContextIdContributor `json:"contextRoot"`               // Let the detected context root contribute to the service identifier calculation.\nThe context root is the first segment of the request URL after the server name. For example, in the www.dynatrace.com/support/help/dynatrace-api/ URL the context root is `support`.
	DetectAsWebRequestService bool                  `json:"detectAsWebRequestService"` // Detect the matching requests as full web services (false) or web request services (true).\n\nSetting this field to true prevents detecting of matching requests as full web services. A web request service is created instead. If you need to further modify the resulting web request service, you need to create a separate [Full web request rule](builtin:service-detection.full-web-request).
	ServerName                *ServiceIdContributor `json:"serverName"`                // Let the detected server name contribute to the service identifier calculation
	WebServiceName            *ServiceIdContributor `json:"webServiceName"`            // Let the detected web service name contribute to the service identifier calculation
	WebServiceNamespace       *ServiceIdContributor `json:"webServiceNamespace"`       // Let the detected web service namespace contribute to the service identifier calculation
}

func (me *IdContributorsType) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"application_id": {
			Type:        hcl.TypeList,
			Description: "Let the detected application identifier contribute to the service identifier calculation",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ServiceIdContributor).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"context_root": {
			Type:        hcl.TypeList,
			Description: "Let the detected context root contribute to the service identifier calculation.\nThe context root is the first segment of the request URL after the server name. For example, in the www.dynatrace.com/support/help/dynatrace-api/ URL the context root is `support`.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ContextIdContributor).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detect_as_web_request_service": {
			Type:        hcl.TypeBool,
			Description: "Detect the matching requests as full web services (false) or web request services (true).\n\nSetting this field to true prevents detecting of matching requests as full web services. A web request service is created instead. If you need to further modify the resulting web request service, you need to create a separate [Full web request rule](builtin:service-detection.full-web-request).",
			Optional:    true,
		},
		"server_name": {
			Type:        hcl.TypeList,
			Description: "Let the detected server name contribute to the service identifier calculation",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ServiceIdContributor).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"web_service_name": {
			Type:        hcl.TypeList,
			Description: "Let the detected web service name contribute to the service identifier calculation",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ServiceIdContributor).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"web_service_namespace": {
			Type:        hcl.TypeList,
			Description: "Let the detected web service namespace contribute to the service identifier calculation",
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
		"application_id":                me.ApplicationID,
		"context_root":                  me.ContextRoot,
		"detect_as_web_request_service": me.DetectAsWebRequestService,
		"server_name":                   me.ServerName,
		"web_service_name":              me.WebServiceName,
		"web_service_namespace":         me.WebServiceNamespace,
	})
}

func (me *IdContributorsType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"application_id":                &me.ApplicationID,
		"context_root":                  &me.ContextRoot,
		"detect_as_web_request_service": &me.DetectAsWebRequestService,
		"server_name":                   &me.ServerName,
		"web_service_name":              &me.WebServiceName,
		"web_service_namespace":         &me.WebServiceNamespace,
	})
}
