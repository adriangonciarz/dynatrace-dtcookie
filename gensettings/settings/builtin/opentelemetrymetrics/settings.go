package opentelemetrymetrics

import "github.com/dtcookie/hcl"

type Settings struct {
	AdditionalAttributes                   AdditionalAttributeItems `json:"additionalAttributes"`                   // The attributes configured below are only added as dimensions to ingested OTLP metrics when the setting *Add configured resource and scope attributes as dimensions* above is enabled.\n\n**Notes:**\n\n* Modifying this setting (renaming, disabling or removing attributes) will cause the metric to change. This may have an impact on existing dashboards, events and alerts that make use of these dimensions. In this case, they will need to be updated manually.\n\n* Dynatrace does not recommend changing/removing the attributes starting with \"dt.\". Dynatrace leverages these attributes to [Enrich metrics](https://www.dynatrace.com/support/help/extend-dynatrace/extend-metrics/reference/enrich-metrics).
	AdditionalAttributesToDimensionEnabled bool                     `json:"additionalAttributesToDimensionEnabled"` // When enabled, the attributes defined in the list below will be automatically added as dimensions to ingested OTLP metrics if they are present in the OpenTelemetry resource or in the instrumentation scope.\n\n**Note:** Modifying this setting will cause the metric to change. This may have an impact on existing dashboards, events and alerts that make use of these dimensions. In this case, they will need to be updated manually.
	MeterNameToDimensionEnabled            bool                     `json:"meterNameToDimensionEnabled"`            // When enabled, the Meter name (also referred to as InstrumentationScope or InstrumentationLibrary in OpenTelemetry SDKs) and version will be automatically added as dimensions (`otel.scope.name` and `otel.scope.version`) to ingested OTLP metrics.\n\n**Note:** Modifying this setting will cause the metric to change. This may have an impact on existing dashboards, events and alerts that make use of these dimensions. In this case, they will need to be updated manually.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"additional_attributes": {
			Type:        hcl.TypeList,
			Description: "The attributes configured below are only added as dimensions to ingested OTLP metrics when the setting *Add configured resource and scope attributes as dimensions* above is enabled.\n\n**Notes:**\n\n* Modifying this setting (renaming, disabling or removing attributes) will cause the metric to change. This may have an impact on existing dashboards, events and alerts that make use of these dimensions. In this case, they will need to be updated manually.\n\n* Dynatrace does not recommend changing/removing the attributes starting with \"dt.\". Dynatrace leverages these attributes to [Enrich metrics](https://www.dynatrace.com/support/help/extend-dynatrace/extend-metrics/reference/enrich-metrics).",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(AdditionalAttributeItems).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"additional_attributes_to_dimension_enabled": {
			Type:        hcl.TypeBool,
			Description: "When enabled, the attributes defined in the list below will be automatically added as dimensions to ingested OTLP metrics if they are present in the OpenTelemetry resource or in the instrumentation scope.\n\n**Note:** Modifying this setting will cause the metric to change. This may have an impact on existing dashboards, events and alerts that make use of these dimensions. In this case, they will need to be updated manually.",
			Optional:    true,
		},
		"meter_name_to_dimension_enabled": {
			Type:        hcl.TypeBool,
			Description: "When enabled, the Meter name (also referred to as InstrumentationScope or InstrumentationLibrary in OpenTelemetry SDKs) and version will be automatically added as dimensions (`otel.scope.name` and `otel.scope.version`) to ingested OTLP metrics.\n\n**Note:** Modifying this setting will cause the metric to change. This may have an impact on existing dashboards, events and alerts that make use of these dimensions. In this case, they will need to be updated manually.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"additional_attributes":                      me.AdditionalAttributes,
		"additional_attributes_to_dimension_enabled": me.AdditionalAttributesToDimensionEnabled,
		"meter_name_to_dimension_enabled":            me.MeterNameToDimensionEnabled,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"additional_attributes":                      &me.AdditionalAttributes,
		"additional_attributes_to_dimension_enabled": &me.AdditionalAttributesToDimensionEnabled,
		"meter_name_to_dimension_enabled":            &me.MeterNameToDimensionEnabled,
	})
}
