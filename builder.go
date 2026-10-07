package neoqueries

import (
	"fmt"
)

type Params map[string]any
type Props map[Ref]string
type List[V any] []V

type Buildable interface {
	build(builder *Builder) string
}

type builtRefs[T comparable] map[T]Ref

type Builder struct {
	nodeRefs     builtRefs[*Node]
	relationRefs builtRefs[*Relation]
	tokenRefs    builtRefs[string]
	propsRefs    builtRefs[*Props]
	listRefs     builtRefs[any]
	valueRefs    builtRefs[any]
	ptrRefs      builtRefs[any]

	params Params

	nodeID     int
	relationID int
	tokenID    int
	propsID    int
	listID     int
	valueID    int
}

func NewBuilder() *Builder {
	return &Builder{
		params:       make(Params),
		nodeRefs:     make(builtRefs[*Node]),
		relationRefs: make(builtRefs[*Relation]),
		tokenRefs:    make(builtRefs[string]),
		propsRefs:    make(builtRefs[*Props]),
		listRefs:     make(builtRefs[any]),
		valueRefs:    make(builtRefs[any]),
		ptrRefs:      make(builtRefs[any]),
	}
}

func (b *Builder) GetParams() Params {
	return b.params
}

func (b *Builder) next[T comparable](obj T, ref Ref, id *int, refMap builtRefs[T]) Ref {
	if savedRef, ok := refMap[obj]; ok {
		return savedRef
	}

	ref = Ref(fmt.Sprintf("%s%d", ref, *id))
	*id++

	refMap[obj] = ref
	return ref
}

func (b *Builder) addToParams(obj any, ref Ref) {
	b.params[ref.String()] = obj
}

func (b *Builder) ensureNodeRef(element *Node) Ref {
	return b.next(element, NodeRef, &b.nodeID, b.nodeRefs)
}
func (b *Builder) ensureRelationRef(element *Relation) Ref {
	return b.next(element, RelationRef, &b.relationID, b.relationRefs)
}
func (b *Builder) ensureTokenRef(value string) Ref {
	ref := b.next(value, TokenRef, &b.tokenID, b.tokenRefs)
	b.addToParams(value, ref)
	return ref
}
func (b *Builder) ensurePropsRef(value *Props) Ref {
	ref := b.next(value, PropsRef, &b.propsID, b.propsRefs)
	b.addToParams(*value, ref)
	return ref
}

func (b *Builder) ensureValueRef(value any) Ref {
	ref := b.next(value, ValueRef, &b.valueID, b.valueRefs)
	b.addToParams(value, ref)
	return ref
}

func (b *Builder) ensureValuePtrRef[T any](value *T) Ref {
	if savedRef, ok := b.listRefs[value]; ok {
		return savedRef
	}

	ref := Ref(fmt.Sprintf("%s%d", ValueRef, b.valueID))
	b.ptrRefs[value] = ref

	if value == nil {
		b.addToParams(nil, ref)
	} else {
		b.addToParams(*value, ref)
	}

	return ref
}

// ensureListRef doesn't use b.next() because *List[V] doesn't match with b.next() generic
func (b *Builder) ensureListRef[V any](value *List[V]) Ref {
	if savedRef, ok := b.listRefs[value]; ok {
		return savedRef
	}

	ref := Ref(fmt.Sprintf("%s%d", ListRef, b.listID))
	b.listID++

	b.listRefs[value] = ref
	b.addToParams(*value, ref)

	return ref
}
