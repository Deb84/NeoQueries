package neoqueries

import (
	"errors"
	"fmt"
)

type PropOrRef interface {
	Prop | Ref
	String() string
}

type UnbuiltPropOrVar[T PropOrRef] interface {
	buildRef(builder *Builder) T
}

type Ref string

func (v Ref) String() string {
	return string(v)
}

type Prop string

func (p Prop) String() string {
	return string(p)
}

type UnbuiltRef struct {
	owner ElementInterface
}

func newUnbuiltRef(owner ElementInterface) *UnbuiltRef {
	return &UnbuiltRef{
		owner: owner,
	}
}

func (r *UnbuiltRef) string() (string, error) {
	v := r.owner.getRef()
	if v != "" {
		return v.String(), nil
	}

	return "", errors.New("variable not built")
}

func (r *UnbuiltRef) buildRef(builder *Builder) Ref {
	return r.owner.buildElement(builder)
}

func (r *UnbuiltRef) Prop(prop string) *UnbuiltProp {
	return newProp(r, prop)
}

type UnbuiltProp struct {
	variable *UnbuiltRef
	prop     string
}

func newProp(variable *UnbuiltRef, prop string) *UnbuiltProp {
	return &UnbuiltProp{
		variable: variable,
		prop:     prop,
	}
}

func (p *UnbuiltProp) buildRef(builder *Builder) Prop {
	vari := p.variable.buildRef(builder)

	template := "%s[$%s]"
	fieldRef := builder.nextTokenRef(p.prop)

	prop := fmt.Sprintf(template, vari, fieldRef)

	return Prop(prop)
}
