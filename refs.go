package neoqueries

import (
	"fmt"
)

// PropOrRef defines a constraint for types that are either a Prop or a Ref
type PropOrRef interface {
	Prop | Ref
	String() string
}

// UnbuiltPropOrVar defines an unbuilt PropOrRef constraint
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

// Reference defines a constraint for types that are either A Ref, NodeRef or RelationRef
type Reference interface {
	NodeRef | RelationRef | Ref
	String() string
}

// Prop is a representation of neo4j prop
type Prop string

func (p Prop) String() string {
	return string(p)
}

// UnbuiltRef represents a Ref which will be built at build
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

// UnbuiltProp represents a Prop which will be built at build
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
