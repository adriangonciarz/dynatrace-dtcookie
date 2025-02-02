package autotags

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// EntitySelectorBasedRule The entity-selector-based rule for auto tag usage. It allows tagging entities via an entity selector
type EntitySelectorBasedRule struct {
	Enabled       *bool                      `json:"enabled"`        // The rule is enabled (`true`) or disabled (`false`).
	Selector      string                     `json:"entitySelector"` // The entity selector string, by which the entities are selected
	ValueFormat   *string                    `json:"valueFormat"`    // The value of the entity-selector-based auto-tag. If specified, the tag is used in the `name:valueFormat` format. \n\nFor example, you can extend the `Infrastructure` tag to `Infrastructure:Windows` and `Infrastructure:Linux`
	Normalization *string                    `json:"normalization"`  // Changes applied to the value after applying the value format. Default is LEAVE_TEXT_AS_IS
	Unknowns      map[string]json.RawMessage `json:"-"`
}

func (me *EntitySelectorBasedRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "The rule is enabled (`true`) or disabled (`false`)",
			Optional:    true,
		},
		"selector": {
			Type:        hcl.TypeString,
			Description: "The entity selector string, by which the entities are selected",
			Optional:    true,
		},
		"value_format": {
			Type:        hcl.TypeString,
			Description: "The value of the entity-selector-based auto-tag. If specified, the tag is used in the `name:valueFormat` format. \n\nFor example, you can extend the `Infrastructure` tag to `Infrastructure:Windows` and `Infrastructure:Linux`",
			Optional:    true,
		},
		"normalization": {
			Type:        hcl.TypeString,
			Description: "Changes applied to the value after applying the value format. Possible values are `LEAVE_TEXT_AS_IS`, `TO_LOWER_CASE` and `TO_UPPER_CASE`. Default is `LEAVE_TEXT_AS_IS`",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *EntitySelectorBasedRule) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["selector"] = me.Selector
	result["enabled"] = opt.Bool(me.Enabled)
	if me.ValueFormat != nil {
		result["value_format"] = *me.ValueFormat
	}
	if me.Normalization != nil {
		result["normalization"] = *me.Normalization
	}
	return result, nil
}

func (me *EntitySelectorBasedRule) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
	}
	if value, ok := decoder.GetOk("selector"); ok {
		me.Selector = value.(string)
	}
	if _, value := decoder.GetChange("enabled"); value != nil {
		me.Enabled = opt.NewBool(value.(bool))
	}
	if value, ok := decoder.GetOk("value_format"); ok {
		me.ValueFormat = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("normalization"); ok {
		me.Normalization = opt.NewString(value.(string))
	} else if me.Unknowns != nil {
		if value, ok := me.Unknowns["normalization"]; ok {
			normalization := ""
			json.Unmarshal(value, &normalization)
			me.Normalization = &normalization
		}
	}
	if me.Unknowns != nil {
		delete(me.Unknowns, "selector")
		delete(me.Unknowns, "enabled")
		delete(me.Unknowns, "value_format")
		delete(me.Unknowns, "normalization")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	return nil
}

func (me *EntitySelectorBasedRule) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(me.Selector)
		if err != nil {
			return nil, err
		}
		m["entitySelector"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(opt.Bool(me.Enabled))
		if err != nil {
			return nil, err
		}
		m["enabled"] = rawMessage
	}
	if me.ValueFormat != nil {
		rawMessage, err := json.Marshal(me.ValueFormat)
		if err != nil {
			return nil, err
		}
		m["valueFormat"] = rawMessage
	}
	if me.Normalization != nil {
		rawMessage, err := json.Marshal(me.Normalization)
		if err != nil {
			return nil, err
		}
		m["normalization"] = rawMessage
	}
	return json.Marshal(m)
}

func (me *EntitySelectorBasedRule) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["enabled"]; found {
		if err := json.Unmarshal(v, &me.Enabled); err != nil {
			return err
		}
	}
	if v, found := m["entitySelector"]; found {
		if err := json.Unmarshal(v, &me.Selector); err != nil {
			return err
		}
	}
	if v, found := m["valueFormat"]; found {
		if err := json.Unmarshal(v, &me.ValueFormat); err != nil {
			return err
		}
	}
	if v, found := m["normalization"]; found {
		if err := json.Unmarshal(v, &me.Normalization); err != nil {
			return err
		}
	}
	delete(m, "entitySelector")
	delete(m, "enabled")
	delete(m, "valueFormat")
	delete(m, "normalization")

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
