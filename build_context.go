package neoqueries

import (
	"fmt"
)

type Params map[string]any
type Props map[Ref]string
type List[V any] []V

type Buildable interface {
	build(builder *BuildContext) string
}

type builtRefs[T comparable] map[T]Ref

type BuildContext struct {
	config Config

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

func NewBuildContext(config Config) *BuildContext {
	return &BuildContext{
		config:       config,
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

func (b *BuildContext) GetParams() Params {
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

func (b *BuildContext) next[T comparable](obj T, ref Ref, id *int, refMap builtRefs[T]) Ref {
	return storeRef(obj, ref, id, refMap)
}

func (b *BuildContext) nextWithPtr[T comparable](obj T, ref Ref, id *int, refMap builtRefs[any]) Ref {
	return storeAnyRef(obj, ref, id, refMap)
}

func (b *BuildContext) nextWithParams[T comparable](obj T, params any, ref Ref, id *int, refMap builtRefs[T]) Ref {
	ref = storeRef(obj, ref, id, refMap)
	b.addToParams(params, ref)
	return ref
}

func (b *BuildContext) nextWithParamsPtr[T comparable](obj T, params any, ref Ref, id *int, refMap builtRefs[any]) Ref {
	ref = storeAnyRef(obj, ref, id, refMap)
	b.addToParams(params, ref)
	return ref
}

func (b *BuildContext) addToParams(obj any, ref Ref) {
	b.params[ref.String()] = obj
}

func (b *BuildContext) ensureNodeRef(element *Node) Ref {
	return b.next(element, b.config.Ref.NodeRef, &b.nodeID, b.nodeRefs)
}
func (b *BuildContext) ensureRelationRef(element *Relation) Ref {
	return b.next(element, b.config.Ref.RelationRef, &b.relationID, b.relationRefs)
}
func (b *BuildContext) ensureTokenRef(value string) Ref {
	return b.nextWithParams(value, value, b.config.Ref.TokenRef, &b.tokenID, b.tokenRefs)
}
func (b *BuildContext) ensurePropsRef(value *Props) Ref {
	return b.nextWithParams(value, *value, b.config.Ref.PropsRef, &b.propsID, b.propsRefs)
}

func (b *BuildContext) ensureValueRef(value any) Ref {
	ref := b.nextWithParams(value, value, b.config.Ref.ValueRef, &b.valueID, b.valueRefs)
	return ref
}

func (b *BuildContext) ensureValuePtrRef[T any](value *T) Ref {
	ref := b.nextWithPtr(value, b.config.Ref.ValueRef, &b.valueID, b.ptrRefs)

	if value == nil {
		b.addToParams(nil, ref)
	} else {
		b.addToParams(*value, ref)
	}

	return ref
}

// ensureListRef doesn't use b.next() because *List[V] doesn't match with b.next() generic
func (b *BuildContext) ensureListRef[V any](value *List[V]) Ref {
	return b.nextWithParamsPtr(value, *value, b.config.Ref.ListRef, &b.listID, b.listRefs)
}
