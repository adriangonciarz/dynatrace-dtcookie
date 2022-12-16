package usersessionexportsettingsv2

import "github.com/dtcookie/hcl"

type Settings struct {
	Authentication     *Authentication     `json:"authentication"`     // Dynatrace can automatically send bulk data to Elasticsearch. You can use an SSL certificate, basic authentication or OAuth 2.0 to secure access. For Dynatrace SaaS installations, the Elasticsearch instance must be publicly reachable by the Dynatrace Cluster.
	EndpointDefinition *EndpointDefinition `json:"endpointDefinition"` // Dynatrace will send JSON data periodically to this endpoint. You can also pause and disable an endpoint to stop the data stream from Dynatrace to your endpoint.
	ExportBehavior     *ExportBehavior     `json:"exportBehavior"`     // Define the scope of your export by using a specific management zone. You can also disable UI notifications for failing exports, or add special settings provided by Dynatrace support for troubleshooting.
	SendDirect         *SendDirect         `json:"sendDirect"`         // Activate this if you want to export user session data to your own Elasticsearch cluster. If you run Elasticsearch 7, make sure to enter <em>_doc</em> as the type. For Elasticsearch 8 omit the type. If you really want to use a type, then you have to add <em>include_type_name=true</em> when creating your Elasticsearch index. See more information in the Dynatrace help.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"authentication": {
			Type:        hcl.TypeList,
			Description: "Dynatrace can automatically send bulk data to Elasticsearch. You can use an SSL certificate, basic authentication or OAuth 2.0 to secure access. For Dynatrace SaaS installations, the Elasticsearch instance must be publicly reachable by the Dynatrace Cluster.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Authentication).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"endpoint_definition": {
			Type:        hcl.TypeList,
			Description: "Dynatrace will send JSON data periodically to this endpoint. You can also pause and disable an endpoint to stop the data stream from Dynatrace to your endpoint.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EndpointDefinition).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"export_behavior": {
			Type:        hcl.TypeList,
			Description: "Define the scope of your export by using a specific management zone. You can also disable UI notifications for failing exports, or add special settings provided by Dynatrace support for troubleshooting.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ExportBehavior).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"send_direct": {
			Type:        hcl.TypeList,
			Description: "Activate this if you want to export user session data to your own Elasticsearch cluster. If you run Elasticsearch 7, make sure to enter <em>_doc</em> as the type. For Elasticsearch 8 omit the type. If you really want to use a type, then you have to add <em>include_type_name=true</em> when creating your Elasticsearch index. See more information in the Dynatrace help.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SendDirect).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"authentication":      me.Authentication,
		"endpoint_definition": me.EndpointDefinition,
		"export_behavior":     me.ExportBehavior,
		"send_direct":         me.SendDirect,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"authentication":      &me.Authentication,
		"endpoint_definition": &me.EndpointDefinition,
		"export_behavior":     &me.ExportBehavior,
		"send_direct":         &me.SendDirect,
	})
}
