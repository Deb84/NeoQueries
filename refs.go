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

type refBase string

func (v refBase) String() string {
	return string(v)
}

// Ref is an exported version of ref to avoid exposing Ref on NodeRef and RelationRef
type Ref struct {
	refBase
}

type NodeRef struct {
	refBase
}

type RelationRef struct {
	refBase
}

type Reference interface {
	NodeRef | RelationRef | Ref
	String() string
}

type Prop string

func (p Prop) String() string {
	return string(p)
}

type UnbuiltRef[T Reference] struct {
	owner Element[T]
}

func newUnbuiltRef[T Reference](owner Element[T]) *UnbuiltRef[T] {
	return &UnbuiltRef[T]{
		owner: owner,
	}
}

func (r *UnbuiltRef[T]) buildRef(context *BuildContext) T {
	return r.owner.buildElement(context)
}

func (r *UnbuiltRef[T]) Prop(prop string) *UnbuiltProp[T] {
	return newProp(r, prop)
}

type UnbuiltProp[T Reference] struct {
	variable *UnbuiltRef[T]
	prop     string
}

func newProp[T Reference](variable *UnbuiltRef[T], prop string) *UnbuiltProp[T] {
	return &UnbuiltProp[T]{
		variable: variable,
		prop:     prop,
	}
}

func (p *UnbuiltProp[T]) buildRef(context *BuildContext) Prop {
	vari := p.variable.buildRef(context)

	template := "%s[$%s]"
	fieldRef := context.ensureTokenRef(p.prop)

	prop := fmt.Sprintf(template, vari, fieldRef)

	return Prop(prop)
}
