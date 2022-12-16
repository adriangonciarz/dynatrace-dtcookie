package notifications

import (
	"github.com/dtcookie/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type WebHookNotificationHeaders []*WebHookNotificationHeader

func (me *WebHookNotificationHeaders) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"header": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new(WebHookNotificationHeader).Schema()},
		},
	}
}

func (me WebHookNotificationHeaders) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me) > 0 {
		entries := []interface{}{}
		for _, entry := range me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["header"] = entries
	}
	return result, nil
}

func (me *WebHookNotificationHeaders) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("header"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new(WebHookNotificationHeader)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "header", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

type WebHookNotificationHeader struct {
	Name        string `json:"name"`        // The name of the HTTP header.
	Secret      bool   `json:"secret"`      // Secret HTTP header value
	SecretValue string `json:"secretValue"` // The secret value of the HTTP header. May contain an empty value.
	Value       string `json:"value"`       // The value of the HTTP header. May contain an empty value.
}

func (me *WebHookNotificationHeader) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the HTTP header.",
			Required:    true,
		},
		"secret": {
			Type:        hcl.TypeBool,
			Description: "Secret HTTP header value",
			Optional:    true,
		},
		"secret_value": {
			Type:        hcl.TypeString,
			Description: "The secret value of the HTTP header. May contain an empty value.",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value of the HTTP header. May contain an empty value.",
			Required:    true,
		},
	}
}

func (me *WebHookNotificationHeader) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"name":         me.Name,
		"secret":       me.Secret,
		"secret_value": me.SecretValue,
		"value":        me.Value,
	})
}

func (me *WebHookNotificationHeader) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"name":         &me.Name,
		"secret":       &me.Secret,
		"secret_value": &me.SecretValue,
		"value":        &me.Value,
	})
}
