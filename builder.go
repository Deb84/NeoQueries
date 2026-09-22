package queries

import "fmt"

type Params map[Ref]any
type Props map[Ref]string
type List[V any] []V

type Buildable interface {
	build() string
	GetParams() Params
}

type Refs struct {
	element Ref
	props   Ref
	token   map[string]Ref
}

func newRefs() Refs {
	return Refs{
		token: make(map[string]Ref),
	}
}

type Builder struct {
	builtRefs map[any]Ref
	params    Params

	nodeID     int
	relationID int
	tokenID    int
	propsID    int
	listID     int
	valueID    int
}

func NewBuilder() *Builder {
	return &Builder{
		params:    make(Params),
		builtRefs: make(map[any]Ref),
	}
}

func (b *Builder) GetParams() Params {
	return b.params
}

func (b *Builder) next(obj any, ref Ref, id *int) Ref {
	if savedRef := b.builtRefs[obj]; savedRef != "" {
		return savedRef
	}
	ref = Ref(fmt.Sprintf("%s%d", ref, *id))
	b.builtRefs[obj] = ref
	*id++
	return ref
}

func (b *Builder) addToParams(obj any, ref Ref) {
	b.params[ref] = obj
}

func (b *Builder) nextNodeRef(element *NodeBuilder) Ref {
	return b.next(element, NodeRef, &b.nodeID)
}
func (b *Builder) nextRelationRef(element *RelationBuilder) Ref {
	return b.next(element, RelationRef, &b.relationID)
}
func (b *Builder) nextTokenRef(value string) Ref {
	ref := b.next(value, TokenRef, &b.tokenID)
	b.addToParams(value, ref)
	return ref
}
func (b *Builder) nextPropsRef(value *Props) Ref {
	ref := b.next(value, PropsRef, &b.propsID)
	b.addToParams(*value, ref)
	return ref
}

func (b *Builder) nextListRef[V any](value *List[V]) Ref {
	ref := b.next(value, ListRef, &b.listID)
	b.addToParams(*value, ref)
	return ref
}

func (b *Builder) nextValueRef(value any) Ref {
	ref := b.next(value, ValueRef, &b.valueID)
	b.addToParams(value, ref)
	return ref
}
