package usersessionexportsettingsv2

import "github.com/dtcookie/hcl"

type SendDirect struct {
	Active    bool   `json:"active"`    // Activate
	DocType   string `json:"docType"`   // Type of documents in the Elasticsearch index
	IndexName string `json:"indexName"` // Name of the index where data is sent
}

func (me *SendDirect) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"active": {
			Type:        hcl.TypeBool,
			Description: "Activate",
			Optional:    true,
		},
		"doc_type": {
			Type:        hcl.TypeString,
			Description: "Type of documents in the Elasticsearch index",
			Optional:    true,
		},
		"index_name": {
			Type:        hcl.TypeString,
			Description: "Name of the index where data is sent",
			Required:    true,
		},
	}
}

func (me *SendDirect) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"active":     me.Active,
		"doc_type":   me.DocType,
		"index_name": me.IndexName,
	})
}

func (me *SendDirect) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"active":     &me.Active,
		"doc_type":   &me.DocType,
		"index_name": &me.IndexName,
	})
}
