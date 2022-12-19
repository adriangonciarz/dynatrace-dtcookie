package iam

import "github.com/dtcookie/hcl"

type Policy struct {
	// UUID           string   `json:"uuid,omitempty"`
	Name           string   `json:"name"`
	Tags           []string `json:"tags,omitempty"`
	Description    string   `json:"description,omitempty"`
	StatementQuery string   `json:"statementQuery"`
}

func (me *Policy) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:     hcl.TypeString,
			Required: true,
		},
		"description": {
			Type:     hcl.TypeString,
			Optional: true,
		},
		"statement_query": {
			Type:     hcl.TypeString,
			Required: true,
		},
	}
}

func (me *Policy) MarshalHCL() (map[string]interface{}, error) {
	return map[string]interface{}{
		"name":            me.Name,
		"description":     me.Description,
		"statement_query": me.StatementQuery,
	}, nil
}

func (me *Policy) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"name":            &me.Name,
		"description":     &me.Description,
		"statement_query": &me.StatementQuery,
	})
}
