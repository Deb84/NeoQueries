package neoqueries

import (
	"strings"
)

type ReturnElementsBuilder struct {
	*Builder
	elements []Element
}

func newReturnElementsBuilder(builder *Builder, elements ...Element) *ReturnElementsBuilder {
	return &ReturnElementsBuilder{
		Builder:  builder,
		elements: elements,
	}
}

func NewReturnElementsBuilder(elements Element) *ReturnElementsBuilder {
	return newReturnElementsBuilder(NewBuilder(), elements)
}

func (qb *QueryBuilder) ReturnElements(elements ...Element) *QueryBuilder {
	qb.addPart(newReturnElementsBuilder(qb.Builder, elements...))
	return qb
}

func (rb *ReturnElementsBuilder) build() string {
	query := `RETURN`
	var b strings.Builder

	b.WriteString(query)
	b.WriteByte(' ')

	for i, element := range rb.elements {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(element.GetRef().String())
	}

	return b.String()
}

func (rb *ReturnElementsBuilder) setBuilder(builder *Builder) {
	rb.Builder = builder
}

type ReturnBuilder struct {
	*Builder
	refs []QueryRef
}

func newReturnBuilder(builder *Builder, refs ...QueryRef) *ReturnBuilder {
	return &ReturnBuilder{
		Builder: builder,
		refs:    refs,
	}
}

func NewReturnBuilder(refs ...QueryRef) *ReturnBuilder {
	return newReturnBuilder(nil, refs...)
}

func (rb *ReturnBuilder) build() string {
	query := `RETURN`
	var b strings.Builder

	b.WriteString(query)
	b.WriteByte(' ')

	for i, ref := range rb.refs {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(ref.toRef(rb.Builder).String())
	}

	return b.String()
}

func (rb *ReturnBuilder) setBuilder(builder *Builder) {
	rb.Builder = builder
}
