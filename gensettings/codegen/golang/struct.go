package golang

import (
	"bytes"
	"sort"
	"strings"
	"text/template"

	"github.com/dtcookie/dynatrace/gensettings/reflection"

	"github.com/dtcookie/dynatrace/gensettings/codegen/camel"
)

type Property struct {
	Name     string
	Type     string
	Comment  string
	HCLTag   string
	JSONTag  string
	Optional bool
	MinItems int
	MaxItems int
	HCLType  string
	Elem     string
}

func (me *Property) Prettify() {
	me.Name = PrettyProp(me.Name)
}

type Struct struct {
	Name       string
	Plural     string
	Comments   []string
	Properties []*Property
}

func (me *Struct) Bytes() ([]byte, error) {
	tmpl, err := template.New("GO_STRUCT").Parse(GO_STRUCT_TEMPLATE)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, me)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (me *Struct) Prettify() {
	me.Name = PrettyProp(me.Name)
	me.Plural = Plural(me.Name)
	if len(me.Properties) > 0 {
		for _, property := range me.Properties {
			property.Prettify()
		}
	}
}

func (me *Struct) FileName() string {
	return camel.Strip(me.Name) + ".go"
}

func (me *Struct) Imports() []string {
	return []string{"github.com/dtcookie/hcl"}
}

func (me *Struct) Kind() CodeContributorKind {
	return StructContrib
}

func (me *Struct) SortName() string {
	return me.Name
}

func NewStruct(t *reflection.Type) *Struct {
	structDef := Struct{
		Name:     TypeName(t),
		Plural:   Plural(TypeName(t)),
		Comments: Comments(strings.Split(t.Comment, "\n")),
	}

	for propertyName, property := range t.Properties {
		propType := property.Type
		if propType.Kind == reflection.StructKind {
			propType = propType.Pointer()
		}
		// if propType.Kind == reflection.ListKind {
		// 	if propType.Elem.Kind == reflection.StructKind {
		// 		propType = propType.Elem.Pointer().List()
		// 	}
		// }
		// if propType.Kind == reflection.SetKind {
		// 	if propType.Elem.Kind == reflection.StructKind {
		// 		propType = propType.Elem.Pointer().Set()
		// 	}
		// }

		structDef.Properties = append(structDef.Properties, HCLKind(propType, &Property{
			Name:     PropertyName(propertyName),
			Comment:  strings.ReplaceAll(Comment(property.Comment), "\n", "\\n"),
			HCLTag:   camel.Strip(propertyName),
			JSONTag:  propertyName,
			Type:     TypeName(propType),
			Optional: property.Optional,
		}))
	}
	sort.Slice(structDef.Properties, func(i, j int) bool {
		return strings.Compare(structDef.Properties[i].Name, structDef.Properties[j].Name) == -1
	})
	return &structDef
}

func StructTypeName(name string) string {
	return strings.ToUpper(name[0:1]) + name[1:]
}

func PropertyName(name string) string {
	propertyName := strings.ToUpper(name[0:1]) + name[1:]
	propertyName = strings.ReplaceAll(propertyName, " ", "")
	propertyName = strings.ReplaceAll(propertyName, "-", "_")
	propertyName = strings.ReplaceAll(propertyName, "+", "_")
	propertyName = strings.ReplaceAll(propertyName, ".", "_")
	propertyName = strings.ReplaceAll(propertyName, ":", "_")
	return PrettyProp(propertyName)
}

func HCLKind(propType *reflection.Type, property *Property) *Property {
	switch propType.UnrefPtr().Kind {
	case reflection.BooleanKind:
		property.HCLType = "hcl.TypeBool"
		property.Optional = true
	case reflection.FloatKind:
		property.HCLType = "hcl.TypeFloat"
	case reflection.EnumKind:
		property.HCLType = "hcl.TypeString"
	case reflection.StringKind:
		property.HCLType = "hcl.TypeString"
	case reflection.IntKind:
		property.HCLType = "hcl.TypeInt"
	case reflection.StructKind:
		property.HCLType = "hcl.TypeList"
		property.MaxItems = 1
		property.MinItems = 1
		property.Elem = "&hcl.Resource{Schema: new(" + TypeName(propType.Elem.Unref()) + ").Schema()}"
	case reflection.SliceAliasKind, reflection.SetAliasKind:
		property.HCLType = "hcl.TypeList"
		property.MinItems = 1
		property.MaxItems = 1
		property.Elem = "&hcl.Resource{Schema: new(" + TypeName(propType) + ").Schema()}"
	case reflection.ListKind:
		property.HCLType = "hcl.TypeList"
		switch propType.Elem.Kind {
		case reflection.StringKind:
			property.Elem = "&hcl.Schema{Type: hcl.TypeString}"
		case reflection.EnumKind:
			property.Elem = "&hcl.Schema{Type: hcl.TypeString}"
		case reflection.IntKind:
			property.Elem = "&hcl.Schema{Type: hcl.TypeInt}"
		case reflection.FloatKind:
			property.Elem = "&hcl.Schema{Type: hcl.TypeFloat}"
		case reflection.BooleanKind:
			property.Elem = "&hcl.Schema{Type: hcl.TypeBool}"
		default:
			property.MinItems = 1
			property.Elem = "&hcl.Resource{Schema: new(" + TypeName(propType.Elem.Unref()) + ").Schema()}"
		}
	case reflection.SetKind:
		property.HCLType = "hcl.TypeSet"
		switch propType.Elem.Kind {
		case reflection.StringKind:
			property.Elem = "&hcl.Schema{Type: hcl.TypeString}"
		case reflection.EnumKind:
			property.Elem = "&hcl.Schema{Type: hcl.TypeString}"
		case reflection.IntKind:
			property.Elem = "&hcl.Schema{Type: hcl.TypeInt}"
		case reflection.FloatKind:
			property.Elem = "&hcl.Schema{Type: hcl.TypeFloat}"
		case reflection.BooleanKind:
			property.Elem = "&hcl.Schema{Type: hcl.TypeBool}"
		default:
			property.MinItems = 1
			property.Elem = "&hcl.Resource{Schema: new(" + TypeName(propType.Elem.Unref()) + ").Schema()}"
		}
	}
	return property
}
