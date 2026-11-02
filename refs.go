package neoqueries

import (
	"fmt"
)

type PropOrRef interface {
	Prop | Ref
	String() string
}

type UnbuiltPropOrVar[T PropOrRef] interface {
	buildRef(context *BuildContext) T
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
	owner Element
}

func newUnbuiltRef(owner Element) *UnbuiltRef {
	return &UnbuiltRef{
		owner: owner,
	}
}

func (r *UnbuiltRef) buildRef(context *BuildContext) Ref {
	return r.owner.buildElement(context)
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

func (p *UnbuiltProp) buildRef(context *BuildContext) Prop {
	vari := p.variable.buildRef(context)

	template := "%s[$%s]"
	fieldRef := context.ensureTokenRef(p.prop)

	prop := fmt.Sprintf(template, vari, fieldRef)

	return Prop(prop)
}
