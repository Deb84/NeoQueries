package neoqueries

import "strings"

type DeleteBuilder struct {
	*Builder
	elements []Element
	query    string
}

func newBaseDeleteBuilder(builder *Builder, query string, element Element, elements []Element) *DeleteBuilder {
	elements = append([]Element{element}, elements...)

	for _, e := range elements {
		e.setBuilder(builder)
	}

	return &DeleteBuilder{
		Builder:  builder,
		elements: elements,
		query:    query,
	}
}

func newDeleteBuilder(builder *Builder, element Element, elements []Element) *DeleteBuilder {
	return newBaseDeleteBuilder(builder, "DELETE", element, elements)
}

func newDetachDeleteBuilder(builder *Builder, element Element, elements []Element) *DeleteBuilder {
	return newBaseDeleteBuilder(builder, "DETACH DELETE", element, elements)
}

func NewDeleteBuilder(element Element, elements ...Element) *DeleteBuilder {
	return newBaseDeleteBuilder(nil, "DELETE", element, elements)
}

func NewDetachDeleteBuilder(element Element, elements ...Element) *DeleteBuilder {
	return newBaseDeleteBuilder(nil, "DETACH DELETE", element, elements)
}

func (b *DeleteBuilder) build() string {
	var s strings.Builder

	s.WriteString(b.query)
	s.WriteByte(' ')

	for i, element := range b.elements {
		if i > 0 {
			s.WriteString(", ")
		}
		s.WriteString(element.build())
	}

	return s.String()
}
