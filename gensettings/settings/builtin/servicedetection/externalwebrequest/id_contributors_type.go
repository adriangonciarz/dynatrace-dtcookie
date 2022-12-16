package externalwebrequest

import "github.com/dtcookie/hcl"

type IdContributorsType struct {
	ApplicationID    *ServiceIdContributor      `json:"applicationId"`    // Let the detected application identifier contribute to the service identifier calculation.
	ContextRoot      *ContextIdContributor      `json:"contextRoot"`      // Let the detected context root contribute to the service identifier calculation.\nThe context root is the first segment of the request URL after the server name. For example, in the www.dynatrace.com/support/help/dynatrace-api/ URL the context root is `support`.
	PortForServiceID bool                       `json:"portForServiceId"` // Let the port contribute to the service identifier calculation.
	PublicDomainName *PublicDomainIdContributor `json:"publicDomainName"` // Let the web request's domain name contribute to the service identifier calculation.
}

func (me *IdContributorsType) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"application_id": {
			Type:        hcl.TypeList,
			Description: "Let the detected application identifier contribute to the service identifier calculation.",
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
		"port_for_service_id": {
			Type:        hcl.TypeBool,
			Description: "Let the port contribute to the service identifier calculation.",
			Optional:    true,
		},
		"public_domain_name": {
			Type:        hcl.TypeList,
			Description: "Let the web request's domain name contribute to the service identifier calculation.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(PublicDomainIdContributor).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *IdContributorsType) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"application_id":      me.ApplicationID,
		"context_root":        me.ContextRoot,
		"port_for_service_id": me.PortForServiceID,
		"public_domain_name":  me.PublicDomainName,
	})
}

func (me *IdContributorsType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"application_id":      &me.ApplicationID,
		"context_root":        &me.ContextRoot,
		"port_for_service_id": &me.PortForServiceID,
		"public_domain_name":  &me.PublicDomainName,
	})
}
