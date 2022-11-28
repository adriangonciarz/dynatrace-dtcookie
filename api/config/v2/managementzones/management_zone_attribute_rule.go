package managementzones

import "github.com/dtcookie/hcl"

// No documentation available
type ManagementZoneAttributeRule struct {
	PgToServicePropagation                     bool                  `json:"pgToServicePropagation" hcl:"pg_to_service_propagation"`                                            // Apply to all services provided by the process groups
	EntityType                                 ManagementZoneMeType  `json:"entityType" hcl:"entity_type"`                                                                      // Rule applies to
	ServiceToHostPropagation                   bool                  `json:"serviceToHostPropagation" hcl:"service_to_host_propagation"`                                        // Apply to underlying hosts of matching services
	PgToHostPropagation                        bool                  `json:"pgToHostPropagation" hcl:"pg_to_host_propagation"`                                                  // Apply to underlying hosts of matching process groups
	AzureToPGPropagation                       bool                  `json:"azureToPGPropagation" hcl:"azure_to_pgpropagation"`                                                 // Apply to process groups connected to matching Azure entities
	Conditions                                 []*AttributeCondition `json:"conditions" hcl:"conditions"`                                                                       // Conditions
	ServiceToPGPropagation                     bool                  `json:"serviceToPGPropagation" hcl:"service_to_pgpropagation"`                                             // Apply to underlying process groups of matching services
	HostToPGPropagation                        bool                  `json:"hostToPGPropagation" hcl:"host_to_pgpropagation"`                                                   // Apply to processes running on matching hosts
	CustomDeviceGroupToCustomDevicePropagation bool                  `json:"customDeviceGroupToCustomDevicePropagation" hcl:"custom_device_group_to_custom_device_propagation"` // Apply to custom devices in a custom device group
	AzureToServicePropagation                  bool                  `json:"azureToServicePropagation" hcl:"azure_to_service_propagation"`                                      // Apply to services provided by matching Azure entities
}

func (me *ManagementZoneAttributeRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"pg_to_service_propagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to all services provided by the process groups",
			Optional:    true,
		},
		"entity_type": {
			Type:        hcl.TypeString,
			Description: "Rule applies to",
			Required:    true,
		},
		"service_to_host_propagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to underlying hosts of matching services",
			Optional:    true,
		},
		"pg_to_host_propagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to underlying hosts of matching process groups",
			Optional:    true,
		},
		"azure_to_pgpropagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to process groups connected to matching Azure entities",
			Optional:    true,
		},
		"conditions": {
			Type:        hcl.TypeList,
			Description: "Conditions",
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(AttributeCondition).Schema()},
			Required:    true,
		},
		"service_to_pgpropagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to underlying process groups of matching services",
			Optional:    true,
		},
		"host_to_pgpropagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to processes running on matching hosts",
			Optional:    true,
		},
		"custom_device_group_to_custom_device_propagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to custom devices in a custom device group",
			Optional:    true,
		},
		"azure_to_service_propagation": {
			Type:        hcl.TypeBool,
			Description: "Apply to services provided by matching Azure entities",
			Optional:    true,
		},
	}
}
