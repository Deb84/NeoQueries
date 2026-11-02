package neoqueries

import "strings"

type DeleteBuilder struct {
	elements []Element
	query    string
}

func newBaseDeleteBuilder(query string, element Element, elements []Element) *DeleteBuilder {
	elements = append([]Element{element}, elements...)

	return &DeleteBuilder{
		elements: elements,
		query:    query,
	}
}

func newDeleteBuilder(element Element, elements []Element) *DeleteBuilder {
	return newBaseDeleteBuilder("DELETE", element, elements)
}

func newDetachDeleteBuilder(element Element, elements []Element) *DeleteBuilder {
	return newBaseDeleteBuilder("DETACH DELETE", element, elements)
}

func NewDeleteBuilder(element Element, elements ...Element) *DeleteBuilder {
	return newBaseDeleteBuilder("DELETE", element, elements)
}

func NewDetachDeleteBuilder(element Element, elements ...Element) *DeleteBuilder {
	return newBaseDeleteBuilder("DETACH DELETE", element, elements)
}

func (b *DeleteBuilder) build(context *BuildContext) string {
	var s strings.Builder

	s.WriteString(b.query)
	s.WriteByte(' ')

	for i, element := range b.elements {
		if i > 0 {
			s.WriteString(", ")
		}
		s.WriteString(element.buildElement(context).String())
	}

	return s.String()
}
