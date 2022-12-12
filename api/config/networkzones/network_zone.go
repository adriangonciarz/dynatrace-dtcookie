package networkzones

import (
	"github.com/dtcookie/hcl"
)

// NetworkZone TODO: documentation
type NetworkZone struct {
	ID                           *string   `json:"id,omitempty"`                           // The ID of the network zone
	Description                  *string   `json:"description,omitempty"`                  // A short description of the network zone
	AltZones                     []*string `json:"alternativeZones,omitempty"`             // A list of alternative network zones.
	NumOfOneAgentsFromOtherZones *int      `json:"numOfOneAgentsFromOtherZones,omitempty"` // The number of OneAgents from other network zones that are using ActiveGates in the network zone.
	NumOfOneAgentsUsing          *int      `json:"numOfOneAgentsUsing,omitempty"`          // The number of OneAgents that are using ActiveGates in the network zone.
	NumofConfiguredActiveGates   *int      `json:"numOfConfiguredActiveGates,omitempty"`   // The number of ActiveGates in the network zone.
	NumOfConfiguredOneAgents     *int      `json:"numOfConfiguredOneAgents,omitempty"`     // The number of OneAgents that are configured to use the network zone as primary.
}

type NetworkZones struct {
	Zones []NetworkZone `json:"networkZones"`
}

func (me *NetworkZone) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"description": {
			Type:        hcl.TypeString,
			Description: "A short description of the network zone",
			Optional:    true,
		},
		"alternative_zones": {
			Type:        hcl.TypeSet,
			Description: "A list of alternative network zones.",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"num_of_oneagents_from_other_zones": {
			Type:        hcl.TypeInt,
			Description: "The number of OneAgents from other network zones that are using ActiveGates in the network zone.",
			Optional:    true,
		},
		"num_of_oneagents_using": {
			Type:        hcl.TypeInt,
			Description: "The number of OneAgents that are using ActiveGates in the network zone.",
			Optional:    true,
		},
		"num_of_configured_activegates": {
			Type:        hcl.TypeInt,
			Description: "The number of ActiveGates in the network zone.",
			Optional:    true,
		},
		"num_of_configured_oneagents": {
			Type:        hcl.TypeInt,
			Description: "The number of OneAgents that are configured to use the network zone as primary.",
			Optional:    true,
		},
	}
}

func (me *NetworkZone) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"description":                       me.Description,
		"alternative_zones":                 me.AltZones,
		"num_of_oneagents_from_other_zones": me.NumOfOneAgentsFromOtherZones,
		"num_of_oneagents_using":            me.NumOfOneAgentsUsing,
		"num_of_configured_activegates":     me.NumofConfiguredActiveGates,
		"num_of_configured_oneagents":       me.NumOfConfiguredOneAgents,
	})
}

func (me *NetworkZone) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"description":       &me.Description,
		"alternative_zones": &me.AltZones,
	})
}
