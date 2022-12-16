package advanceddetectionrule

import "github.com/dtcookie/hcl"

type ProcessGroupExtraction struct {
	Delimiter      *Delimiter `json:"delimiter"` // Optionally delimit this property between *From* and *To*.
	Property       string     `json:"property"`
	StandaloneRule bool       `json:"standaloneRule"` // If this option is selected, the default Dynatrace behavior is disabled for these detected processes. Only this rule is used to separate the process group.\n\nIf this option is not selected, this rule contributes to the default Dynatrace process group detection. \n\n[See our help page for examples.](https://dt-url.net/1722wrz)
}

func (me *ProcessGroupExtraction) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"delimiter": {
			Type:        hcl.TypeList,
			Description: "Optionally delimit this property between *From* and *To*.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Delimiter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"property": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"standalone_rule": {
			Type:        hcl.TypeBool,
			Description: "If this option is selected, the default Dynatrace behavior is disabled for these detected processes. Only this rule is used to separate the process group.\n\nIf this option is not selected, this rule contributes to the default Dynatrace process group detection. \n\n[See our help page for examples.](https://dt-url.net/1722wrz)",
			Optional:    true,
		},
	}
}

func (me *ProcessGroupExtraction) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"delimiter":       me.Delimiter,
		"property":        me.Property,
		"standalone_rule": me.StandaloneRule,
	})
}

func (me *ProcessGroupExtraction) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"delimiter":       &me.Delimiter,
		"property":        &me.Property,
		"standalone_rule": &me.StandaloneRule,
	})
}
