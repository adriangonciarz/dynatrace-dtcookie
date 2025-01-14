package managementzones

import (
	"encoding/json"
	"sort"

	"github.com/dtcookie/dynatrace/api/config/entityruleengine"
	"github.com/dtcookie/hcl"
)

// Rule The rule of the management zone usage. It defines how the management zone applies.
// Each rule is evaluated independently of all other rules.
type Rule struct {
	Conditions       []*entityruleengine.Condition `json:"conditions"`                 // A list of matching rules for the management zone.  The management zone applies only if **all** conditions are fulfilled.
	Enabled          bool                          `json:"enabled"`                    // The rule is enabled (`true`) or disabled (`false`).
	PropagationTypes []PropagationType             `json:"propagationTypes,omitempty"` // How to apply the management zone to underlying entities:  * `SERVICE_TO_HOST_LIKE`: Apply to underlying hosts of matching services\n   - `SERVICE_TO_PROCESS_GROUP_LIKE`: Apply to underlying process groups of matching services\n   - `PROCESS_GROUP_TO_HOST`: Apply to underlying hosts of matching process groups\n   - `PROCESS_GROUP_TO_SERVICE`: Apply to all services provided by matching process groups\n   - `HOST_TO_PROCESS_GROUP_INSTANCE`: Apply to processes running on matching hosts\n   - `CUSTOM_DEVICE_GROUP_TO_CUSTOM_DEVICE`: Apply to custom devices in matching custom device groups\n   - `AZURE_TO_PG`: Apply to process groups connected to matching Azure entities\n   - `AZURE_TO_SERVICE`: Apply to services provided by matching Azure entities.
	Type             RuleType                      `json:"type"`                       // The type of Dynatrace entities the management zone can be applied to.
	Unknowns         map[string]json.RawMessage    `json:"-"`
}

func (mzr *Rule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "The type of Dynatrace entities the management zone can be applied to",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "The rule is enabled (`true`) or disabled (`false`)",
			Optional:    true,
		},
		"propagation_types": {
			Type:        hcl.TypeSet,
			Description: "How to apply the management zone to underlying entities:\n   - `SERVICE_TO_HOST_LIKE`: Apply to underlying hosts of matching services\n   - `SERVICE_TO_PROCESS_GROUP_LIKE`: Apply to underlying process groups of matching services\n   - `PROCESS_GROUP_TO_HOST`: Apply to underlying hosts of matching process groups\n   - `PROCESS_GROUP_TO_SERVICE`: Apply to all services provided by matching process groups\n   - `HOST_TO_PROCESS_GROUP_INSTANCE`: Apply to processes running on matching hosts\n   - `CUSTOM_DEVICE_GROUP_TO_CUSTOM_DEVICE`: Apply to custom devices in matching custom device groups\n   - `AZURE_TO_PG`: Apply to process groups connected to matching Azure entities\n   - `AZURE_TO_SERVICE`: Apply to services provided by matching Azure entities",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"conditions": {
			Type:        hcl.TypeList,
			Description: "A list of matching rules for the management zone. The management zone applies only if **all** conditions are fulfilled",
			MinItems:    1,
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(entityruleengine.Condition).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (mzr *Rule) SortConditions() {
	if len(mzr.Conditions) > 0 {
		conds := []*entityruleengine.Condition{}
		condStrings := sort.StringSlice{}
		for _, entry := range mzr.Conditions {
			condBytes, _ := json.Marshal(entry)
			condStrings = append(condStrings, string(condBytes))
		}
		condStrings.Sort()
		for _, condString := range condStrings {
			cond := entityruleengine.Condition{}
			json.Unmarshal([]byte(condString), &cond)
			conds = append(conds, &cond)
		}
		mzr.Conditions = conds
	}
}

func (mzr *Rule) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(mzr.Unknowns) > 0 {
		data, err := json.Marshal(mzr.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["enabled"] = mzr.Enabled
	result["type"] = string(mzr.Type)
	if len(mzr.PropagationTypes) > 0 {
		entries := []interface{}{}
		for _, entry := range mzr.PropagationTypes {
			entries = append(entries, string(entry))
		}
		result["propagation_types"] = entries
	}
	if len(mzr.Conditions) > 0 {
		mzr.SortConditions()
		entries := []interface{}{}
		for _, entry := range mzr.Conditions {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["conditions"] = entries
	}
	return result, nil
}

func (mzr *Rule) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), mzr); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &mzr.Unknowns); err != nil {
			return err
		}
		delete(mzr.Unknowns, "type")
		delete(mzr.Unknowns, "enabled")
		delete(mzr.Unknowns, "propagation_types")
		delete(mzr.Unknowns, "conditions")
		if len(mzr.Unknowns) == 0 {
			mzr.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		mzr.Type = RuleType(value.(string))
	}
	if _, value := decoder.GetChange("enabled"); value != nil {
		mzr.Enabled = value.(bool)
	}
	if propagationTypes := decoder.GetStringSet("propagation_types"); len(propagationTypes) > 0 {
		mzr.PropagationTypes = []PropagationType{}
		for _, propagationType := range propagationTypes {
			mzr.PropagationTypes = append(mzr.PropagationTypes, PropagationType(propagationType))
		}
	}

	if result, ok := decoder.GetOk("conditions.#"); ok {
		mzr.Conditions = []*entityruleengine.Condition{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(entityruleengine.Condition)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "conditions", idx)); err != nil {
				return err
			}
			mzr.Conditions = append(mzr.Conditions, entry)
		}
	}
	return nil
}

