package frequentissues

import (
	"github.com/dtcookie/hcl"
)

type FrequentIssues struct {
	DetectApps  bool `json:"detectFrequentIssuesInApplications"`            // Detect frequent issues within applications, enabled (`true`) or disabled (`false`)
	DetectTxn   bool `json:"detectFrequentIssuesInTransactionsAndServices"` // Detect frequent issues within transactions and services, enabled (`true`) or disabled (`false`)
	DetectInfra bool `json:"detectFrequentIssuesInInfrastructure"`          // Detect frequent issues within infrastructure, enabled (`true`) or disabled (`false`)
}

func (me *FrequentIssues) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"detect_apps": {
			Type:        hcl.TypeBool,
			Required:    true,
			Description: "Detect frequent issues within applications, enabled (`true`) or disabled (`false`)",
		},
		"detect_txn": {
			Type:        hcl.TypeBool,
			Required:    true,
			Description: "Detect frequent issues within transactions and services, enabled (`true`) or disabled (`false`)",
		},
		"detect_infra": {
			Type:        hcl.TypeBool,
			Required:    true,
			Description: "Detect frequent issues within infrastructure, enabled (`true`) or disabled (`false`)",
		},
	}
}

func (me *FrequentIssues) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"detect_apps":  me.DetectApps,
		"detect_txn":   me.DetectTxn,
		"detect_infra": me.DetectInfra,
	})
}

func (me *FrequentIssues) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"detect_apps":  &me.DetectApps,
		"detect_txn":   &me.DetectTxn,
		"detect_infra": &me.DetectInfra,
	})
}
