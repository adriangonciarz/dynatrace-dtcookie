package browser

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

type KeyStrokesEvent struct {
	EventBase
	TextValue         string         `json:"textValue,omitempty"`  // The text to enter
	Masked            *bool          `json:"masked,omitempty"`     // Indicates whether the `textValue` is encrypted (`true`) or not (`false`)
	SimulateBlurEvent bool           `json:"simulateBlurEvent"`    // Defines whether to blur the text field when it loses focus.\nSet to `true` to trigger the blur the `textValue`
	Wait              *WaitCondition `json:"wait,omitempty"`       // The wait condition for the event—defines how long Dynatrace should wait before the next action is executed
	Validate          Validations    `json:"validate,omitempty"`   // The validation rule for the event—helps you verify that your browser monitor loads the expected page content or page element
	Target            *Target        `json:"target,omitempty"`     // The tab on which the page should open
	Credential        *Credential    `json:"credential,omitempty"` // Credentials for this event
}

type Credential struct {
	ID    string `json:"id"`
	Field string `json:"field"`
}

func (me *Credential) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"vault_id": {
			Type:        hcl.TypeString,
			Description: "The ID of the credential within the Credentials Vault",
			Required:    true,
		},
		"field": {
			Type:        hcl.TypeString,
			Description: "Either `username` or `password`",
			Required:    true,
		},
	}
}

func (me *Credential) MarshalHCL() (map[string]interface{}, error) {
	return map[string]interface{}{
		"vault_id": me.ID,
		"field":    me.Field,
	}, nil
}

func (me *Credential) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("vault_id", &me.ID); err != nil {
		return err
	}
	if err := decoder.Decode("field", &me.Field); err != nil {
		return err
	}
	return nil
}

func (me *KeyStrokesEvent) GetType() EventType {
	return EventTypes.KeyStrokes
}

func (me *KeyStrokesEvent) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"credential": {
			Type:        hcl.TypeList,
			Description: "Credentials for this event",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(Credential).Schema()},
		},
		"text": {
			Type:        hcl.TypeString,
			Description: "The text to enter. Must not be specified if `credentials` from the vault are being used",
			Optional:    true,
		},
		"masked": {
			Type:        hcl.TypeBool,
			Description: "Indicates whether the `textValue` is encrypted (`true`) or not (`false`). Must not be specified if `credentials` from the vault are being used",
			Optional:    true,
		},
		"simulate_blur_event": {
			Type:        hcl.TypeBool,
			Description: "Defines whether to blur the text field when it loses focus.\nSet to `true` to trigger the blur the `textValue`",
			Optional:    true,
		},
		"wait": {
			Type:        hcl.TypeList,
			Description: "The wait condition for the event—defines how long Dynatrace should wait before the next action is executed",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(WaitCondition).Schema()},
		},
		"validate": {
			Type:        hcl.TypeList,
			Description: "The validation rules for the event—helps you verify that your browser monitor loads the expected page content or page element",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(Validations).Schema()},
		},
		"target": {
			Type:        hcl.TypeList,
			Description: "The tab on which the page should open",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(Target).Schema()},
		},
	}
}

func (me *KeyStrokesEvent) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if me.Target != nil {
		if marshalled, err := me.Target.MarshalHCL(); err == nil {
			result["target"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Credential != nil {
		if marshalled, err := me.Credential.MarshalHCL(); err == nil {
			result["credential"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Wait != nil {
		if marshalled, err := me.Wait.MarshalHCL(); err == nil {
			result["wait"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if len(me.Validate) > 0 {
		if marshalled, err := me.Validate.MarshalHCL(); err == nil {
			result["validate"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Credential == nil {
		if me.Masked != nil && *me.Masked {
			result["masked"] = me.Masked
		} else {
			result["masked"] = false
		}
	} else {
		result["masked"] = false
	}
	result["text"] = me.TextValue
	result["simulate_blur_event"] = me.SimulateBlurEvent
	return result, nil
}

func (me *KeyStrokesEvent) UnmarshalHCL(decoder hcl.Decoder) error {
	me.Type = EventTypes.Tap
	if err := decoder.Decode("text", &me.TextValue); err != nil {
		return err
	}
	if err := decoder.Decode("masked", &me.Masked); err != nil {
		return err
	}
	if err := decoder.Decode("simulate_blur_event", &me.SimulateBlurEvent); err != nil {
		return err
	}
	if err := decoder.Decode("wait", &me.Wait); err != nil {
		return err
	}
	if err := decoder.Decode("credential", &me.Credential); err != nil {
		return err
	}
	if me.Credential == nil && me.Masked == nil {
		me.Masked = opt.NewBool(false)
	}
	if err := decoder.Decode("validate", &me.Validate); err != nil {
		return err
	}
	if err := decoder.Decode("target", &me.Target); err != nil {
		return err
	}
	return nil
}
