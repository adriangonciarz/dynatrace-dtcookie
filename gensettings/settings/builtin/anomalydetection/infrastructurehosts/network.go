package infrastructurehosts

import "github.com/dtcookie/hcl"

type Network struct {
	HighNetworkDetection               *HighNetworkDetection               `json:"highNetworkDetection"`
	NetworkDroppedPacketsDetection     *NetworkDroppedPacketsDetection     `json:"networkDroppedPacketsDetection"`
	NetworkErrorsDetection             *NetworkErrorsDetection             `json:"networkErrorsDetection"`
	NetworkHighRetransmissionDetection *NetworkHighRetransmissionDetection `json:"networkHighRetransmissionDetection"`
	NetworkTcpProblemsDetection        *NetworkTcpProblemsDetection        `json:"networkTcpProblemsDetection"`
}

func (me *Network) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"high_network_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(HighNetworkDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"network_dropped_packets_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(NetworkDroppedPacketsDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"network_errors_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(NetworkErrorsDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"network_high_retransmission_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(NetworkHighRetransmissionDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"network_tcp_problems_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(NetworkTcpProblemsDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Network) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"high_network_detection":                me.HighNetworkDetection,
		"network_dropped_packets_detection":     me.NetworkDroppedPacketsDetection,
		"network_errors_detection":              me.NetworkErrorsDetection,
		"network_high_retransmission_detection": me.NetworkHighRetransmissionDetection,
		"network_tcp_problems_detection":        me.NetworkTcpProblemsDetection,
	})
}

func (me *Network) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"high_network_detection":                &me.HighNetworkDetection,
		"network_dropped_packets_detection":     &me.NetworkDroppedPacketsDetection,
		"network_errors_detection":              &me.NetworkErrorsDetection,
		"network_high_retransmission_detection": &me.NetworkHighRetransmissionDetection,
		"network_tcp_problems_detection":        &me.NetworkTcpProblemsDetection,
	})
}
