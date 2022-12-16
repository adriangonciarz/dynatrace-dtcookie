package custommetrics

import "github.com/dtcookie/hcl"

type Settings struct {
	Dimensions []string     `json:"dimensions"` // Defines the fields that are used as dimensions. A dimension is a collection of reference information about a metric data point that is of interest to your business. Dimensions are parameters like \"browserFamily\", \"userType\", \"country\". For example, using \"userType\" as a dimension allows you to split chart data based on user types.
	Enabled    bool         `json:"enabled"`    // Enable custom metric
	Filters    Filters      `json:"filters"`    // Defines the filters for the user session. Filters apply at the moment of extracting the data and only sessions that satisfy the filtering criteria will be used to extract the custom metrics. You will not be able to modify these filters in the metric data explorer. For example, using \"userType equals REAL_USER\" will give you only data from real users, while forcing the synthetic sessions to be ignored.
	MetricKey  string       `json:"metricKey"`  // Metric key
	Value      *MetricValue `json:"value"`      // Defines the type of value to be extracted from the user session. When using **User session counter**, the number of user sessions is counted (similar to count(*) when using USQL). When using **User session field value**, the value of a user session field is extracted.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dimensions": {
			Type:        hcl.TypeList,
			Description: "Defines the fields that are used as dimensions. A dimension is a collection of reference information about a metric data point that is of interest to your business. Dimensions are parameters like \"browserFamily\", \"userType\", \"country\". For example, using \"userType\" as a dimension allows you to split chart data based on user types.",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enable custom metric",
			Optional:    true,
		},
		"filters": {
			Type:        hcl.TypeList,
			Description: "Defines the filters for the user session. Filters apply at the moment of extracting the data and only sessions that satisfy the filtering criteria will be used to extract the custom metrics. You will not be able to modify these filters in the metric data explorer. For example, using \"userType equals REAL_USER\" will give you only data from real users, while forcing the synthetic sessions to be ignored.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Filters).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"metric_key": {
			Type:        hcl.TypeString,
			Description: "Metric key",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeList,
			Description: "Defines the type of value to be extracted from the user session. When using **User session counter**, the number of user sessions is counted (similar to count(*) when using USQL). When using **User session field value**, the value of a user session field is extracted.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MetricValue).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"dimensions": me.Dimensions,
		"enabled":    me.Enabled,
		"filters":    me.Filters,
		"metric_key": me.MetricKey,
		"value":      me.Value,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"dimensions": &me.Dimensions,
		"enabled":    &me.Enabled,
		"filters":    &me.Filters,
		"metric_key": &me.MetricKey,
		"value":      &me.Value,
	})
}
