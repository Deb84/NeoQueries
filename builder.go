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

func storeRefBase(ref Ref, id *int) Ref {
	*id++
	return Ref(fmt.Sprintf("%s%d", ref, *id))
}

func storeRef[V comparable](value V, ref Ref, id *int, refMap builtRefs[V]) Ref {
	if savedRef, ok := refMap[value]; ok {
		return savedRef
	}

	ref = storeRefBase(ref, id)

	refMap[value] = ref
	return ref
}

func storeAnyRef[V comparable](value V, ref Ref, id *int, refMap builtRefs[any]) Ref {
	if savedRef, ok := refMap[value]; ok {
		return savedRef
	}

	ref = storeRefBase(ref, id)

	refMap[value] = ref
	return ref
}

func (b *Builder) next[T comparable](obj T, ref Ref, id *int, refMap builtRefs[T]) Ref {
	return storeRef(obj, ref, id, refMap)
}

func (b *Builder) nextWithPtr[T comparable](obj T, ref Ref, id *int, refMap builtRefs[any]) Ref {
	return storeAnyRef(obj, ref, id, refMap)
}

func (b *Builder) nextWithParams[T comparable](obj T, params any, ref Ref, id *int, refMap builtRefs[T]) Ref {
	ref = storeRef(obj, ref, id, refMap)
	b.addToParams(params, ref)
	return ref
}

func (b *Builder) nextWithParamsPtr[T comparable](obj T, params any, ref Ref, id *int, refMap builtRefs[any]) Ref {
	ref = storeAnyRef(obj, ref, id, refMap)
	b.addToParams(params, ref)
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
	return b.nextWithParams(value, value, TokenRef, &b.tokenID, b.tokenRefs)
}
func (b *Builder) ensurePropsRef(value *Props) Ref {
	return b.nextWithParams(value, *value, PropsRef, &b.propsID, b.propsRefs)
}

func (b *Builder) ensureValueRef(value any) Ref {
	ref := b.nextWithParams(value, value, ValueRef, &b.valueID, b.valueRefs)
	return ref
}

func (b *Builder) ensureValuePtrRef[T any](value *T) Ref {
	ref := b.nextWithPtr(value, ValueRef, &b.valueID, b.ptrRefs)

	if value == nil {
		b.addToParams(nil, ref)
	} else {
		b.addToParams(*value, ref)
	}

	return ref
}

// ensureListRef doesn't use b.next() because *List[V] doesn't match with b.next() generic
func (b *Builder) ensureListRef[V any](value *List[V]) Ref {
	return b.nextWithParamsPtr(value, *value, ValueRef, &b.listID, b.listRefs)
}
