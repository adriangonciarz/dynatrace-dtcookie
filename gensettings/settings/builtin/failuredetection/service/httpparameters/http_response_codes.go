package httpparameters

import "github.com/dtcookie/hcl"

type HttpResponseCodes struct {
	ClientSideErrors                    string `json:"clientSideErrors"`                    // HTTP response codes which indicate client side errors
	FailOnMissingResponseCodeClientSide bool   `json:"failOnMissingResponseCodeClientSide"` // Treat missing HTTP response code as client side error
	FailOnMissingResponseCodeServerSide bool   `json:"failOnMissingResponseCodeServerSide"` // Treat missing HTTP response code as server side errors
	ServerSideErrors                    string `json:"serverSideErrors"`                    // HTTP response codes which indicate an error on the server side
}

func (me *HttpResponseCodes) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"client_side_errors": {
			Type:        hcl.TypeString,
			Description: "HTTP response codes which indicate client side errors",
			Required:    true,
		},
		"fail_on_missing_response_code_client_side": {
			Type:        hcl.TypeBool,
			Description: "Treat missing HTTP response code as client side error",
			Optional:    true,
		},
		"fail_on_missing_response_code_server_side": {
			Type:        hcl.TypeBool,
			Description: "Treat missing HTTP response code as server side errors",
			Optional:    true,
		},
		"server_side_errors": {
			Type:        hcl.TypeString,
			Description: "HTTP response codes which indicate an error on the server side",
			Required:    true,
		},
	}
}

func (me *HttpResponseCodes) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"client_side_errors":                        me.ClientSideErrors,
		"fail_on_missing_response_code_client_side": me.FailOnMissingResponseCodeClientSide,
		"fail_on_missing_response_code_server_side": me.FailOnMissingResponseCodeServerSide,
		"server_side_errors":                        me.ServerSideErrors,
	})
}

func (me *HttpResponseCodes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"client_side_errors":                        &me.ClientSideErrors,
		"fail_on_missing_response_code_client_side": &me.FailOnMissingResponseCodeClientSide,
		"fail_on_missing_response_code_server_side": &me.FailOnMissingResponseCodeServerSide,
		"server_side_errors":                        &me.ServerSideErrors,
	})
}
