package notifications

import (
	"fmt"
	"sort"
	"strings"

	"hash/crc32"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

type HTTPHeaders []*HTTPHeader

func (me *HTTPHeaders) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"header": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "An additional HTTP Header to include when sending requests",
			Elem: &hcl.Resource{
				Schema: new(HTTPHeader).Schema(),
			},
		},
	}
}

func (me HTTPHeaders) Sort() {
	sort.Sort(me)
}

func (me HTTPHeaders) Len() int {
	return len(me)
}

func (me HTTPHeaders) Less(i, j int) bool {
	return strings.Compare(me[i].Name, me[j].Name) == -1
}

func (me HTTPHeaders) Swap(i, j int) {
	tmp := me[i]
	me[j] = me[i]
	me[i] = tmp
}

func (me HTTPHeaders) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("header", me)
}

func hashString(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func (me *HTTPHeaders) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("header"); ok {
		resourceSet := value.(hcl.Set)
		for _, m := range resourceSet.List() {
			hash := hashString(m.(map[string]interface{})["name"].(string))
			var header HTTPHeader
			if err := header.UnmarshalHCL(hcl.NewDecoder(decoder, fmt.Sprintf("header.%d", hash))); err != nil {
				return err
			}
			*me = append(*me, &header)
		}
	}
	return nil
}

type HTTPHeader struct {
	Name        string  `json:"name"`                  // The name of the HTTP header
	Secret      bool    `json:"secret"`                // The value of this HTTP header is a secret (`true`) or not (`false`).
	Value       *string `json:"value,omitempty"`       // The value of the HTTP header. May contain an empty value
	SecretValue *string `json:"secretValue,omitempty"` // The secret value of the HTTP header. May contain an empty value
}

func (me *HTTPHeader) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the HTTP header",
			Required:    true,
		},
		"secret_value": {
			Type:        hcl.TypeString,
			Description: "The value of the HTTP header as a sensitive property. May contain an empty value. `secret_value` and `value` are mutually exclusive. Only one of those two is allowed to be specified.",
			Sensitive:   true,
			Optional:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value of the HTTP header. May contain an empty value. `secret_value` and `value` are mutually exclusive. Only one of those two is allowed to be specified.",
			Optional:    true,
		},
	}
}

func (me *HTTPHeader) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["name"] = me.Name
	result["value"] = me.Value
	// Must not serialize secret values
	// REST API only offers them in scrambled form
	// result["secret_value"] = me.SecretValue

	return result, nil
}

func (me *HTTPHeader) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	if value, ok := decoder.GetOk("value"); ok {
		me.Value = opt.NewString(value.(string))
		me.Secret = false
	}
	if value, ok := decoder.GetOk("secret_value"); ok {
		me.SecretValue = opt.NewString(value.(string))
		me.Secret = true
	}
	return nil
}
