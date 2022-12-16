package advanceddetectionrule

import "github.com/dtcookie/hcl"

type ProcessInstanceExtraction struct {
	Delimiter *Delimiter `json:"delimiter"` // Optionally delimit this property between *From* and *To*.
	Property  string     `json:"property"`
}

func (me *ProcessInstanceExtraction) Schema() map[string]*hcl.Schema {
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
			Optional:    true,
		},
	}
}

func (me *ProcessInstanceExtraction) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"delimiter": me.Delimiter,
		"property":  me.Property,
	})
}

func (me *ProcessInstanceExtraction) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"delimiter": &me.Delimiter,
		"property":  &me.Property,
	})
}
