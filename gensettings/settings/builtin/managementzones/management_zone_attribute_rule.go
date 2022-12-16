package managementzones

import "github.com/dtcookie/hcl"

type ManagementZoneAttributeRule struct {
	AzureToPGPropagation                       bool                 `json:"azureToPGPropagation"`      // Apply to process groups connected to matching Azure entities
	AzureToServicePropagation                  bool                 `json:"azureToServicePropagation"` // Apply to services provided by matching Azure entities
	Conditions                                 AttributeConditions  `json:"conditions"`
	CustomDeviceGroupToCustomDevicePropagation bool                 `json:"customDeviceGroupToCustomDevicePropagation"` // Apply to custom devices in a custom device group
	EntityType                                 ManagementZoneMeType `json:"entityType"`                                 // Rule applies to
	HostToPGPropagation                        bool                 `json:"hostToPGPropagation"`                        // Apply to processes running on matching hosts
	PGToHostPropagation                        bool                 `json:"pgToHostPropagation"`                        // Apply to underlying hosts of matching process groups
	PGToServicePropagation                     bool                 `json:"pgToServicePropagation"`                     // Apply to all services provided by the process groups
	ServiceToHostPropagation                   bool                 `json:"serviceToHostPropagation"`                   // Apply to underlying hosts of matching services
	ServiceToPGPropagation                     bool                 `json:"serviceToPGPropagation"`                     // Apply to underlying process groups of matching services
}

func (me *ManagementZoneAttributeRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"azure_to_pgpropagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to process groups connected to matching Azure entities",
			Optional:    true,
		},
		"azure_to_service_propagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to services provided by matching Azure entities",
			Optional:    true,
		},
		"conditions": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(AttributeConditions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"custom_device_group_to_custom_device_propagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to custom devices in a custom device group",
			Optional:    true,
		},
		"entity_type": {
			Type:        hcl.TypeString,
			Description: "Rule applies to",
			Required:    true,
		},
		"host_to_pgpropagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to processes running on matching hosts",
			Optional:    true,
		},
		"pg_to_host_propagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to underlying hosts of matching process groups",
			Optional:    true,
		},
		"pg_to_service_propagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to all services provided by the process groups",
			Optional:    true,
		},
		"service_to_host_propagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to underlying hosts of matching services",
			Optional:    true,
		},
		"service_to_pgpropagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to underlying process groups of matching services",
			Optional:    true,
		},
	}
}

func (me *ManagementZoneAttributeRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"azure_to_pgpropagation":                           me.AzureToPGPropagation,
		"azure_to_service_propagation":                     me.AzureToServicePropagation,
		"conditions":                                       me.Conditions,
		"custom_device_group_to_custom_device_propagation": me.CustomDeviceGroupToCustomDevicePropagation,
		"entity_type":                                      me.EntityType,
		"host_to_pgpropagation":                            me.HostToPGPropagation,
		"pg_to_host_propagation":                           me.PGToHostPropagation,
		"pg_to_service_propagation":                        me.PGToServicePropagation,
		"service_to_host_propagation":                      me.ServiceToHostPropagation,
		"service_to_pgpropagation":                         me.ServiceToPGPropagation,
	})
}

func (me *ManagementZoneAttributeRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"azure_to_pgpropagation":                           &me.AzureToPGPropagation,
		"azure_to_service_propagation":                     &me.AzureToServicePropagation,
		"conditions":                                       &me.Conditions,
		"custom_device_group_to_custom_device_propagation": &me.CustomDeviceGroupToCustomDevicePropagation,
		"entity_type":                                      &me.EntityType,
		"host_to_pgpropagation":                            &me.HostToPGPropagation,
		"pg_to_host_propagation":                           &me.PGToHostPropagation,
		"pg_to_service_propagation":                        &me.PGToServicePropagation,
		"service_to_host_propagation":                      &me.ServiceToHostPropagation,
		"service_to_pgpropagation":                         &me.ServiceToPGPropagation,
	})
}
