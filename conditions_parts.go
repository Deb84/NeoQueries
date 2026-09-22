package NeoQueries

import (
	"fmt"
)

type conditionPart interface {
	build(*ConditionBuilder) string
}

type refConditionPart struct {
	ref QueryRef
}

func (p *refConditionPart) build(cb *ConditionBuilder) string {
	return p.ref.toRef(cb.Builder).String()
}

type nestedConditionPart struct {
	condition ConditionBuilderInterface
}

func (p *nestedConditionPart) build(cb *ConditionBuilder) string {
	p.condition.setBuilder(cb.Builder)
	return "(" + p.condition.buildCondition() + ")"
}

type rawPart string

func (p rawPart) build(*ConditionBuilder) string {
	return string(p)
}

type anyPart struct {
	value any
}

func (p *anyPart) build(cb *ConditionBuilder) string {
	return cb.nextValueRef(p.value).String()
}

type listPart[V any] struct {
	list *List[V]
}

func (p listPart[V]) build(cb *ConditionBuilder) string {
	ref := cb.nextListRef(p.list)
	return "$" + ref.String()
}

type ConditionStart struct {
	builder *ConditionBuilder
}

func (c *ConditionStart) String(value string) *StartString {
	c.builder.addPart(&anyPart{value: value})
	return &StartString{builder: c.builder}
}

func (c *ConditionStart) Value[T any](value T) *AnyState[T] {
	c.builder.addPart(&anyPart{value: value})
	return &AnyState[T]{builder: c.builder}
}

func (c *ConditionStart) Ref(value QueryRef) *StartRef {
	c.builder.addPart(&refConditionPart{ref: value})
	return newStartRef(c.builder)
}

func (c *ConditionStart) Condition(value *ConditionBuilder) *CompleteCondition {
	c.builder.addPart(&nestedConditionPart{condition: value})
	return &CompleteCondition{c.builder}
}

type StartRef struct {
	builder *ConditionBuilder
	*AnyState[any]
	*StartString
}

func newStartRef(builder *ConditionBuilder) *StartRef {
	return &StartRef{
		builder:     builder,
		AnyState:    &AnyState[any]{builder: builder},
		StartString: &StartString{builder: builder},
	}
}

type StartString struct {
	builder *ConditionBuilder
}

func (p *StartString) StartsWith() *EndString {
	p.builder.addPart(rawPart("STARTS WITH"))
	return &EndString{builder: p.builder}
}

func (p *StartString) EndsWith() *EndString {
	p.builder.addPart(rawPart("ENDS WITH"))
	return &EndString{builder: p.builder}
}

func (p *StartString) Contains() *EndString {
	p.builder.addPart(rawPart("CONTAINS"))
	return &EndString{builder: p.builder}
}

type EndString struct {
	builder *ConditionBuilder
}

func (p *EndString) Ref(value QueryRef) *CompleteCondition {
	p.builder.addPart(&refConditionPart{ref: value})
	return &CompleteCondition{p.builder}
}

func (p *EndString) String(value string) *CompleteCondition {
	p.builder.addPart(&anyPart{value: value})
	return &CompleteCondition{p.builder}
}

type AnyState[T any] struct {
	builder *ConditionBuilder
}

func (p *AnyState[T]) In() *ListState[T] {
	p.builder.addPart(rawPart("IN"))
	return &ListState[T]{builder: p.builder}
}

func (p *AnyState[T]) IsNull() *CompleteCondition {
	p.builder.addPart(rawPart("IS NULL"))
	return &CompleteCondition{p.builder}
}

func (p *AnyState[T]) IsNotNull() *CompleteCondition {
	p.builder.addPart(rawPart("IS NOT NULL"))
	return &CompleteCondition{p.builder}
}

func (p *AnyState[T]) Eq() *Comparison[T] {
	p.builder.addPart(rawPart(Equal))
	return &Comparison[T]{builder: p.builder}
}

func (p *AnyState[T]) Neq() *Comparison[T] {
	p.builder.addPart(rawPart(NotEqual))
	return &Comparison[T]{builder: p.builder}
}

func (p *AnyState[T]) Gt() *Comparison[T] {
	p.builder.addPart(rawPart(GreaterThan))
	return &Comparison[T]{builder: p.builder}
}

func (p *AnyState[T]) Lt() *Comparison[T] {
	p.builder.addPart(rawPart(LessThan))
	return &Comparison[T]{builder: p.builder}
}

func (p *AnyState[T]) Gte() *Comparison[T] {
	p.builder.addPart(rawPart(GreaterThanOrEqual))
	return &Comparison[T]{builder: p.builder}
}

func (p *AnyState[T]) Lte() *Comparison[T] {
	p.builder.addPart(rawPart(LessThanOrEqual))
	return &Comparison[T]{builder: p.builder}
}

type ListState[V any] struct {
	builder *ConditionBuilder
}

func (p *ListState[V]) List(list *List[V]) *CompleteCondition {
	p.builder.addPart(listPart[V]{list: list})
	return &CompleteCondition{p.builder}
}

func (p *ListState[T]) Ref(value QueryRef) *CompleteCondition {
	p.builder.addPart(rawPart(fmt.Sprint(value)))
	return &CompleteCondition{p.builder}
}

type Comparison[T any] struct {
	builder *ConditionBuilder
}

func (p *Comparison[T]) Ref(value QueryRef) *CompleteCondition {
	p.builder.addPart(&refConditionPart{ref: value})
	return &CompleteCondition{p.builder}
}

func (p *Comparison[T]) Value(value T) *CompleteCondition {
	p.builder.addPart(&anyPart{value: value})
	return &CompleteCondition{p.builder}
}

type CompleteCondition struct {
	*ConditionBuilder
}

func (p *CompleteCondition) And() *ConditionStart {
	p.addPart(rawPart("AND"))
	return &ConditionStart{builder: p.ConditionBuilder}
}

func (p *CompleteCondition) Or() *ConditionStart {
	p.addPart(rawPart("OR"))
	return &ConditionStart{builder: p.ConditionBuilder}
}
