package alerting

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// Profile represents an Alerting Profile in Dynatrace
type Profile struct {
	Name           string        `json:"name"`                    // The name of the Alerting Profile
	ManagementZone *string       `json:"managementZone"`          // Define management zone filter for profile
	SeverityRules  SeverityRules `json:"severityRules,omitempty"` // Define severity rules for profile. A maximum of 100 severity rules is allowed.
	EventFilters   EventFilters  `json:"eventFilters,omitempty"`  // Define event filters for profile. A maximum of 100 event filters is allowed.
}

// Schema provides a map for terraform, containing all the current supported properties
func (me *Profile) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the alerting profile, displayed in the UI",
			Required:    true,
		},
		"management_zone": {
			Type:        hcl.TypeString,
			Description: "The ID of the management zone to which the alerting profile applies",
			Optional:    true,
		},
		"rules": {
			Type:        hcl.TypeList,
			Description: "A list of rules for management zone usage.  Each rule is evaluated independently of all other rules",
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(SeverityRules).Schema()},
		},
		"filters": {
			Type:        hcl.TypeList,
			Description: "The list of event filters.  For all filters that are *negated* inside of these event filters, that is all `Predefined` as well as `Custom` (Title and/or Description) ones the AND logic applies. For all *non-negated* ones the OR logic applies. Between these two groups, negated and non-negated, the AND logic applies.  If you specify both severity rule and event filter, the AND logic applies",
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(EventFilters).Schema()},
		},
	}
}

// EnsurePredictableOrder is currently an inconvenient necessity. The REST API does not guarantee to deliever the Severity Rules for an Alerting Profile in a predictable order.
// Terraform does however insist on the same order every time. Therefore the the Severity Rules are getting ordered based on their JSON representation, before any HCL code is getting produced.
func (me *Profile) EnsurePredictableOrder() {
	if len(me.SeverityRules) == 0 {
		return
	}
	conds := SeverityRules{}
	condStrings := sort.StringSlice{}
	for _, entry := range me.SeverityRules {
		condBytes, _ := json.Marshal(entry)
		condStrings = append(condStrings, string(condBytes))
	}
	condStrings.Sort()
	for _, condString := range condStrings {
		cond := SeverityRule{}
		json.Unmarshal([]byte(condString), &cond)
		conds = append(conds, &cond)
	}
	me.SeverityRules = conds
}

// MarshalHCL produces HCL structures for Terraform
func (me *Profile) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["name"] = me.Name
	if me.ManagementZone != nil {
		result["management_zone"] = me.ManagementZone
	}
	if len(me.SeverityRules) > 0 {
		me.EnsurePredictableOrder()
		marshalled, err := me.SeverityRules.MarshalHCL()
		if err != nil {
			return nil, err
		}
		result["rules"] = []interface{}{marshalled}
	}
	if len(me.EventFilters) > 0 {
		filters := append(EventFilters{}, me.EventFilters...)
		sort.Slice(filters, func(i, j int) bool {
			d1, _ := json.Marshal(filters[i])
			d2, _ := json.Marshal(filters[j])
			cmp := strings.Compare(string(d1), string(d2))
			return (cmp == -1)
		})
		marshalled, err := me.EventFilters.MarshalHCL()
		if err != nil {
			return nil, err
		}
		result["filters"] = []interface{}{marshalled}
	}
	return result, nil
}

// UnmarshalHCL decodes HCL code and fills this object
func (me *Profile) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	if value, ok := decoder.GetOk("management_zone"); ok {
		me.ManagementZone = opt.NewString(value.(string))
	}
	if _, ok := decoder.GetOk("rules.#"); ok {
		me.SeverityRules = SeverityRules{}
		if err := me.SeverityRules.UnmarshalHCL(hcl.NewDecoder(decoder, "rules", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("filters.#"); ok {
		me.EventFilters = EventFilters{}
		if err := me.EventFilters.UnmarshalHCL(hcl.NewDecoder(decoder, "filters", 0)); err != nil {
			return err
		}
	}
	return nil
}
