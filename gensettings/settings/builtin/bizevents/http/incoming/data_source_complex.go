package incoming

import "github.com/dtcookie/hcl"

type DataSourceComplex struct {
	DataSource DataSourceEnum `json:"dataSource"` // Data source
	Path       string         `json:"path"`       // [See our documentation](https://dt-url.net/ei034bx)
}

func (me *DataSourceComplex) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"data_source": {
			Type:        hcl.TypeString,
			Description: "Data source",
			Required:    true,
		},
		"path": {
			Type:        hcl.TypeString,
			Description: "[See our documentation](https://dt-url.net/ei034bx)",
			Required:    true,
		},
	}
}

func (me *DataSourceComplex) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"data_source": me.DataSource,
		"path":        me.Path,
	})
}

func (me *DataSourceComplex) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"data_source": &me.DataSource,
		"path":        &me.Path,
	})
}
