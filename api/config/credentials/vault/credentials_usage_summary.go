package vault

import (
	"github.com/dtcookie/hcl"
)

type CredentialUsageSummary []*CredentialUsageObj

type CredentialUsageObj struct {
	MonitorType MonitorType `json:"type"`
	Count       int32       `json:"count"`
}

func (me *CredentialUsageObj) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "Type of usage, `HTTP_MONITOR` or `BROWSER_MONITOR`",
			Required:    true,
		},
		"count": {
			Type:        hcl.TypeInt,
			Description: "The number of uses",
			Required:    true,
		},
	}
}

func (me *CredentialUsageObj) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["type"] = string(me.MonitorType)
	result["count"] = int(me.Count)
	return result, nil
}

func (me *CredentialUsageObj) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("type"); ok {
		me.MonitorType = MonitorType(value.(string))
	}
	if value, ok := decoder.GetOk("count"); ok {
		me.Count = int32(value.(int))
	}
	return nil
}

type MonitorType string

var MonitorTypes = struct {
	HTTPMonitor     MonitorType
	BrowserMonitorl MonitorType
}{
	MonitorType("HTTP_MONITOR"),
	MonitorType("BROWSER_MONITOR"),
}
