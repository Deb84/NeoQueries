package neoqueries

import (
	"fmt"
)

type Params map[string]any
type Props map[string]string
type List[V any] []V

type Buildable interface {
	build(builder *BuildContext) string
}

type builtRefs[T comparable] map[T]refBase

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

func storeRefBase[T Reference](ref T, id *int) refBase {
	*id++
	return refBase(fmt.Sprintf("%s%d", ref.String(), *id))
}

func storeRef[V comparable, R Reference](value V, ref R, id *int, refMap builtRefs[V]) refBase {
	if savedRef, ok := refMap[value]; ok {
		return savedRef
	}

	r := storeRefBase(ref, id)

	refMap[value] = r
	return r
}

func storeAnyRef[V comparable, R Reference](value V, ref R, id *int, refMap builtRefs[any]) refBase {
	if savedRef, ok := refMap[value]; ok {
		return savedRef
	}

	r := storeRefBase(ref, id)

	refMap[value] = r
	return r
}

func (b *BuildContext) next[T comparable, R Reference](obj T, ref R, id *int, refMap builtRefs[T]) refBase {
	return storeRef(obj, ref, id, refMap)
}

func (b *BuildContext) nextWithPtr[T comparable, R Reference](obj T, ref R, id *int, refMap builtRefs[any]) refBase {
	return storeAnyRef(obj, ref, id, refMap)
}

func (b *BuildContext) nextWithParams[T comparable, R Reference](obj T, params any, ref R, id *int, refMap builtRefs[T]) refBase {
	r := storeRef(obj, ref, id, refMap)
	b.addToParams(params, r)
	return r
}

func (b *BuildContext) nextWithParamsPtr[T comparable, R Reference](obj T, params any, ref R, id *int, refMap builtRefs[any]) refBase {
	r := storeAnyRef(obj, ref, id, refMap)
	b.addToParams(params, r)
	return r
}

func (b *BuildContext) addToParams(obj any, ref refBase) {
	b.params[ref.String()] = obj
}

func (b *BuildContext) ensureNodeRef(element *Node) NodeRef {
	return NodeRef{b.next(element, b.config.Ref.NodeRef, &b.nodeID, b.nodeRefs)}
}
func (b *BuildContext) ensureRelationRef(element *Relation) RelationRef {
	return RelationRef{b.next(element, b.config.Ref.RelationRef, &b.relationID, b.relationRefs)}
}
func (b *BuildContext) ensureTokenRef(value string) Ref {
	return Ref{b.nextWithParams(value, value, b.config.Ref.TokenRef, &b.tokenID, b.tokenRefs)}
}
func (b *BuildContext) ensurePropsRef(value *Props) Ref {
	return Ref{b.nextWithParams(value, *value, b.config.Ref.PropsRef, &b.propsID, b.propsRefs)}
}

func (b *BuildContext) ensureValueRef(value any) Ref {
	return Ref{b.nextWithParams(value, value, b.config.Ref.ValueRef, &b.valueID, b.valueRefs)}
}

func (b *BuildContext) ensureValuePtrRef[T any](value *T) Ref {
	ref := b.nextWithPtr(value, b.config.Ref.ValueRef, &b.valueID, b.ptrRefs)

	if value == nil {
		b.addToParams(nil, ref)
	} else {
		b.addToParams(*value, ref)
	}

	return Ref{ref}
}

// ensureListRef doesn't use b.next() because *List[V] doesn't match with b.next() generic
func (b *BuildContext) ensureListRef[V any](value *List[V]) Ref {
	return Ref{b.nextWithParamsPtr(value, *value, b.config.Ref.ListRef, &b.listID, b.listRefs)}
}
