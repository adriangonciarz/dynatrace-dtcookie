package txstartfilters

import "github.com/dtcookie/hcl"

type Settings struct {
	IncludedCicsTerminalTransactionIds []string `json:"includedCicsTerminalTransactionIds"` // You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	IncludedCicsTransactionIds         []string `json:"includedCicsTransactionIds"`         // You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	IncludedImsTransactionIds          []string `json:"includedImsTransactionIds"`          // You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"included_cics_terminal_transaction_ids": {
			Type:        hcl.TypeSet,
			Description: "You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"included_cics_transaction_ids": {
			Type:        hcl.TypeSet,
			Description: "You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"included_ims_transaction_ids": {
			Type:        hcl.TypeSet,
			Description: "You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"included_cics_terminal_transaction_ids": me.IncludedCicsTerminalTransactionIds,
		"included_cics_transaction_ids":          me.IncludedCicsTransactionIds,
		"included_ims_transaction_ids":           me.IncludedImsTransactionIds,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"included_cics_terminal_transaction_ids": &me.IncludedCicsTerminalTransactionIds,
		"included_cics_transaction_ids":          &me.IncludedCicsTransactionIds,
		"included_ims_transaction_ids":           &me.IncludedImsTransactionIds,
	})
}
