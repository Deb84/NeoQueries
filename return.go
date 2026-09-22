package queries

import (
	"NeoQueries"
	"strings"
)

type ReturnElementsBuilder struct {
	*NeoQueries.Builder
	elements []Element
}

func newReturnElementsBuilder(builder *NeoQueries.Builder, elements ...Element) *ReturnElementsBuilder {
	return &ReturnElementsBuilder{
		NeoQueries.Builder: builder,
		elements:           elements,
	}
}

func NewReturnElementsBuilder(elements Element) *ReturnElementsBuilder {
	return newReturnElementsBuilder(NeoQueries.NewBuilder(), elements)
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

func (rb *ReturnElementsBuilder) setBuilder(builder *NeoQueries.Builder) {
	rb.Builder = builder
}

type ReturnBuilder struct {
	*NeoQueries.Builder
	refs []QueryRef
}

func newReturnBuilder(builder *NeoQueries.Builder, refs ...QueryRef) *ReturnBuilder {
	return &ReturnBuilder{
		NeoQueries.Builder: builder,
		refs:               refs,
	}
}

func NewReturnBuilder(refs ...QueryRef) *ReturnBuilder {
	return newReturnBuilder(NeoQueries.NewBuilder(), refs...)
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

func (rb *ReturnBuilder) setBuilder(builder *NeoQueries.Builder) {
	rb.Builder = builder
}
