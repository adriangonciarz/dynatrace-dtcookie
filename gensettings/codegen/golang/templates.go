package golang

const GO_LIST_ALIAS_TEMPLATE = `
type {{.Name}} []*{{.Elem}}

func (me *{{.Name}}) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"{{.Item}}": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new({{.Elem}}).Schema()},
		},
	}
}

func (me {{.Name}}) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("{{.Item}}", me)
}

func (me *{{.Name}}) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("{{.Item}}", me)
}
`
const GO_SET_ALIAS_TEMPLATE = `
type {{.Name}} []*{{.Elem}}

func (me *{{.Name}}) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"{{.Item}}": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &hcl.Resource{Schema: new({{.Elem}}).Schema()},
		},
	}
}

func (me {{.Name}}) MarshalHCL() (map[string]interface{}, error) {
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
		result["{{.Item}}"] = entries
	}
	return result, nil
}

func (me *{{.Name}}) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("{{.Item}}"); ok {

		entrySet := value.(*schema.Set)

		for _, entryMap := range entrySet.List() {
			hash := entrySet.F(entryMap)
			entry := new({{.Elem}})
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "{{.Item}}", hash)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}
`

const GO_STRUCT_TEMPLATE = `
{{ range .Comments }}// {{.}}{{end}}
type {{.Name}} struct { {{with .Properties}}{{ range .}}
	{{.Name}} {{.Type}} ` + "`" + `json:"{{.JSONTag}}"` + "`" + ` {{if .Comment}}// {{.Comment}}{{end}}{{end}}{{end}}
}

func (me *{{.Name}}) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{ {{with .Properties}}{{ range .}}
		"{{.HCLTag}}":  {
			Type: {{.HCLType}},
			Description: "{{if .Comment}}{{.Comment}}{{else}}no documentation available{{end}}",
			{{if .Optional}} Optional: true {{else}} Required: true {{end}},
			{{if .Elem}} Elem: {{.Elem}},{{end}}
			{{if .MinItems}} MinItems: {{.MinItems}},{{end}}
			{{if .MaxItems}} MaxItems: {{.MaxItems}},{{end}}
		},{{end}}{{end}}
	}
}

func (me *{{.Name}}) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{ {{with .Properties}}{{ range .}}
		"{{.HCLTag}}":  me.{{.Name}},{{end}}{{end}}
	})
}

func (me *{{.Name}}) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{ {{with .Properties}}{{ range .}}
		"{{.HCLTag}}":  &me.{{.Name}},{{end}}{{end}}
	})
}
`

const GO_ENUM_TEMPLATE = `
type {{.Name}} string
{{$name := .Name}}
var {{.Plural}} = struct{
	{{with .Instances}}{{ range .}}
	{{.Name}} {{$name}}{{end}}{{end}}
}{
	{{with .Instances}}{{ range .}}{{$name}}("{{.Literal}}"),
	{{end}}{{end}} }
`

const GO_FILE_TEMPLATE = `
package {{.Package}}

{{ range .Imports}}import "{{ . }}"
{{end}}
	
{{ .Code }}
`