func (mzr *Rule) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(mzr.Unknowns) > 0 {
		for k, v := range mzr.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(mzr.Enabled)
		if err != nil {
			return nil, err
		}
		m["enabled"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(mzr.Type)
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	if len(mzr.PropagationTypes) > 0 {
		rawMessage, err := json.Marshal(mzr.PropagationTypes)
		if err != nil {
			return nil, err
		}
		m["propagationTypes"] = rawMessage
	}
	if len(mzr.Conditions) > 0 {
		rawMessage, err := json.Marshal(mzr.Conditions)
		if err != nil {
			return nil, err
		}
		m["conditions"] = rawMessage
	}

	return json.Marshal(m)
}

func (mzr *Rule) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["enabled"]; found {
		if err := json.Unmarshal(v, &mzr.Enabled); err != nil {
			return err
		}
	}
	if v, found := m["type"]; found {
		if err := json.Unmarshal(v, &mzr.Type); err != nil {
			return err
		}
	}
	if v, found := m["propagationTypes"]; found {
		if err := json.Unmarshal(v, &mzr.PropagationTypes); err != nil {
			return err
		}
	}
	if v, found := m["conditions"]; found {
		if err := json.Unmarshal(v, &mzr.Conditions); err != nil {
			return err
		}

	}
	delete(m, "enabled")
	delete(m, "type")
	delete(m, "propagationTypes")
	delete(m, "conditions")
	if len(m) > 0 {
		mzr.Unknowns = m
	}
	return nil
}

// PropagationType has no documentation
type PropagationType string

// PropagationTypes offers the known enum values
var PropagationTypes = struct {
	AzureToPg                       PropagationType
	AzureToService                  PropagationType
	CustomDeviceGroupToCustomDevice PropagationType
	HostToProcessGroupInstance      PropagationType
	ProcessGroupToHost              PropagationType
	ProcessGroupToService           PropagationType
	ServiceToHostLike               PropagationType
	ServiceToProcessGroupLike       PropagationType
}{
	"AZURE_TO_PG",
	"AZURE_TO_SERVICE",
	"CUSTOM_DEVICE_GROUP_TO_CUSTOM_DEVICE",
	"HOST_TO_PROCESS_GROUP_INSTANCE",
	"PROCESS_GROUP_TO_HOST",
	"PROCESS_GROUP_TO_SERVICE",
	"SERVICE_TO_HOST_LIKE",
	"SERVICE_TO_PROCESS_GROUP_LIKE",
}

// RuleType The type of Dynatrace entities the management zone can be applied to.
type RuleType string

// RuleTypes offers the known enum values
var RuleTypes = struct {
	AppMonServer                 RuleType
	AppMonSystemProfile          RuleType
	AWSAccount                   RuleType
	AWSApplicationLoadBalancer   RuleType
	AWSAutoScalingGroup          RuleType
	AWSClassicLoadBalancer       RuleType
	AWSNetworkLoadBalancer       RuleType
	AWSRelationalDatabaseService RuleType
	Azure                        RuleType
	BrowserMonitor               RuleType
	CloudApplication             RuleType
	CloudApplicationNamespace    RuleType
	CloudFoundryFoundation       RuleType
	CustomApplication            RuleType
	CustomDevice                 RuleType
	CustomDeviceGroup            RuleType
	DataCenterService            RuleType
	EnterpriseApplication        RuleType
	ESXIHost                     RuleType
	ExternalMonitor              RuleType
	Host                         RuleType
	HostGroup                    RuleType
	HTTPMonitor                  RuleType
	KubernetesCluster            RuleType
	MobileApplication            RuleType
	OpenStackAccount             RuleType
	ProcessGroup                 RuleType
	Service                      RuleType
	WebApplication               RuleType
}{
	"APPMON_SERVER",
	"APPMON_SYSTEM_PROFILE",
	"AWS_ACCOUNT",
	"AWS_APPLICATION_LOAD_BALANCER",
	"AWS_AUTO_SCALING_GROUP",
	"AWS_CLASSIC_LOAD_BALANCER",
	"AWS_NETWORK_LOAD_BALANCER",
	"AWS_RELATIONAL_DATABASE_SERVICE",
	"AZURE",
	"BROWSER_MONITOR",
	"CLOUD_APPLICATION",
	"CLOUD_APPLICATION_NAMESPACE",
	"CLOUD_FOUNDRY_FOUNDATION",
	"CUSTOM_APPLICATION",
	"CUSTOM_DEVICE",
	"CUSTOM_DEVICE_GROUP",
	"DATA_CENTER_SERVICE",
	"ENTERPRISE_APPLICATION",
	"ESXI_HOST",
	"EXTERNAL_MONITOR",
	"HOST",
	"HOST_GROUP",
	"HTTP_MONITOR",
	"KUBERNETES_CLUSTER",
	"MOBILE_APPLICATION",
	"OPENSTACK_ACCOUNT",
	"PROCESS_GROUP",
	"SERVICE",
	"WEB_APPLICATION",
}
