package frequentissues

import "github.com/dtcookie/hcl"

type Settings struct {
	DetectFrequentIssuesInApplications            bool `json:"detectFrequentIssuesInApplications"`            // Detect frequent issues within applications
	DetectFrequentIssuesInInfrastructure          bool `json:"detectFrequentIssuesInInfrastructure"`          // Detect frequent issues within infrastructure
	DetectFrequentIssuesInTransactionsAndServices bool `json:"detectFrequentIssuesInTransactionsAndServices"` // Detect frequent issues within transactions and services
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"detect_frequent_issues_in_applications": {
			Type:        hcl.TypeBool,
			Description: "Detect frequent issues within applications",
			Optional:    true,
		},
		"detect_frequent_issues_in_infrastructure": {
			Type:        hcl.TypeBool,
			Description: "Detect frequent issues within infrastructure",
			Optional:    true,
		},
		"detect_frequent_issues_in_transactions_and_services": {
			Type:        hcl.TypeBool,
			Description: "Detect frequent issues within transactions and services",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"detect_frequent_issues_in_applications":              me.DetectFrequentIssuesInApplications,
		"detect_frequent_issues_in_infrastructure":            me.DetectFrequentIssuesInInfrastructure,
		"detect_frequent_issues_in_transactions_and_services": me.DetectFrequentIssuesInTransactionsAndServices,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"detect_frequent_issues_in_applications":              &me.DetectFrequentIssuesInApplications,
		"detect_frequent_issues_in_infrastructure":            &me.DetectFrequentIssuesInInfrastructure,
		"detect_frequent_issues_in_transactions_and_services": &me.DetectFrequentIssuesInTransactionsAndServices,
	})
}
