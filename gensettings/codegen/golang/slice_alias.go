package golang

import (
	"bytes"
	"text/template"

	"github.com/dtcookie/dynatrace/gensettings/reflection"

	"github.com/dtcookie/dynatrace/gensettings/codegen/camel"
)

type SliceAlias struct {
	Name string
	Elem string
	Item string
}

func (me *SliceAlias) Prettify() {
	me.Name = PrettyProp(me.Name)
}

func (me *SliceAlias) Bytes() ([]byte, error) {
	tmpl, err := template.New("GO_LIST_ALIAS").Parse(GO_LIST_ALIAS_TEMPLATE)
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

func (me *SliceAlias) FileName() string {
	return camel.Strip(me.Elem) + ".go"
}

func (me *SliceAlias) Imports() []string {
	return []string{"github.com/dtcookie/hcl"}
}

func (me *SliceAlias) Kind() CodeContributorKind {
	return AliasContrib
}

func (me *SliceAlias) SortName() string {
	return me.Name
}

func NewSliceAlias(t *reflection.Type, item string) *SliceAlias {
	return &SliceAlias{
		Name: Plural(TypeName(t.Elem.Unref())),
		Elem: TypeName(t.Elem.Unref()),
		Item: item,
	}
}
